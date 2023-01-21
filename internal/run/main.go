package run

import (
	"fmt"

	"interact_cli_test/pkg/ui"
)

func Execute() error {
	fmt.Println("本项目用来模拟 DevStream 的 rs 子项目的用户交互部分")
	//fmt.Println("（红色）！您还未配置SCM的认证信息，请先执行 auth add 命令添加认证信息")
	ui.Select("请选择目的SCM", []string{"GitHub/aFlyBird0"})
	selectRepoScaffolding()
	ui.SetPromptWithDefault("请输入目的仓库的名称", "dtm-repo-scaffolding-golang-gin")
	return nil
}
