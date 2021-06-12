package cmd

import (
	"flag"
	"fmt"
	"gtc/structs"
	"os"
)

func Execute() {
	if len(os.Args) == 1 {
		// Root flags here
		fmt.Printf("Root: ")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "workspace":
		WorkspaceCMD()

	case "todo":
		Todo()

	default: // Root with flags
		// Add root flags here
		// testRoot := flag.String("test", "", "test flag")
		flag.PrintDefaults()
		// fmt.Printf("Root test: %s", *testRoot)
		os.Exit(1)
	}

	// Check witch subcommand was Parsed using the FlagSet.Parsed() function.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)
	// if workspaceCommand.Parsed() {
	// 	fmt.Printf("workspaceTest: %s", *workspaceTest)
	// }

	// if todoCommand.Parsed() {
	// 	fmt.Printf("workspaceTest: %s", *todoTest)
	// }
}

func WorkspaceCMD() {
	workspaceCommand := flag.NewFlagSet("workspace", flag.ExitOnError)

	switch os.Args[2] {
	case "create":
		// Workspace subcommand flags
		name := workspaceCommand.String("name", "", "Workspace name")
		desc := workspaceCommand.String("desc", "", "Workspace description")
		workspaceCommand.Parse(os.Args[3:])

		if *name == "" {
			fmt.Printf("Name is required")
			os.Exit(1)
		}

		ws := structs.CreateWorkspace(*name, *desc)
		fmt.Printf("%+v is created", *ws)

	case "list":
		wss := structs.GetWorkspaces()
		for _, ws := range wss {
			fmt.Printf("%+v\n", *ws)
		}

	case "use":
		workspaceCommand.Parse(os.Args[3:])

	case "delete":
		workspaceCommand.Parse(os.Args[3:])

	case "status":
		workspaceCommand.Parse(os.Args[3:])

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func Todo() {
	todoCommand := flag.NewFlagSet("todo", flag.ExitOnError)
	switch os.Args[2] {
	case "add":
		todoCommand.Parse(os.Args[3:])

	case "delete":
		todoCommand.Parse(os.Args[3:])

	case "status":
		todoCommand.Parse(os.Args[3:])

	case "complete":
		todoCommand.Parse(os.Args[3:])

	default:
		flag.PrintDefaults()
		os.Exit(1)
	}
}
