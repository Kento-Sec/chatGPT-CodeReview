package helper

import (
	"fmt"
	"os"
)

func CreatConfig() {
	fmt.Println("找不到配置文件...正在创建配置文件...")
	setConfigPath, _ := os.UserHomeDir()
	getConfigPathName := setConfigPath[0:len(setConfigPath)]
	configPath := getConfigPathName + "/.config/chatGPT-CodeReview"
	configPathName := getConfigPathName + "/.config/chatGPT-CodeReview/config.json"
	err := os.MkdirAll(configPath, 0777)
	saveResultName := getConfigPathName + "/chatGPT-CodeReview"
	err2 := os.MkdirAll(saveResultName, 0777)
	if err != nil {
		fmt.Println("创建目录失败")
	}
	if err2 != nil {
		fmt.Println("创建目录失败")
	}
	f, err := os.Create(configPathName)
	if err != nil {
		fmt.Println(err.Error())
	}
	if err != nil {
	} else {
		_, err = f.Write([]byte("{" + "\n"))
		_, err = f.Write([]byte("\t\"apikey\": {" + "\n"))
		_, err = f.Write([]byte("\t\t\"1\": \"\"" + "\n"))
		_, err = f.Write([]byte("\t}" + "\n"))
		_, err = f.Write([]byte("}" + "\n"))
		fmt.Println("配置文件创建好，在配置文件中输入你的api-key")
		os.Exit(0)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
		}
	}(f)

}
