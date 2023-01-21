package ui

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SetPrompt(label string) (string, error) {
	return SetPromptWithDefault(label, "")
}

func SetPromptWithDefault(label string, defaults string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: NotEmpty,
		Default:  defaults,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func SetSecret(label string) (string, error) {
	prompt := promptui.Prompt{
		Label:    label,
		Mask:     '*',
		Validate: NotEmpty,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}
