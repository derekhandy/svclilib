package svclilib

import (
	"fmt"
	"strings"
)

var defaultCommand = Command{
	Name:  "Unknown Name",
	Desc:  "No description found.",
	Usage: "No usage found.",
}

// Summary:
//
//	Log Command
func Logc(env Environment, args []string) {
	fmt.Println(env.Glyphs[0] + strings.Join(args, " "))
}

// Summary:
//
//	Log Footer
func Logf(env Environment) {
	if env.Footer != "" {
		fmt.Print(env.Footer)
	}
}

// Summary:
//
//	Log Header
func Logh(env Environment) {
	if env.Header != "" {
		fmt.Print(env.Header)
		fmt.Print(env.Spacing)
	}
}

// Summary:
//
//	Log Message
func Logm(env Environment, text string) {
	fmt.Println(env.Prefix + text)
}

// Summary:
//
//	Log Output
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

// Summary:
//
//	Log Usage
func Logu(env Environment, cmds []Command) {
	format := env.UsageFormat

	for i := 0; i < len(cmds); i++ {
		usage := format
		usage = replaceOrDefault(usage, "{name}", cmds[i].Name)
		usage = replaceOrDefault(usage, "{desc}", cmds[i].Desc)
		usage = replaceOrDefault(usage, "{usage}", cmds[i].Usage)
		usage = strings.ReplaceAll(usage, "{prefix}", env.Prefix)
		usage = strings.ReplaceAll(usage, "{spacing}", env.Spacing)
		fmt.Println(usage)
	}
}

// Summary:
//
//	Log Line Spacing
func Logv(env Environment) {
	fmt.Print(env.Glyphs[1])
}

// Summary:
//
//	Replace parameter in string
func replaceOrDefault(str string, match string, replace string) string {
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
