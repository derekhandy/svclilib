package svclilib

import (
	"fmt"
	"strings"
)

func Printh(env Environment) {
	fmt.Print(env.Spacing)
	if env.Header != "" {
		fmt.Print(env.Header)
		fmt.Print(env.Spacing)
	}
	fmt.Print(env.Spacing)
}

func Printm(env Environment, text string) {
	fmt.Println(env.Prefix + " " + text)
}

func Printc(env Environment, args []string) {
	fmt.Println(env.Glyphs[0] + " " + strings.Join(args, ", "))
}

func Printu(env Environment, cmds []Command) {
	// format := env.UsageFormat

	// for i := 0; i < len(cmds); i++ {
	// 	usage := format
	// 	usage = strings.ReplaceAll(usage, "{name}", cmds[i].Name)
	// 	usage = strings.ReplaceAll(usage, "{desc}", cmds[i].Desc)
	// 	usage = strings.ReplaceAll(usage, "{prefix}", env.Prefix)
	// 	usage = strings.ReplaceAll(usage, "{spacing}", env.Spacing)
	// }
}

func Printf(env Environment) {
	fmt.Print(env.Spacing)
	fmt.Print(env.Footer)
	fmt.Print(env.Spacing)
}
