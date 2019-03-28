package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	version        = "0.1.0-190325"
	configFileName = "Upconfig.ini"
	language       = "zh-cn"
	source         map[string]string
)

type Cache struct {
	Name    string
	Channel string
	Version string
}

func main() {
	fmt.Println("Upd8er", "v"+version, "\n")
	fmt.Println("欢迎使用 Upd8er 自动升级程序！")
	if _, err := os.Stat(configFileName); os.IsNotExist(err) { // 检测是否存在配置文件
		noConfig() // 没有，启动配置向导
	} else {
		hasConfig() // 有，读取配置文件
	}
}
func noConfig() { // 配置向导
	time.Sleep(1 * time.Second)
	fmt.Println("检测到这是你第一次运行程序，请首先完成初始设置向导。\n")
	fmt.Println("【初始设置向导】\n")
	fmt.Println("您想要将更新器绑定至哪一个程序？")
	choice := radio([2]string{"Era.js", "EraLife"})
	// 写入配置文件
	cache := Cache{
		Name: choice,
	}
	b, _ := json.MarshalIndent(cache, "", "")
	ioutil.WriteFile(configFileName, b, 0777)
}
func hasConfig() { // 读取配置文件
	time.Sleep(1 * time.Second)
	fmt.Println("正在查询新版本…")
	fmt.Println("您想要将更新器绑定至哪一个程序？")
	b, _ := ioutil.ReadFile(configFileName)
	var cache Cache
	json.Unmarshal(b, &cache)
	downloadToFile("https://raw.githubusercontent.com/miswanting/Era.js/master/README.md", "a.md")
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
func downloadToFile(source string, target string) {
	file, _ := os.Create(target)
	defer file.Close()

	// Get the data
	res, err := http.Get(source)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	// Writer the body to file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println(err)
	}
}
func downloadSource(source string) (data io.ReadCloser) {
	// Get the data
	res, err := http.Get(source)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	var 
	return res.Body
}
