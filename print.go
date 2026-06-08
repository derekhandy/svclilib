package svclilib

import "fmt"

func PRINT(env Environment, text string) {
	fmt.Print(env.Spacing)
	if env.Header != "" {
		fmt.Print(env.Header)
		fmt.Print(env.Spacing)
	}
	fmt.Print(env.Prefix + " " + text)
	fmt.Print(env.Spacing)
}

func PRINTM(env Environment, text string, hasInput bool) {
	fmt.Print(env.Spacing)
	fmt.Print(env.Prefix + " " + text)
	if hasInput {
		fmt.Print(env.Spacing)
		fmt.Print(env.Prefix + " ")
	}
}
