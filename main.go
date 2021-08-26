package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	app, cleanFunc, err := BuildApp()
	if err != nil {
		panic(err)
	}
	router := app.Router.NewRouter()
	go func() {
		_ = router.Run(app.Config.App.Addr)
	}()

EXIT:
	for {
		sig := <-sc
		app.Logger.Infof("接收到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	app.Logger.Infof("服务退出")
	time.Sleep(time.Second)
	os.Exit(state)
}
