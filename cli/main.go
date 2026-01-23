package main

import (
	"github.com/alecthomas/kong"
)

type CLI struct {
	UpdateInngest UpdateInngestCmd `cmd:"" name:"update-inngest" help:"Update Inngest dependencies across all example projects"`
}

func main() {
	var cli CLI
	ctx := kong.Parse(&cli,
		kong.Name("cli"),
		kong.Description("Inngest sandbox CLI"),
		kong.UsageOnError(),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
