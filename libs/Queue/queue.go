package Queue

import (
	"fmt"
	"time"
	"strings"
	"github.com/kr/beanstalk"
	"go-beanstalk/libs/Service"
)

/**
 * 调度队列处理
 */
var (
	HostPort    string = "127.0.0.1:11300"
	TubeName1   string = "default"
	TubeName2   string = "default2"
	TubeNameEnd string = "TubeNameEnd"
)

//队列生产
func Producer(ProducerName, tubeName string) {
	if ProducerName == "" || tubeName == "" {
		return
	}

	c, err := beanstalk.Dial("tcp", HostPort)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true
	fmt.Println(ProducerName, " [Producer] tubeName:", tubeName, " c.Tube.Name:", c.Tube.Name)
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("for %s %d", tubeName, i)
		c.Put([]byte(msg), 30, 0, 120*time.Second)
		fmt.Println(ProducerName, " [Producer] beanstalk put body:", msg)
		//time.Sleep(1 * time.Second)
	}

	c.Close()
	fmt.Println("Producer() end.")
}

//队列消费
func Consumer(ConsumerName string, tubeName string) {
	if ConsumerName == "" {
		panic("消费者不能为空！")
		return
	}

	if ConsumerName == "" || tubeName == "" {
		panic("队列名不能为空！")
		return
	}

	c, err := beanstalk.Dial("tcp", HostPort)
	if err != nil {
		panic("初始化队列失败！")
		panic(err)
	}
	defer c.Close()

	c.Tube.Name = tubeName
	c.TubeSet.Name[tubeName] = true

	fmt.Println("ConsumerName: ", ConsumerName, ", tubeName: ", tubeName, ", c.Tube.Name:", c.Tube.Name)

	substr := "timeout"
	for {
		//从队列中取出
		id, body, err := c.Reserve(1 * time.Second)
		if err != nil {
			if !strings.Contains(err.Error(), substr) {
				fmt.Println(ConsumerName, " [Consumer] [", c.Tube.Name, "] err:", err, " id:", id)
			}
			continue
		}
		fmt.Println(ConsumerName, " [Consumer] [", c.Tube.Name, "] job:", id, " body:", string(body))

		//队列任务处理逻辑
		go Service.HandleTask(string(body))

		//从队列中清掉
		err = c.Delete(id)
		if err != nil {
			fmt.Println(ConsumerName, " [Consumer] [", c.Tube.Name, "] Delete err:", err, " id:", id)
		} else {
			fmt.Println(ConsumerName, " [Consumer] [", c.Tube.Name, "] Successfully deleted. id:", id)
		}
	}
	fmt.Println("Consumer() end. ")
}
