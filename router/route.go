package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v7"
	"github.com/go-xorm/xorm"
	"hechuangfil/conf"
	"hechuangfil/controller"
	"hechuangfil/middleware"
	"hechuangfil/models"
	"hechuangfil/result"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Serve(
	client *redis.Client,
	engine *xorm.Engine,
	repo *conf.Repo,
	pro *conf.Project,
	logConf *conf.LogFile,
	user *controller.User,
	login *controller.Login,
	yungo controller.Yungo,
	news *models.News,
){

	r := gin.Default()
	//cors跨域配置
	r.Use(conf.Cors())
	//文件大小限制
	//r.MaxMultipartMemory = repo.MaxSize
	//云构
	yunGroup := r.Group("/yun")

	yunGroup.Use(middleware.JwtCheck(client))
	{
		//yunGroup.POST("/upload", yungo.Upload)
		//yungoGroup.GET("/list/imports", yungo.ListImportFiles)
		//yunGroup.GET("/ftypes",yungo.ListFileTypes)
		//重置密码
		yunGroup.POST("/resetpwd",yungo.ReSetPwd)


		// 个人信息
		yunGroup.GET("/memberinfo",user.MemberInfo)
		yunGroup.POST("/updatemember",user.UpdateMember) // 修改个人信息

		//yungoGroup.GET("/download",yungo.DownLoad)

		//订单模块
		yunGroup.POST("/placeAnOrder",user.PlaceAnOrder) //下单
		yunGroup.GET("/tashratePrice",user.TashratePrice) //算力单价
		yunGroup.GET("/orderlist",user.Orderlist)	//获取订单列表
		yunGroup.POST("/rechargeCoin",user.RechargeCoin) //充值
		//yunGroup.GET("/hashrateTotal",user.HashrateTotal)         //总算力

		//财务
		yunGroup.GET("/collectInfo1",user.GetCollectInfo) //挖矿收益
		yunGroup.GET("/settleLogList",user.SettleLogList) //收益记录
		yunGroup.GET("/settleLogById",user.SettleLogById) //


		yunGroup.GET("/collectInfo2",user.GetCollectInfo2) //分润收益
		yunGroup.GET("/lowerLevelPerformance",user.LowerLevelPerformance) //直接下级业绩
		yunGroup.GET("/orderProfitList",user.OrderProfitList) //分润收益记录
		yunGroup.GET("/orderProfitById",user.OrderProfitById) //

		yunGroup.GET("/getWalletAddress",user.GetWalletAddress) //充币钱包
		yunGroup.GET("/rechargeList", user.GetRechargeList)      // 获取充值记录列表
		yunGroup.GET("/rechargeById", user.GetRechargeById)      // 获取充值记录

		yunGroup.GET("/withdrawList", user.GetWithdrawList)       // 获取提现记录列表
		yunGroup.GET("/withdrawById", user.GetWithdrawById)       // 获取提现记录
		yunGroup.POST("/withdrawCoin",user.WithdrawCoin)         //提现
		yunGroup.GET("/withdrawAddress",user.GetWithdrawAddress)         //获取提现钱包
		yunGroup.POST("/updateWithdrawAddress",user.UpdateWithdrawAddress)         //修改提现钱包

		yunGroup.POST("/changePayPasswordByCode",user.ChangePayPasswordByCode)         // 修改支付密码根据验证码
		yunGroup.POST("/changePayPasswordByPassword",user.ChangePayPasswordByPassword) // 修改支付密码根据登录密码
		yunGroup.GET("/getCustomerPhone", user.GetCustomerPhone)                       // 获取用户手机号
		yunGroup.GET("/getCustomerCode", user.GetCustomerCode)                       // 获取用户邀请码


		yunGroup.POST("/add", func(ctx *gin.Context) {
			ctx.JSON(200,result.Ok(""))
		})

	}

	//无权限接口
	//登录
	r.POST("/dologin", login.DoLogin)

	r.POST("/verificationCode", login.VerificationCode) //发送验证码
	//用户注册
	r.POST("/customerRegister",  login.CustomerRegister)		// 注册
	r.POST("/customerLoginByCode", login.CustomerLoginByCode)			//	验证码登录
	r.POST("/customerLoginByPassword", login.CustomerLoginByPassword)	//  账号密码登录
	r.POST("/changePassword", login.ChangePassword)  				// 修改密码

	r.GET("/getNews", news.GetNews)
	r.GET("/getNewsContent", news.GetNewsContent)

	//swagger文档
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//log.Printf("swagger : %s", "http://localhost:3000/swagger/index.html")
	srv := &http.Server{
		Addr:    ":"+pro.Host,
		Handler: r,
	}
	sign := make(chan os.Signal)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM,syscall.SIGHUP, syscall.SIGQUIT)
	//go func() {
	//	<-sign
	//	log.Println("Shutdown............ ...")
	//	if err := srv.Shutdown(context.Background()); err != nil {
	//		log.Fatal("Server Shutdown:", err)
	//	}
	//}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	log.Println("Release connection resources")
	client.Close()
	engine.Close()
	log.Println("Server exiting")

}