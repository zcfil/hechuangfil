package main

import (
	"github.com/BurntSushi/toml"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
	"hechuangfil/auth"
	"hechuangfil/conf"
	"hechuangfil/controller"
	"hechuangfil/db"
	_ "hechuangfil/docs"
	"hechuangfil/models"
	"hechuangfil/router"
	"hechuangfil/service"
	"log"
	"math/rand"
	"os"
	"time"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email 1503780117@qq.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 183.61.251.226:3000
// @BasePath
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name token
func main() {
	rand.Seed(time.Now().UnixNano())
	app := &cli.App{
		Flags: []cli.Flag {
			&cli.StringFlag{
				Name: "path",
				Value: "./conf/config.toml",
				Usage: "启动的配置文件路径格式为toml",
			},
			&cli.StringFlag{
				Name: "geo",
				Value: "./conf/GeoLite2-City.mmdb",
				Usage: "GeoIP数据库文件路径,将IP地址解析出地理位置（国家，城市，经纬度）",
			},
		},
		Action: func(c *cli.Context) error {

			var config conf.Config
			if _, err := toml.DecodeFile(c.String("path"), &config); err != nil {
				// handle error
				return  err
			}

			container := dig.New()
			container.Provide(func() *conf.Dysmsapi{
				return &conf.Dysmsapi{
					config.Dysmsapi.AccessKeyID,
					config.Dysmsapi.AccessKeySecret,
					config.Dysmsapi.RegisterMsgTemplate,
					config.Dysmsapi.LoginMsgTemplate,
					config.Dysmsapi.SignName,
				}
			})
			container.Provide(func() *conf.Repo{
				return &conf.Repo{
					config.Repo.PaymentImagePath,
					config.Repo.ProductImagePath,
					config.Repo.MaxSize,
				}
			})
			container.Provide(func() *conf.Project{
				return &conf.Project{
					config.Project.Host,
					config.Project.LotusUrl,
					config.Project.Salesmanratio,
				}
			})
			container.Provide(func() *conf.LogFile{
				return &conf.LogFile{
					FileName: config.LogFile.FileName,
				}
			})
			container.Provide(func() *conf.AuthCode{
				return &conf.AuthCode{
					Token: config.AuthCode.Token,
				}
			})
			container.Provide(service.NewResponse)
			container.Provide(controller.NewLog(config.LogFile))
			container.Provide(db.NewSqlEngine(config.Mysql))
			container.Provide(conf.NewRedisClient(config.Redis))
			container.Provide(models.NewUserLevelModels)
			container.Provide(models.NewOrderProfitModels)
			container.Provide(models.NewSysConfigModels)
			container.Provide(controller.NewLogin)
			container.Provide(controller.NewUser)
			container.Provide(controller.NewFinance)
			container.Provide(models.NewNews)
			err := container.Invoke(auth.UpdateOrderStatus)
			//geoDB配置路径
			//container.Provide(conf.NewGeoDB(c.String("geo")))

			container.Provide(conf.LoadCasbin(config.Mysql))
			err = container.Invoke(router.Serve)

			return err
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("启动失败!", err)
	}
}
