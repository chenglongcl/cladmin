package main

import (
	"apiserver/config"
	"apiserver/router"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/lexkong/log"
	"net/http"
	"time"
	"apiserver/model"
	"apiserver/router/middleware"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {
	pflag.Parse()
	// init config
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// init db
	model.DB.Init()
	defer model.DB.Close()
	// init redis
	model.RD.Init()
	defer model.RD.Close()
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
	log.Info(http.ListenAndServe(viper.GetString("addr"), g).Error())
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
