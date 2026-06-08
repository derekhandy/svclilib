package svclilib

//
//	Summary:
//
//		Header		Prints before a command is called. Acts as task start identifier.
//		Prefix		Prefix before any message is printed excluding header and footer.
//		Footer		Prints after a command is executed. Acts as task end identifier.
//		Glyphs		Input and output prefixes for command execution Logs.
//
//		Spacing		Spacing string used to separate:
// 						- Console output from header
// 						- Header from message
//						- Message from footer
//						- Footer from continuing console
//
//		Commands	List of functions that can be called by name.
//		Silent		Only prints error Logs.
//
type Environment struct {
	Header      string
	Prefix      string
	Footer      string
	Glyphs      []string
	Spacing     string
	Commands    []Command
	Silent      bool
	UsageFormat string
}

//
//	Summary:
//
//		Use "" as string for func(args []string) (error, string)
// 		return parameter to prevent Printr() from printing.
//
//		Example: return nil, ""
//
type Command struct {
	Name        string
	Desc        string
	ArgRequired int
	Function    func(args []string) error
}

func Execute(env Environment, args []string) {
	Printh(env)

	RunCommand(env, args)

	Printf(env)
}

func RunCommand(env Environment, args []string) {
	if len(args) == 0 {
		Printm(env, "Error: No command provided")
		return
	}

	Printc(env, args)

	for _, command := range env.Commands {
		if command.Name == args[0] {
			if len(args) < command.ArgRequired+1 {
				Printm(env, "Error: Not enough arguments for command "+command.Name)
			} else if err := command.Function(args[1:]); err != nil {
				Printm(env, "Error: "+err.Error())
			} /* else if !env.Silent {
				Printr(env)
			}*/
			return
		}
	}

	Printm(env, "Error: Unknown command "+args[0])
}
