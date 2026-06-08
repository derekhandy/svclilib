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

func Logr(env Environment, out string) {
	if out == "" {
		return
	}
	normalized := strings.ReplaceAll(out, "\r\n", "\n")
	log := strings.Split(normalized, "\n")

	for i := 0; i < len(log); i++ {
		fmt.Println(env.Glyphs[1] + " " + log[0])
	}
}

func Logu(env Environment, cmds []Command) {
	format := env.UsageFormat

	for i := 0; i < len(cmds); i++ {
		//basic test
		usage := format
		prevUsage := usage
		usage = strings.ReplaceAll(prevUsage, "{name}", cmds[i].Name)
		prevUsage = usage
		usage = strings.ReplaceAll(prevUsage, "{desc}", cmds[i].Desc)
		prevUsage = usage
		usage = strings.ReplaceAll(prevUsage, "{prefix}", env.Prefix)
		prevUsage = usage
		usage = strings.ReplaceAll(prevUsage, "{spacing}", env.Spacing)
		prevUsage = usage
	}
}

func Logf(env Environment) {
	fmt.Print(env.Spacing)
	fmt.Print(env.Footer)
	fmt.Print(env.Spacing)
}
