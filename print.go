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

// func Logr(env Environment, out string) {
// 	if out == "" {
// 		return
// 	}
// 	normalized := strings.ReplaceAll(out, "\r\n", "\n")
// 	log := strings.Split(normalized, "\n")

// 	for i := 0; i < len(log); i++ {
// 		fmt.Println(env.Glyphs[1] + " " + log[0])
// 	}
// }

func Logu(env Environment, cmds []Command) {
	format := env.UsageFormat

	for i := 0; i < len(cmds); i++ {
		usage := format
		usage = strings.ReplaceAll(usage, "{name}", cmds[i].Name)
		usage = strings.ReplaceAll(usage, "{desc}", cmds[i].Desc)
		usage = strings.ReplaceAll(usage, "{prefix}", env.Prefix)
		usage = strings.ReplaceAll(usage, "{spacing}", env.Spacing)
	}
}

func Logf(env Environment) {
	fmt.Print(env.Spacing)
	fmt.Print(env.Footer)
	fmt.Print(env.Spacing)
}
