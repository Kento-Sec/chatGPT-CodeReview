package main

import (
	"codereviwe/helper"
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	var root string
	flag.StringVar(&root, "f", "", "specify the root folder path")
	flag.Parse()

	if root == "" {
		fmt.Println("Please specify the root folder path using -f flag")
		os.Exit(1)
	}
	err := filepath.Walk(root, visit)
	if err != nil {
		fmt.Println("error:", err)
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		fmt.Println("directory:", path)
	} else {
		setConfigPath, _ := os.UserHomeDir()
		getConfigPathName := setConfigPath[0:len(setConfigPath)]
		configPathName := getConfigPathName + "/.config/chatGPT-CodeReview/"
		viper.AddConfigPath(configPathName)
		viper.AddConfigPath("$HOME/")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
		err := viper.ReadInConfig()
		if err != nil {
			helper.CreatConfig()
			os.Exit(0)
		}
		apiKey := fmt.Sprintf("%s", viper.Get("apikey.1"))
		var record [][]string
		csvFileTitle := [][]string{
			{"Code", "Result"},
		}
		record = append(record, csvFileTitle...)
		fileName := path + ".csv"
		csvdatafile, err := os.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		writer := csv.NewWriter(csvdatafile)
		defer csvdatafile.Close()
		contents, err := ioutil.ReadFile(path)
		fmt.Println(string(contents))
		if err != nil {
			fmt.Println(err)
			return nil
		}
		responseData := GetResponse(apiKey, string(contents))
		if len(responseData) > 0 {
			data := [][]string{
				{
					string(contents),
					responseData,
				},
			}
			fmt.Println(responseData)
			record = append(record, data...)
		} else {
			fmt.Println("No choices available in response")
		}
		err = writer.WriteAll(record)
		if err != nil {
			return err
		}
		writer.Flush()
	}
	return nil
}

func GetResponse(apiKey string, contents string) string {
	client := openai.NewClient(apiKey)
	prompt := "使用中文分析一下这段代码存在什么漏洞：\n" + contents
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
	}
	if len(resp.Choices) > 0 {
		return resp.Choices[0].Message.Content
	}
	return ""
}
