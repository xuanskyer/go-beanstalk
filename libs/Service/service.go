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
	} else {
		fmt.Println("json err:", err)
	}

	fmt.Println("sleep 5 seconds...")
	time.Sleep(5 * time.Second)
}
