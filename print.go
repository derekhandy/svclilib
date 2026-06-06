package svclilib

import "fmt"

func PRINT(env Environment, text string) {
	fmt.Println()
	if env.Header != "" {
		fmt.Println(env.Header)
		fmt.Println()
	}
	fmt.Println(env.Prefix + " " + text)
}

func PRINTM(env Environment, text string, hasInput bool) {
	fmt.Println(env.Prefix + " " + text)
	if hasInput {
		fmt.Print(env.Prefix + " ")
	}
}
