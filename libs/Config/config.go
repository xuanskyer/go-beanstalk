package Config

import (
	"gopkg.in/ini.v1"
	"fmt"
	"strings"
)

/**
* 调度变量定义
 */

type TaskArgs struct {
	Type     string            `json:"type"`
	Source   string            `json:"source"`
	Callback []string        `json:"callback"`
}

type TaskStruct struct {
	Action     string            `json:"action"`
	Queue_name string            `json:"queue_name"`
	Args       TaskArgs   `json:"args"`
}



func Get(full_name string) interface{} {

	split_name := strings.Split(full_name, ".")

	count := len(split_name)
	fmt.Println(count)
	if count <= 0 {
		panic("读取配置参数不能为空!")
	}
	ini_file := "conf/" + split_name[0] + ".ini"

	cfg, err := ini.Load(ini_file)
	if err != nil {
		fmt.Println("读取配置文件错误：", err)
	}
	if 1 == count {
		sections := cfg.Sections()
		return sections
	} else if 2 == count {
		section := cfg.Section(split_name[1]).KeysHash()
		return section
	} else if 3 == count {
		key := cfg.Section(split_name[1]).Key(split_name[2]).Value()
		return key
	} else {
		panic("读取配置参数层级太多！")
	}

}
