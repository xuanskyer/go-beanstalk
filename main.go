package main

/**
 * beanstalk 队列处理入口
 */
import (
	"runtime"
	"go-beanstalk/libs/Logger"
	"go-beanstalk/libs/Queue"
	"go-beanstalk/libs/Db"
)

/**
 * 入口方法
 * 支持多队列多任务并发处理
 */
func main() {

	Logger.Notice(" ============== 队列任务处理脚本启动 ============== ")

	Db.Test()


	runtime.GOMAXPROCS(runtime.NumCPU())

	go Queue.Consumer("CA", Queue.TubeName1)
	go Queue.Consumer("CB", Queue.TubeName2)

	//无意义处理，保证脚本不退出
	Queue.Consumer("Consumer", Queue.TubeNameEnd)

	Logger.Notice("============== 队列任务处理脚本退出 ============== ")
}
