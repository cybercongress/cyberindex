package main

import (
	"github.com/cybercongress/cyberindex/modules"
	"github.com/desmos-labs/juno/cmd"
	junomessages "github.com/desmos-labs/juno/modules/messages"

	cybermessages "github.com/cybercongress/cyberindex/modules/messages"
	cyberapp "github.com/cybercongress/go-cyber/app"
	parsecmd "github.com/desmos-labs/juno/cmd/parse"
	"github.com/forbole/bdjuno/database"
)

func main() {
	parseCfg := parsecmd.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(cyberapp.MakeTestEncodingConfig).
		WithRegistrar(modules.NewRegistrar(getAddressesParser()))

	cfg := cmd.NewConfig("cyberindex").WithParseConfig(parseCfg)

	executor := cmd.BuildDefaultExecutor(cfg)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func getAddressesParser() junomessages.MessageAddressesParser {
	return junomessages.JoinMessageParsers(
		junomessages.CosmosMessageAddressesParser,
		cybermessages.CyberMessageAddressesParser,
	)
}
