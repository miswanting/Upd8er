package main

import (
	"fmt"
	"os"
	// "time"
	"strconv"
)

var (
	version        = "0.1.0-190325"
	configFileName = "Upconfig.ini"
	language       = "zh-cn"
)

func main() {
	fmt.Println("Upd8er", "v"+version, "\n")
	if _, err := os.Stat(configFileName); os.IsNotExist(err) { // 检测是否存在配置文件
		setup() // 没有，启动配置向导
	} else {
		readConfig() // 有，读取配置文件
	}
}
func setup() { // 配置向导
	// time.Sleep(1 * time.Second)
	fmt.Println("欢迎使用 Upd8er 自动升级程序！")
	fmt.Println("检测到这是你第一次运行程序，请首先完成初始设置向导。\n")
	fmt.Println("【初始设置向导】\n")
	fmt.Println("您想要更新哪一个程序？")
	fmt.Println(radio([2]string{"Era.js", "EraLife"}))
	// 写入配置文件
}
func readConfig() { // 读取配置文件
	fmt.Println("readConfig")
}
func radio(radioList [2]string) (choice string) {
	for i := 0; i < len(radioList); i++ {
		fmt.Println(strconv.Itoa(i+1)+".", radioList[i])
	}
	fmt.Printf("请输入序号[1-" + strconv.Itoa(len(radioList)) + "]：")
	var choiceS string
	fmt.Scanln(&choiceS)
	choiceI, _ := strconv.ParseInt(choiceS, 10, 0)
	return radioList[int(choiceI-1)]
}
