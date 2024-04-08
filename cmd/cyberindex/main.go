package main

import (
	"github.com/cybercongress/cyberindex/v2/modules"
	"github.com/forbole/bdjuno/v3/types/config"
	"github.com/forbole/juno/v3/cmd"
	junomessages "github.com/forbole/juno/v3/modules/messages"

	cybermessages "github.com/cybercongress/cyberindex/v2/modules/messages"
	cyberapp "github.com/cybercongress/go-cyber/v2/app"
	"github.com/forbole/bdjuno/v3/database"
	initcmd "github.com/forbole/juno/v3/cmd/init"
	parsecmd "github.com/forbole/juno/v3/cmd/parse"
	parsetypes "github.com/forbole/juno/v3/cmd/parse/types"
	startcmd "github.com/forbole/juno/v3/cmd/start"
)

func main() {
	initCfg := initcmd.NewConfig().
		WithConfigCreator(config.Creator)

	parseCfg := parsetypes.NewConfig().
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(cyberapp.MakeTestEncodingConfig).
		WithRegistrar(modules.NewRegistrar(getAddressesParser()))

	cfg := cmd.NewConfig("cyberindex").
		WithInitConfig(initCfg).
		WithParseConfig(parseCfg)

	rootCmd := cmd.RootCmd(cfg.GetName())

	rootCmd.AddCommand(
		cmd.VersionCmd(),
		initcmd.NewInitCmd(cfg.GetInitConfig()),
		parsecmd.NewParseCmd(cfg.GetParseConfig()),
		startcmd.NewStartCmd(cfg.GetParseConfig()),
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
		cybermessages.WasmMessageAddressesParser,
		cybermessages.LiquidityMessageAddressesParser,
	)
}
