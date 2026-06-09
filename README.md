# svclilib


<b> svclilib </b> is a Go library for building small command-driven CLI and REPL applications. It provides command registration, argument validation, formatted output helpers, usage printing, and environment-driven console styling. Cross-platform support.

Requires <b> Go 1.26 </b> or newer.


```go
// Usage Example

package main

import (
	"fmt"

	. "github.com/derekhandy/svclilib"
)

func hello(args []string) (error, string) {
	return nil, "Hello " + args[0]
}

func main() {
	env := Environment{
		Header:  "\n  example\n",
		Prefix:  "  ",
		Footer:  "  -\n\n",
		Glyphs:  []string{"  >  ", "  |\n", "  |   "},
		Spacing: "\n",
		UsageFormat: "{prefix}|   NAME\t{name}\n" +
			"{prefix}|   DESC\t{desc}\n" +
			"{prefix}|   USAGE\t{usage}",
		Commands: []Command{
			{
				Name:        "hello",
				Desc:        "prints a greeting",
				Usage:       "hello <name>",
				ArgRequired: 1,
				Function:    hello,
			},
		},
	}

	Execute(env, []string{"hello", "World"})
	fmt.Println("done")
}
```

## Install

Download from repository or add the module to an existing Go project:

```bash
go get github.com/derekhandy/svclilib
go mod tidy
```

## Use

There are two supported ways to call command logic: full execution with headers and footers, or direct command dispatch.

```go
// Prints header, command output, and footer when env.REPL is false.
svclilib.Execute(env, []string{"hello", "World"})
```

```go
// Dispatches a command without the Execute wrapper.
svclilib.RunCommand(env, []string{"hello", "World"})
```

## API

```go
// Defines console formatting and available commands.
type Environment struct {
	Header      string
	Prefix      string
	Footer      string
	Glyphs      []string
	Spacing     string
	Commands    []Command
	UsageFormat string
	REPL        bool
}

// Defines one callable command.
type Command struct {
	Name        string
	Desc        string
	Usage       string
	ArgRequired int
	Function    func(args []string) (error, string)
}

// Runs command dispatch with optional header and footer output.
svclilib.Execute(env svclilib.Environment, args []string)

// Runs command dispatch directly.
svclilib.RunCommand(env svclilib.Environment, args []string)
```

## Commands

```go
// Set a command with one required argument.

Command{
	Name:        "build",
	Desc:        "builds the current project",
	Usage:       "build <output>",
	ArgRequired: 1,
	Function: func(args []string) (error, string) {
		return nil, "building " + args[0]
	},
}
```

`ArgRequired` is checked before the command function runs. If the caller does not provide enough arguments, `svclilib` prints an error and the command usage block.

## Output Helpers

```go
// Logs the command invocation.
svclilib.Logc(env, []string{"build", "app"})

// Logs a footer, header, message, result output, usage block, or line spacer.
svclilib.Logf(env)
svclilib.Logh(env)
svclilib.Logm(env, "message")
svclilib.Logr(env, "output")
svclilib.Logu(env, []svclilib.Command{command})
svclilib.Logv(env)
```

## Info

`UsageFormat` supports these replacement fields:

```text
{name}
{desc}
{usage}
{prefix}
{spacing}
```

When `Environment.REPL` is true, `Execute` skips the header, command echo, and footer so the caller can manage the REPL display directly.

Command functions should return an empty string when no result output should be printed:

```go
return nil, ""
```

## NOTICE

<b> svclilib only dispatches functions provided by the calling application. Validation, permissions, side effects, and command safety remain the responsibility of the application using the library.</b>
