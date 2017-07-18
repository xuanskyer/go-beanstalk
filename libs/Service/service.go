package Service

import (
	"fmt"
	"reflect"
	"encoding/json"
	"time"
	"go-beanstalk/libs/Config"
)

/**
 * 任务处理逻辑
 *
 */
func HandleTask(content string) {
	fmt.Println(reflect.TypeOf(content), reflect.ValueOf(content).Kind())
	fmt.Println(content)
	var task_content Config.TaskStruct
	if err := json.Unmarshal([]byte(content), &task_content); err == nil {
		fmt.Println("------------- task_content转struct -------------")
		fmt.Println(task_content)
		fmt.Println("action", task_content.Action)
		fmt.Println("queue_name", task_content.Queue_name)
		fmt.Println("args", task_content.Args)
		fmt.Println("args.member_id", task_content.Args.Member_id)
		fmt.Println("args.domain_id", task_content.Args.Domain_id)
		fmt.Println("args.domain_type", task_content.Args.Domain_type)
	} else {
		fmt.Println("json err:", err)
	}

	fmt.Println("sleep 5 seconds...")
	time.Sleep(5 * time.Second)
}
