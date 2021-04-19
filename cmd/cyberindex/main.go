package main

import (
	"github.com/desmos-labs/juno/cmd"
	juno "github.com/desmos-labs/juno/types"

	"github.com/forbole/bdjuno/database"

	"github.com/cybercongress/cyberindex/x"

	"github.com/cybercongress/go-cyber/app"
)

func main() {
	executor := cmd.BuildDefaultExecutor(
		"cyberindex",
		x.NewModulesRegistrar(
			x.CyberMessageAddressesParser,
		),
		juno.DefaultSetup,
		app.MakeTestEncodingConfig,
		database.Builder,
	)

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
