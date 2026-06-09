package svclilib

import (
	"fmt"
	"strings"
)

func Logh(env Environment) {
	if env.Header != "" {
		fmt.Print(env.Header)
		fmt.Print(env.Spacing)
	}
}

func Logm(env Environment, text string) {
	fmt.Println(env.Prefix + text)
}

func Logc(env Environment, args []string) {
	fmt.Println(env.Glyphs[0] + strings.Join(args, " "))
}

func Logv(env Environment) {
	fmt.Print(env.Glyphs[1])
}

func Logr(env Environment, out string) {
	if out == "" {
		return
	}
	normalized := strings.ReplaceAll(out, "\r\n", "\n")
	log := strings.Split(normalized, "\n")

	for i := 0; i < len(log); i++ {
		if log[i] == "" {
			continue
		}
		fmt.Println(env.Glyphs[2] + log[i])
	}
}

func Logu(env Environment, cmds []Command) {
	format := env.UsageFormat

	for i := 0; i < len(cmds); i++ {
		usage := format
		usage = strings.ReplaceAll(usage, "{name}", cmds[i].Name)
		usage = strings.ReplaceAll(usage, "{desc}", cmds[i].Desc)
		usage = strings.ReplaceAll(usage, "{usage}", cmds[i].Usage)
		usage = strings.ReplaceAll(usage, "{prefix}", env.Prefix)
		usage = strings.ReplaceAll(usage, "{spacing}", env.Spacing)
		fmt.Println(usage)
	}
}

func Logf(env Environment) {
	fmt.Print(env.Footer)
}
