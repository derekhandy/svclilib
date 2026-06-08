package svclilib

import "fmt"

func PRINTH(env Environment) {
	fmt.Print(env.Spacing)
	if env.Header != "" {
		fmt.Print(env.Header)
		fmt.Print(env.Spacing)
	}
	fmt.Print(env.Spacing)
}

func PRINTM(env Environment, text string) {
	fmt.Println(env.Prefix + " " + text)
}

func PRINTI(env Environment, text string) {
	fmt.Println(env.Prefix + " " + text)
	fmt.Println(env.Prefix + " ")
}

func PRINTF(env Environment) {
	fmt.Print(env.Spacing)
	fmt.Print(env.Footer)
	fmt.Print(env.Spacing)
}
