package main

import (
	"cladmin/config"
	"cladmin/dal/cladmindb/cladminmodel"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// init config
	if err := config.Init(*pflag.StringP("config", "c", "", "gormgenrate config file path.")); err != nil {
		panic(err)
	}
	mySQLDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"),
		true,
		"Local")
	// 连接数据库
	db, err := gorm.Open(mysql.Open(mySQLDSN))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	// 生成实例
	g := gen.NewGenerator(gen.Config{
		// 相对执行`go run`时的路径, 会自动创建目录
		OutPath: "../../dal/cladmindb/cladminquery",
		// WithDefaultQuery 生成默认查询结构体(作为全局变量使用), 即`Q`结构体和其字段(各表模型)
		// WithoutContext 生成没有context调用限制的代码供查询
		// WithQueryInterface 生成interface形式的查询代码(可导出), 如`Where()`方法返回的就是一个可导出的接口类型
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	// 设置目标 db
	g.UseDB(db)
	// generate from struct in project
	g.ApplyBasic(
		cladminmodel.SysUser{}, cladminmodel.SysMenu{}, cladminmodel.SysRole{}, cladminmodel.SysRoleMenu{},
		cladminmodel.SysUserRole{}, cladminmodel.SysArticle{}, cladminmodel.SysBulletin{}, cladminmodel.SysCategory{},
		cladminmodel.SysConfig{}, cladminmodel.SysUserToken{},
	)

	g.Execute()
}
