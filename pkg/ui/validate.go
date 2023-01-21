package ui

import (
	"errors"
	"strings"
)

var NotEmpty = func(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("输入不能为空")
	}
	return nil
}
