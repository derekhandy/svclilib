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

var defaultCommand = Command{
	Name:  "Unknown Name",
	Desc:  "No description found.",
	Usage: "No usage found.",
}

func Logu(env Environment, cmds []Command) {
	format := env.UsageFormat

	for i := 0; i < len(cmds); i++ {
		usage := format
		usage = ReplaceOrDefault(usage, "{name}", cmds[i].Name)
		usage = ReplaceOrDefault(usage, "{desc}", cmds[i].Desc)
		usage = ReplaceOrDefault(usage, "{usage}", cmds[i].Usage)
		usage = strings.ReplaceAll(usage, "{prefix}", env.Prefix)
		usage = strings.ReplaceAll(usage, "{spacing}", env.Spacing)
		fmt.Println(usage)
	}
}

func ReplaceOrDefault(str string, match string, replace string) string {
	if match == "" || str == "" {
		return str
	}

	if replace == "" {
		switch match {
		case "{name}":
			replace = defaultCommand.Name
		case "{desc}":
			replace = defaultCommand.Desc
		case "{usage}":
			replace = defaultCommand.Usage
		}
	}

	return strings.ReplaceAll(str, match, replace)
}

func Logf(env Environment) {
	fmt.Print(env.Footer)
}
