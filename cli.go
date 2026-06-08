package svclilib

import "strings"

type Environment struct {
	Header     string
	Prefix     string
	Footer     string
	Spacing    string
	Commands   []Command
	LogSuccess bool
}

type Command struct {
	Name        string
	ArgRequired int
	Function    func(args []string) error
}

func CMD(env Environment, args []string) {
	PRINTH(env)

	RunCommand(env, args)

	PRINTF(env)
}

func RunCommand(env Environment, args []string) {
	if len(args) == 0 {
		PRINTM(env, "Error: No command provided")
		return
	}

	PRINTM(env, strings.Join(args))

	for _, command := range env.Commands {
		if command.Name == args[0] {
			if len(args) < command.ArgRequired+1 {
				PRINTM(env, "Error: Not enough arguments for command "+command.Name)
			} else if err := command.Function(args[1:]); err != nil {
				PRINTM(env, "Error: "+err.Error())
			} else if env.LogSuccess {
				PRINTM(env, "Command "+command.Name+" executed successfully")
			}
			return
		}
	}

	PRINTM(env, "Error: Unknown command "+args[0])
}
