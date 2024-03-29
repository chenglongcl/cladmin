package main

import (
	"cladmin/config"
	"cladmin/pkg/aliyun"
	"cladmin/pkg/cloudstorage"
	"cladmin/pkg/gormx"
	"cladmin/pkg/json"
	"cladmin/pkg/redisgo"
	"cladmin/pkg/redsync"
	v "cladmin/pkg/version"
	"cladmin/router"
	"cladmin/router/middleware"
	"cladmin/router/middleware/inject"
	"errors"
	"fmt"
	"github.com/chenglongcl/log"
	"github.com/gin-gonic/gin"
	"github.com/jpillora/overseer"
	"github.com/jpillora/overseer/fetcher"
	"github.com/json-iterator/go"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

var (
	cfg     = pflag.StringP("config", "c", "", "cladmin config file path.")
	version = pflag.BoolP("version", "v", false, "show version info")
)

func main() {
	overseer.Run(overseer.Config{
		Program: program,
		Address: ":8002",
		Fetcher: &fetcher.File{Path: "public/update/cladmin"},
		Debug:   false,
	})
}

func program(state overseer.State) {
	pflag.Parse()
	if *version {
		info := v.Get()
		marshaled, err := jsoniter.MarshalIndent(&info, "", "  ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshaled))
		return
	}
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	//init gorm v2
	gormx.InitMySQL()
	defer gormx.Close()
	// init redis
	_ = redisgo.Init()
	json.Init()
	//init redsync
	redsync.Init()
	//init Casbin
	inject.Init()
	err := inject.LoadCasbinPolicyData()
	//init cloudStorage
	cloudstorage.InitCloudStorage()
	//init aliYun client
	aliyun.InitAliYunOpenApiClients()
	if err != nil {
		log.Fatal("Failure to load Casbin policy data:", err)
	}
	//Set gin mode
	gin.SetMode(viper.GetString("runmode"))
	//Create the gin engine
	g := gin.New()

	router.Load(
		g,

		//Middlewares
		middleware.RequestId(),
		middleware.Logging(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()
	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	//log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
	log.Infof(http.Serve(state.Listener, g).Error())

}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}
		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
