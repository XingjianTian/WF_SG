package main

import (
	"WF_SG/DataStructure"
	"WF_SG/Web/common"
	"WF_SG/Web/models"
	"WF_SG/Web/routes"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func init() {

	//read config
	config.AddConfigPath("Web/configs")
	config.SetConfigName("mysql")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error reading config, %s", err)
	}
	dbConfig := common.DbConfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
		config.GetString("default.charset"),
		config.GetInt("default.MaxIdleConns"),
		config.GetInt("default.MaxOpenConns"),
	}
	common.DB = dbConfig.InitDB()

	if config.GetBool("default.sql_log") {
		common.DB.LogMode(true)
	}

	/*
		SDKInit.ChannelIdToConfigs = map[string]string{
			"hustgym":      "/channel-artifacts/HUSTgym.tx",
			"hustdomitory": "/channel-artifacts/HUSTdomitory.tx",
		}

		err := SDKInit.SetupInitInfo("")
		if err != nil {
			fmt.Println("Failed in sdk setup" + err.Error())
		}
	*/
}
func main() {
	app := iris.New()
	config.SetConfigName("app")
	if err := config.ReadInConfig(); err != nil {
		log.Fatalf("error reading config, %s", err)
	}
	tmpl := iris.HTML("Web/views", ".html").
		Layout(config.GetString("site.DefaultLayout"))
	if config.GetBool("site.APPDebug") == true {
		app.Logger().SetLevel("debug") //设置debug
		tmpl.Reload(true)
	}

	//public func
	tmpl.AddFunc("TimeToDate", common.TimeToDate)
	tmpl.AddFunc("strToHtml", common.StrToHtml)

	//在这里默认layout被注册
	app.RegisterView(tmpl)
	app.Favicon("Web/favicon.ico")
	app.Use(iris.Gzip)

	//（可选）添加两个内置处理程序
	//可以从任何与http相关的panics中恢复
	//并将请求记录到终端。
	app.Use(recover.New())
	app.Use(logger.New())

	app.StaticWeb("/public", "Web/public")   //设置静态文件目录
	app.StaticWeb("/uploads", "Web/uploads") //设置静态文件目录
	app.StaticWeb("/node_modules", "node_modules")

	//设置公共页面输出1
	app.Use(func(ctx iris.Context) {
		//读session
		if acc := common.SessManager.Start(ctx).GetString("user_session"); acc != "" {
			channel := common.SessManager.Start(ctx).GetString("channel")
			var userModel models.UserModel
			//读数据库
			userInfo, _ := userModel.UserInfo(acc)
			if userInfo.Headico == "" {
				//默认头像
				userInfo.Headico = "/public/imgs/user2-160x160.jpg"
			}
			//保存一个或多个键值
			userName := DataStructure.AiteBefore(userInfo.Account)
			ctx.ViewData("userInfo", userInfo)
			ctx.ViewData("userName", userName)
			ctx.ViewData("channel", channel)
		}
		ctx.ViewData("Title", config.GetString("site.DefaultTitle"))
		now := time.Now().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		ctx.ViewData("CurrentTime", now)
		ctx.Next()
	})
	//设置错误模版
	app.OnAnyErrorCode(func(ctx iris.Context) {
		_, err := ctx.HTML("<center>Sorry, Error Code:" + strconv.Itoa(ctx.GetStatusCode()) + "</center>")
		if err != nil {
			log.Fatalf("Sorry, Error %s", err)
		}
	})

	routes.Routes(app)

	//应用配置文件
	app.Configure(iris.WithConfiguration(iris.YAML("Web/configs/iris.yml")))

	//Run
	www := app.Party("www.")
	{
		currentRoutes := app.GetRoutes()
		for _, r := range currentRoutes {
			//Method就是get/post/delete/update，src就是route地址,handlers是方法
			www.Handle(r.Method, r.Tmpl().Src, r.Handlers...)
		}
	}
	err := app.Run(iris.Addr(config.GetString("server.domain") + ":" + config.GetString("server.port")))
	if err != nil {
		log.Fatalf("Failed to start server, Error: %s", err)
	}
}
