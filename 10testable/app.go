package main

import "github.com/op/go-logging"

/*
@author RandySun
@create 2022-05-03-13:38
*/

//var log = logrus.New()
//
//type App struct{}
//
//func (a *App) Start() {
//	log.Info("app start ...")
//}

//func (a *app) Start() {
//	a.Logger.Info("app start ...")
//
//	// ...
//}

type App struct {
	Logger
}

func (a *App) Start() {
	a.Logger.Info("app start ...")
	// ...
}

// NewApp 构造函数，将依赖项注入
func NewApp(lg Logger) *App {
	return &App{
		Logger: lg, // 使用传入的依赖项完成初始化
	}
}

// Logger 将日志库抽象为接口类型
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
}

func main() {

	//app := &App{}
	//app.Start()

	var log = logging.MustGetLogger("example")
	app := NewApp(log)
	app.Start()

}
