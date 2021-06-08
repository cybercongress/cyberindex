package main

import (
	"github.com/cybercongress/cyberindex/modules"
	"github.com/desmos-labs/juno/cmd"
	"github.com/forbole/bdjuno/types/config"

	cyberapp "github.com/cybercongress/go-cyber/app"
	initcmd "github.com/desmos-labs/juno/cmd/init"
	parsecmd "github.com/desmos-labs/juno/cmd/parse"
	"github.com/forbole/bdjuno/database"
)

func main() {
	initCfg := initcmd.NewConfig().
		WithConfigFlagSetup(config.SetupConfigFlags).
		WithConfigCreator(config.CreateConfig)

	parseCfg := parsecmd.NewConfig().
		WithConfigParser(config.ParseConfig).
		WithRegistrar(modules.NewRegistrar()).
		WithDBBuilder(database.Builder).
		WithEncodingConfigBuilder(cyberapp.MakeTestEncodingConfig)

	cfg := cmd.NewConfig("cyberindex").
		WithInitConfig(initCfg).
		WithParseConfig(parseCfg)

	executor := cmd.BuildDefaultExecutor(cfg)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}
