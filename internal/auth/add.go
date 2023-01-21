package auth

import (
	"fmt"

	"interact_cli_test/pkg/ui"
)

func Add() error {
	fmt.Println("下面开始添加SCM库的认证信息")
	ui.Select("请选择SCM库类型", []string{"GitHub", "GitLab"})
	ui.Select("请选择SCM身份类型", []string{"个人", "组织"})
	ui.SetPrompt("请输入SCM的个人/组织id")
	ui.SetSecret("请输入SCM库的token")
	fmt.Println("认证信息添加成功")
	return nil
}
