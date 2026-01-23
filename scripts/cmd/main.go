package main

import (
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "sandbox",
		Short: "Inngest sandbox management tool",
	}

	root.AddCommand(updateInngestCmd())

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
