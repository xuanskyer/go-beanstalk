package Config

/**
* 调度变量定义
 */

type TaskArgs struct {
	Type        string        	`json:"type"`
	Source      string        	`json:"source"`
	Callback    []string      	`json:"callback"`
}

type TaskStruct struct {
	Action     string        	`json:"action"`
	Queue_name string        	`json:"queue_name"`
	Args       TaskArgs   `json:"args"`
}
