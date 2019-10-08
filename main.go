package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"lottery/config"
	"lottery/controllers"
	"lottery/database"
	"lottery/models"
)

/**
 * 初始化 iris app
 * @method NewApp
 * @return  {[type]}  api      *iris.Application  [iris app]
 */
func newApp() (api *iris.Application) {
	api = iris.New()
	api.Use(logger.New())

	api.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.JSON(controllers.ApiResource(false, nil, "404 Not Found"))
	})
	api.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.WriteString("Oups something went wrong, try again")
	})

	//同步模型数据表
	//如果模型表这里没有添加模型，单元测试会报错数据表不存在。
	//因为单元测试结束，会删除数据表
	database.DB.AutoMigrate(
		&models.Lottery{},
	)

	iris.RegisterOnInterrupt(func() {
		database.DB.Close()
	})


	return
}

func main() {

	go controllers.GetLottery(0)

	app := newApp()

	addr := config.Conf.Get("app.addr").(string)
	app.Run(iris.Addr(addr))
}
