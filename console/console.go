package console

import "fmt"

func PrintAllCommandsDescription() {
	fmt.Println("\033[36m┌──────────────────────────────────────────────────────────┐")
	fmt.Println("│                   🚀  AVAILABLE COMMANDS                  │")
	fmt.Println("└──────────────────────────────────────────────────────────┘\033[0m")

	//fmt.Println("\033[34m  GENERAL\033[0m")
	fmt.Println("    \033[33m-h\033[0m")
	fmt.Println("        List all available commands\n")

	//fmt.Println("\033[32m  GENERATORS\033[0m")
	fmt.Println("    \033[33m-gen\033[0m")
	fmt.Println("        Generate React boilerplate using Feature-Sliced Design (FSD)\n")

	//fmt.Println("\033[33m  CREATE COMMANDS\033[0m")
	fmt.Println("    \033[33m-c page <page-name>\033[0m")
	fmt.Println("        Create a new page inside the \"pages\" folder\n")

	fmt.Println("    \033[33m-c feature <feature-name>\033[0m")
	fmt.Println("        Create a new feature inside the \"features\" folder\n")

	fmt.Println("    \033[33m-c entity <entity-name>\033[0m")
	fmt.Println("        Create a new entity inside the \"entities\" folder")
}

func PrintCreateCommandError() {
	fmt.Println("\033[31m❌ Error:\033[0m Missing required arguments.")

	fmt.Println("\033[36mUsage:\033[0m")
	fmt.Println("  c page <name>")
	fmt.Println("  c feature <name>")
	fmt.Println("  c entity <name>")

	fmt.Println("\033[36mExamples:\033[0m")
	fmt.Println("  c page dashboard")
	fmt.Println("  c feature auth")
	fmt.Println("  c entity user\n")

	fmt.Println("\033[33mTip:\033[0m Run -h to see all available commands.")
}
