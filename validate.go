package main

import (
	"errors"
	"strings"
)

var validate = func(input string) error {
	if strings.TrimSpace(input) == "" {
		return errors.New("输入不能为空")
	}
	return nil
}
