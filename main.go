package main

import (
	ds "WF_SG/Chaincode/DataStructure"
	"WF_SG/Chaincode/Utils"
	"WF_SG/SDKInit"
	"WF_SG/Services"
	"WF_SG/Web/common"
	"WF_SG/Web/models"
	"WF_SG/Web/routes"
	"encoding/json"
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	config "github.com/spf13/viper"
	"log"
	"strconv"
	"time"
)

func InitDb() {
	//bids sig
	var bids []models.BidModel
	db := common.DB
	db.Find(&bids)
	for i := 0; i < len(bids); i++ {

		if bids[i].ContractCompanyOwnerSig != "" {
			continue
		}

		bigJsonBeforeSig, _ := json.Marshal(bids[i])
		bidSig, err := Utils.Sign(bigJsonBeforeSig, bids[i].ContractCompanyOwnerAccount)
		if err != nil {
			fmt.Println(err.Error())
		}
		bids[i].ContractCompanyOwnerSig = bidSig
		db.Save(&bids[i])
	}
	println("bids sig initialized")

	//user keys
	var keypair Utils.KeyPair

	_ = json.Unmarshal([]byte("{\"Skpem\":\"LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ0F1UzduSHBLUTZveC8xdUwKY0NXbU1MYmVkZmNDNUJIVk9WcEVpZ0ZCQ1graFJBTkNBQVI2Q2FqVk1xSGlkSnFPU2dqbUpOM0VoSEgvT0tqWgp3NzMvVmJGS1pZbjhYcW1JclNMQi9qRzdpOWs1QlpMYk1HenFzSVVpUmxBWjdEL0J5VjYvSGRJegotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg==\",\"Pkpem\":\"LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNIRENDQWNLZ0F3SUJBZ0lSQUxJcjNBSlVHVi9PSnhOSVozdXNQTDh3Q2dZSUtvWkl6ajBFQXdJd2FURUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMkoxYVd4a1pYSXVZMjl0TVJjd0ZRWURWUVFERXc1allTNWlkV2xzClpHVnlMbU52YlRBZUZ3MHlNREF4TVRBd09EUTVNREJhRncwek1EQXhNRGN3T0RRNU1EQmFNR2N4Q3pBSkJnTlYKQkFZVEFsVlRNUk13RVFZRFZRUUlFd3BEWVd4cFptOXlibWxoTVJZd0ZBWURWUVFIRXcxVFlXNGdSbkpoYm1OcApjMk52TVE4d0RRWURWUVFMRXdaamJHbGxiblF4R2pBWUJnTlZCQU1NRVVGa2JXbHVRR0oxYVd4a1pYSXVZMjl0Ck1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWVnbW8xVEtoNG5TYWprb0k1aVRkeElSeC96aW8KMmNPOS8xV3hTbVdKL0Y2cGlLMGl3ZjR4dTR2Wk9RV1MyekJzNnJDRklrWlFHZXcvd2NsZXZ4M1NNNk5OTUVzdwpEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3S3dZRFZSMGpCQ1F3SW9BZ2dUbmpkVU9nCldMNVcxNDNjRGlCVnNQNndVZldxSVN6NG04MzlERGVxcXZBd0NnWUlLb1pJemowRUF3SURTQUF3UlFJaEFNb2QKR2pnN0c2UFZjNEMvRnQyMWtHajVWRUVPWUNIeXJacVBNN2xTdlVidkFpQVB4QU1pcWUwVjgrTjkrRjduWjI2bAppOTJNU05lUHltRlNiSENxd3IxL05BPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\"}"), &keypair)

	/*
		keypair = Utils.KeyPair{
			Skpem: []byte("LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0tCk1JR0hBZ0VBTUJNR0J5cUdTTTQ5QWdFR0NDcUdTTTQ5QXdFSEJHMHdhd0lCQVFRZ0F1UzduSHBLUTZveC8xdUwKY0NXbU1MYmVkZmNDNUJIVk9WcEVpZ0ZCQ1graFJBTkNBQVI2Q2FqVk1xSGlkSnFPU2dqbUpOM0VoSEgvT0tqWgp3NzMvVmJGS1pZbjhYcW1JclNMQi9qRzdpOWs1QlpMYk1HenFzSVVpUmxBWjdEL0J5VjYvSGRJegotLS0tLUVORCBQUklWQVRFIEtFWS0tLS0tCg=="),
			Pkpem: []byte("LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNIRENDQWNLZ0F3SUJBZ0lSQUxJcjNBSlVHVi9PSnhOSVozdXNQTDh3Q2dZSUtvWkl6ajBFQXdJd2FURUwKTUFrR0ExVUVCaE1DVlZNeEV6QVJCZ05WQkFnVENrTmhiR2xtYjNKdWFXRXhGakFVQmdOVkJBY1REVk5oYmlCRwpjbUZ1WTJselkyOHhGREFTQmdOVkJBb1RDMkoxYVd4a1pYSXVZMjl0TVJjd0ZRWURWUVFERXc1allTNWlkV2xzClpHVnlMbU52YlRBZUZ3MHlNREF4TVRBd09EUTVNREJhRncwek1EQXhNRGN3T0RRNU1EQmFNR2N4Q3pBSkJnTlYKQkFZVEFsVlRNUk13RVFZRFZRUUlFd3BEWVd4cFptOXlibWxoTVJZd0ZBWURWUVFIRXcxVFlXNGdSbkpoYm1OcApjMk52TVE4d0RRWURWUVFMRXdaamJHbGxiblF4R2pBWUJnTlZCQU1NRVVGa2JXbHVRR0oxYVd4a1pYSXVZMjl0Ck1Ga3dFd1lIS29aSXpqMENBUVlJS29aSXpqMERBUWNEUWdBRWVnbW8xVEtoNG5TYWprb0k1aVRkeElSeC96aW8KMmNPOS8xV3hTbVdKL0Y2cGlLMGl3ZjR4dTR2Wk9RV1MyekJzNnJDRklrWlFHZXcvd2NsZXZ4M1NNNk5OTUVzdwpEZ1lEVlIwUEFRSC9CQVFEQWdlQU1Bd0dBMVVkRXdFQi93UUNNQUF3S3dZRFZSMGpCQ1F3SW9BZ2dUbmpkVU9nCldMNVcxNDNjRGlCVnNQNndVZldxSVN6NG04MzlERGVxcXZBd0NnWUlLb1pJemowRUF3SURTQUF3UlFJaEFNb2QKR2pnN0c2UFZjNEMvRnQyMWtHajVWRUVPWUNIeXJacVBNN2xTdlVidkFpQVB4QU1pcWUwVjgrTjkrRjduWjI2bAppOTJNU05lUHltRlNiSENxd3IxL05BPT0KLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo="),
		}
	*/
	keypair.Sk = Utils.GetPriKey(keypair.Skpem)
	keypair.Pk = Utils.GetPubKey(keypair.Pkpem)

	var users []models.UserModel
	db.Find(&users)
	for i := 0; i < len(users); i++ {
		Utils.KeyMap[users[i].Account] = &keypair
	}
	println("users keypair initialized")

	//init some ieds
	var ied1 = ds.IedModel{
		DeviceId:           "1",
		DeviceName:         "Intelligent Device Generation-2",
		DeviceProducer:     "Smart Bridge Company",
		DeviceWorkingDays:  0,
		DeviceBelongIem:    "Lakeview Community",
		DeviceUserAccount:  users[0].Account,
		DeviceDownInfos:    nil,
		DeviceWorkingInfos: nil,
	}
	iedAsJsonBytes1, _ := json.Marshal(ied1)
	Services.HLservice.AddIedService(iedAsJsonBytes1)

	//init some contract
	var contract1 = ds.ContractModel{
		ContractId:                  "1231",
		ContractVersion:             "1.0",
		ContractName:                "Smart Water",
		ContractCompanyName:         "SanFrancisco Power",
		ContractCompanyOwnerAccount: users[0].Account,
		ContractCompanyOwnerSig:     "",
		ContractDetails:             "saeoijufhnoquawehjnfoiq",
		EnergyType:                  "water",
		EnergyPrice:                 "1.39",
		ContractLastTime:            "2020.01.12",
		ContractSignTime:            "",
		ContractUserAccount:         users[0].Account,
		ContractUserSig:             "",
	}

	contractAsJsonBytes1, _ := json.Marshal(contract1)
	var bid ds.BidModel
	_ = json.Unmarshal(contractAsJsonBytes1, &bid)
	bidAsJsonBytes, _ := json.Marshal(bid)

	//test use private key to sign
	signature1, _ := Utils.Sign(bidAsJsonBytes, bid.ContractCompanyOwnerAccount)
	contract1.ContractCompanyOwnerSig = signature1
	contractAsJsonBytes1, _ = json.Marshal(contract1)

	signature2, _ := Utils.Sign(contractAsJsonBytes1, contract1.ContractUserAccount)
	contract1.ContractUserSig = signature2
	contractAsJsonBytes1, _ = json.Marshal(contract1)

	Services.HLservice.AddContractService(contractAsJsonBytes1)
}

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

	SDKInit.ChannelIdToConfigs = map[string]string{
		"hustgym":      "/channel-artifacts/HUSTgym.tx",
		"hustdomitory": "/channel-artifacts/HUSTdomitory.tx",
	}

	err := SDKInit.SetupInitInfo("")
	if err != nil {
		fmt.Println("Failed in sdk setup" + err.Error())
	}

	InitDb()
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
			userName := ds.AiteBefore(userInfo.Account)
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
