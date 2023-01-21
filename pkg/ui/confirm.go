package ui

import (
	"errors"
	"strings"

	"github.com/manifoldco/promptui"
)

// 官方的confirm有点问题，这里重新封个
// 代码从 https://github.com/manifoldco/promptui/issues/81#issuecomment-906597371 修改而来
func Confirm(label string, defaultSelect bool) (bool, error) {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	if defaultSelect {
		prompt.Default = "y"
	} else {
		prompt.Default = "n"
	}
	validate := func(s string) error {
		if len(s) == 1 && strings.Contains("YyNn", s) || prompt.Default != "" && len(s) == 0 {
			return nil
		}
		return errors.New("非法输入，请输入Y,y,N,n 中的一个")
	}
	prompt.Validate = validate

	_, err := prompt.Run()
	confirmed := !errors.Is(err, promptui.ErrAbort)
	if err != nil && confirmed {
		return false, err
	}

	return confirmed, nil
}
