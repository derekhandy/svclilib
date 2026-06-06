package golangcli

type Environment struct {
	Header   string
	Prefix   string
	Commands []Command
}

type Command struct {
	Name        string
	ArgRequired int
	Function    func(args []string) error
}

func CMD(env Environment, args []string) {
	if len(args) == 0 {
		PRINT(env, "Error: No command provided")
		return
	}
	for _, command := range env.Commands {
		if command.Name == args[0] {
			if len(args) < command.ArgRequired+1 {
				PRINT(env, "Error: Not enough arguments for command "+command.Name)
				return
			}
			if err := command.Function(args[1:]); err != nil {
				PRINT(env, "Error: "+err.Error())
				return
			}
			PRINT(env, "Command "+command.Name+" executed successfully")
			return
		}
	}
	PRINT(env, "Error: Unknown command "+args[0])
}
