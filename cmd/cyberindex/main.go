package main

import (
	"github.com/cybercongress/cyberindex/modules"
	"github.com/forbole/juno/v2/cmd"
	junomessages "github.com/forbole/juno/v2/modules/messages"

	cybermessages "github.com/cybercongress/cyberindex/modules/messages"
	cyberapp "github.com/cybercongress/go-cyber/app"
	"github.com/forbole/bdjuno/v2/database"
	initcmd "github.com/forbole/juno/v2/cmd/init"
	parsecmd "github.com/forbole/juno/v2/cmd/parse"
)

func main() {
	parseCfg := parsecmd.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(cyberapp.MakeTestEncodingConfig).
		WithRegistrar(modules.NewRegistrar(getAddressesParser()))

	cfg := cmd.NewConfig("cyberindex").WithParseConfig(parseCfg)

	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.InitCmd(cfg.GetInitConfig()),
		parsecmd.ParseCmd(cfg.GetParseConfig()),
	)

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
