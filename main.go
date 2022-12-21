package main

import "fmt"

func main() {
	fmt.Println("本项目用来模拟 DevStream 项目的交互式参数设置过程")
	fmt.Println("本项目基于 quickstart 进行模拟，增加了 repo scaffolding 的设置")
	owner, err := setPrompt("输入你的代码托管平台用户名")
	if err != nil {
		panic(err)
	}
	repo, err := setPrompt("输入你的代码仓库名")
	if err != nil {
		panic(err)
	}
	token, err := setSecret("请输入你的代码托管平台的 token")
	if err != nil {
		panic(err)
	}
	repoScaffoldingURL, err := SetRepoScaffolding()
	if err != nil {
		panic(err)
	}
	dockerhubUser, err := setPrompt("请输入你的 DockerHub 用户名")
	if err != nil {
		panic(err)
	}
	dockerhubToken, err := setSecret("请输入你的 DockerHub token")
	if err != nil {
		panic(err)
	}
	fmt.Printf("owner: %s, repo: %s\n", owner, repo)
	fmt.Printf("GitHub token: %s\n", token)
	if repoScaffoldingURL != "" {
		fmt.Printf("repo scaffolding: %s\n", repoScaffoldingURL)
	} else {
		fmt.Println("no repo scaffolding")
	}
	fmt.Printf("dockerhubUser: %s\n", dockerhubUser)
	fmt.Printf("dockerhubToken: %s\n", dockerhubToken)
}
