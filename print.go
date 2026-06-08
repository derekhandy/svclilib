package svclilib

import (
	"fmt"
	"strings"
)

func Logh(env Environment) {
	fmt.Print(env.Spacing)
	if env.Header != "" {
		fmt.Print(env.Header)
		fmt.Print(env.Spacing)
	}
	fmt.Print(env.Spacing)
}

func Logm(env Environment, text string) {
	fmt.Println(env.Prefix + " " + text)
}

func Logc(env Environment, args []string) {
	fmt.Println(env.Glyphs[0] + " " + strings.Join(args, ", "))
}

func Logr(env Environment, output string) {
	normalized := strings.ReplaceAll(output, "\r\n", "\n")
	Log := strings.Split(normalized, "\n")

	for i := 0; i < len(Log); i++ {
		fmt.Println(env.Glyphs[1] + " " + Log[0])
	}
}

func Logf(env Environment) {
	fmt.Print(env.Spacing)
	fmt.Print(env.Footer)
	fmt.Print(env.Spacing)
}
