package utils

import "github.com/manifoldco/promptui"

func GetStringPromt(label string) (string, error) {
	commonNamePrompt := promptui.Prompt{
		Label: label,
	}

	return commonNamePrompt.Run()
}

func GetSelectPromt(label string, items []string) (int, string, error) {
	commonNamePrompt := promptui.Select{
		Label: label,
		Items: items,
	}

	return commonNamePrompt.Run()
}
