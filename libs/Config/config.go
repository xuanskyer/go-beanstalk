package Config

/**
* 调度变量定义
 */

type TaskArgs struct {
	Member_id   interface{}   	`json:"member_id"`
	Action      string        	`json:"action"`
	Domain_id   interface{}   	`json:"domain_id"`
	Domain_type string        	`json:"domain_type"`
	Type        string        	`json:"type"`
	Source      string        	`json:"source"`
	Callback    []string      	`json:"callback"`
}

type TaskStruct struct {
	Action     string        	`json:"action"`
	Queue_name string        	`json:"queue_name"`
	Args       TaskArgs   `json:"args"`
}
