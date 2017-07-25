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

func getCommon(full_name string) interface{} {

	split_name := strings.Split(full_name, ".")

	count := len(split_name)
	if count <= 0 {
		panic("读取配置参数不能为空!")
	}
	ini_file := "conf/" + split_name[0] + ".ini"

	cfg, err := ini.Load(ini_file)
	if err != nil {
		fmt.Println("读取配置文件错误：", err)
	}
	if 1 == count {
		cache := map[string]map[string]string{}
		names := cfg.SectionStrings()

		for _, value := range names {
			cache[value] = cfg.Section(value).KeysHash()
		}
		return cache
	} else if 2 == count {
		section := cfg.Section(split_name[1]).Keys()
		return section
	} else if 3 == count {
		key := cfg.Section(split_name[1]).Key(split_name[2]).Value()
		return key
	} else {
		panic("读取配置参数层级太多！")
	}

}

/**
* 获取整个ini文件的配置信息
 */
func GetAll(full_name string) map[string]map[string]string {

	split_name := strings.Split(full_name, ".")

	count := len(split_name)
	if count <= 0 {
		panic("读取配置参数不能为空!")
	}
	ini_file := "conf/" + split_name[0] + ".ini"

	cfg, err := ini.Load(ini_file)
	if err != nil {
		fmt.Println("读取配置文件错误：", err)
	}
	if 1 == count {
		cache := map[string]map[string]string{}
		names := cfg.SectionStrings()

		for _, value := range names {
			cache[value] = cfg.Section(value).KeysHash()
		}
		return cache
	} else {
		panic("读取配置参数层级不合法：Get(\"ini_file\") ")
	}

}

/**
* 获取指定section的配置或指定key的配置
 */
func Get(full_name string) map[string]string {

	split_name := strings.Split(full_name, ".")

	count := len(split_name)
	if count <= 0 {
		panic("读取配置参数不能为空!")
	}
	ini_file := "conf/" + split_name[0] + ".ini"

	cfg, err := ini.Load(ini_file)
	if err != nil {
		fmt.Println("读取配置文件错误：", err)
	}
	if 2 == count {
		section := cfg.Section(split_name[1]).KeysHash()
		return section
	} else if 3 == count {
		cache := map[string]string{}
		cache[split_name[2]] = cfg.Section(split_name[1]).Key(split_name[2]).Value()

		return cache
	} else {
		panic("读取配置参数层级不合法：Get(\"ini_file.section\") 或 Get(\"ini_file.section.key\") ")
	}

}
