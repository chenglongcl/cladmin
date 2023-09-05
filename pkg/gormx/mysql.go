package gormx

import (
	"cladmin/dal/cladmindb/cladminquery"
	"fmt"
	"github.com/chenglongcl/log"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type Database struct {
	Self   *gorm.DB
	Docker *gorm.DB
}

var database *Database

func InitMySQL() {
	database = &Database{
		Self: initDB("db"),
		//Docker:      initDB("sydst_db"),
	}
	cladminquery.SetDefault(database.Self)
}

func initDB(name string) *gorm.DB {
	if name == "" {
		return nil
	}
	return openDB(
		viper.GetString(name+".username"),
		viper.GetString(name+".password"),
		viper.GetString(name+".addr"),
		viper.GetString(name+".name"),
	)
}

func openDB(username, password, addr, name string) *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
			username,
			password,
			addr,
			name,
			true,
			"Local"), // DSN data source name
	}
	var gormLogMode logger.LogLevel
	if viper.GetBool("gormlog") {
		gormLogMode = logger.Info
	} else {
		gormLogMode = logger.Silent
	}
	db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		Logger: logger.Default.LogMode(gormLogMode),
	})
	if err != nil {
		log.Errorf(err, "Database connection failed. Database name: %s", name)
	}
	setupDB(db)
	return db
}

func setupDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(150) // 用于设置最大打开的连接数，默认值为0表示不限制.设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
	sqlDB.SetMaxIdleConns(100) // 用于设置闲置的连接数.设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
	sqlDB.SetConnMaxLifetime(7200 * time.Second)
}

func Close() {
	selfDB, _ := database.Self.DB()
	_ = selfDB.Close()
}

func SelectDB(name string) *gorm.DB {
	switch name {
	case "self":
		return database.Self
	}
	return nil
}
