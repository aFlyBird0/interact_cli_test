package main

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
)

func SetRepoScaffolding() (url string, err error) {
	useRepoScaffolding, err := askIfUseRepoScaffolding()
	if err != nil {
		return
	}

	if useRepoScaffolding {
		url, err = selectRepoScaffolding()
		if err != nil {
			return
		}
	}

	return
}

func askIfUseRepoScaffolding() (bool, error) {
	return Confirm("是否使用仓库模板？", true)
}

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

type RepoScaffolding struct {
	Name        string
	URL         string
	Language    string
	Framework   string
	Description string
}

func selectRepoScaffolding() (url string, err error) {

	repos := []RepoScaffolding{
		{Name: "dtm-repo-scaffolding-golang-gin",
			URL:         "https://github.com/devstream-io/dtm-repo-scaffolding-golang-gin",
			Language:    "Golang",
			Framework:   "Gin",
			Description: "这是一个Golang Web框架的仓库模板，基于Gin框架，拥有。。。。的功能",
		},
		{Name: "dtm-repo-scaffolding-golang-cli",
			URL:         "https://github.com/devstream-io/dtm-repo-scaffolding-golang-cli",
			Language:    "Golang",
			Framework:   "Cobra",
			Description: "这是一个Golang CLI框架的仓库模板，基于Cobra框架，拥有。。。。的功能",
		},
		{Name: "dtm-repo-scaffolding-python-flask",
			URL:         "https://github.com/devstream-io/dtm-repo-scaffolding-python-flask",
			Language:    "Python",
			Framework:   "Flask",
			Description: "这是一个Python Web框架的仓库模板，基于Flask框架，拥有。。。。的功能",
		},
		{Name: "dtm-repo-scaffolding-java-springboot",
			URL:         "https://github.com/devstream-io/dtm-repo-scaffolding-java-springboot",
			Language:    "Java",
			Framework:   "SpringBoot",
			Description: "这是一个Java Web框架的仓库模板，基于SpringBoot框架，拥有。。。。的功能",
		},
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "🐰 {{ .Name | blue }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "🐰 仓库模板: {{ .Name | blue | cyan }}",
		Details: `
--------- Repo Details ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Language:" | faint }}	{{ .Language }}
{{ "Framework:" | faint }}	{{ .Framework }}
{{ "Description:" | faint }}	{{ .Description }}`,
	}

	searcher := func(input string, index int) bool {
		pepper := repos[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "请选择代码模板",
		Items:     repos,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	return repos[i].URL, nil
}
