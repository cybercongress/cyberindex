package modules

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cyberdb "github.com/cybercongress/cyberindex/database"
	"github.com/cybercongress/cyberindex/modules/graph"
	"github.com/cybercongress/cyberindex/modules/grid"
	"github.com/cybercongress/cyberindex/modules/liquidity"
	"github.com/cybercongress/cyberindex/modules/resources"
	"github.com/cybercongress/cyberindex/modules/wasm"
	"github.com/cybercongress/go-cyber/app"
	"github.com/forbole/bdjuno/v2/database"
	"github.com/forbole/bdjuno/v2/modules"
	"github.com/forbole/bdjuno/v2/modules/auth"
	"github.com/forbole/bdjuno/v2/modules/bank"
	banksource "github.com/forbole/bdjuno/v2/modules/bank/source"
	localbanksource "github.com/forbole/bdjuno/v2/modules/bank/source/local"
	remotebanksource "github.com/forbole/bdjuno/v2/modules/bank/source/remote"
	bjmodules "github.com/forbole/bdjuno/v2/modules/modules"
	jmodules "github.com/forbole/juno/v2/modules"
	"github.com/forbole/juno/v2/modules/messages"
	"github.com/forbole/juno/v2/modules/registrar"
	nodeconfig "github.com/forbole/juno/v2/node/config"
	"github.com/forbole/juno/v2/node/local"
	"github.com/forbole/juno/v2/node/remote"
	//"github.com/tendermint/liquidity/x/liquidity"

	//"github.com/tendermint/liquidity/x/liquidity"
	"github.com/tendermint/tendermint/libs/log"
	"os"
)

var (
	_ registrar.Registrar = &Registrar{}
)

type Registrar struct {
	parser messages.MessageAddressesParser
}

func NewRegistrar(parser messages.MessageAddressesParser) *Registrar {
	return &Registrar{
		parser: modules.UniqueAddressesParser(parser),
	}
}

func (r *Registrar) BuildModules(ctx registrar.Context) jmodules.Modules {
	cdc := ctx.EncodingConfig.Marshaler
	bigDipperBd := database.Cast(ctx.Database)
	cyberDb := &cyberdb.CyberDb{Db: bigDipperBd}
	sources, err := BuildSources(ctx.JunoConfig.Node, ctx.EncodingConfig)
	if err != nil {
		panic(err)
	}
	authModule := auth.NewModule(r.parser, cdc, bigDipperBd)
	bankModule := bank.NewModule(r.parser, sources.BankSource, cdc, bigDipperBd)
	graphModule := graph.NewModule(r.parser, cdc, cyberDb)
	gridModule := grid.NewModule(r.parser, cdc, cyberDb)
	wasmModule := wasm.NewModule(r.parser, cdc, cyberDb)
	resourceModule := resources.NewModule(r.parser, cdc, cyberDb)
	liquidityModule := liquidity.NewModule(r.parser, cdc, bankModule, authModule, cyberDb)

	return []jmodules.Module{
		messages.NewModule(r.parser, cdc, ctx.Database),
		authModule,
		bankModule,
		bjmodules.NewModule(ctx.JunoConfig.Chain, bigDipperBd),
		graphModule,
		gridModule,
		resourceModule,
		wasmModule,
		liquidityModule,
	}
}

type Sources struct {
	BankSource      banksource.Source
}

func BuildSources(nodeCfg nodeconfig.Config, encodingConfig *params.EncodingConfig) (*Sources, error) {
	switch cfg := nodeCfg.Details.(type) {
	case *remote.Details:
		return buildRemoteSources(cfg)
	case *local.Details:
		return buildLocalSources(cfg, encodingConfig)

	default:
		return nil, fmt.Errorf("invalid configuration type: %T", cfg)
	}
}

func buildLocalSources(cfg *local.Details, encodingConfig *params.EncodingConfig) (*Sources, error) {
	source, err := local.NewSource(cfg.Home, encodingConfig)
	if err != nil {
		return nil, err
	}

	app := app.NewApp(
		log.NewTMLogger(log.NewSyncWriter(os.Stdout)), source.StoreDB, nil, true, map[int64]bool{},
		cfg.Home, 0, app.MakeEncodingConfig(), nil, nil, nil,
	)

	return &Sources{
		BankSource:     localbanksource.NewSource(source, banktypes.QueryServer(app.BankKeeper)),
	}, nil
}

func buildRemoteSources(cfg *remote.Details) (*Sources, error) {
	source, err := remote.NewSource(cfg.GRPC)
	if err != nil {
		return nil, fmt.Errorf("error while creating remote source: %s", err)
	}

	return &Sources{
		BankSource:     remotebanksource.NewSource(source, banktypes.NewQueryClient(source.GrpcConn)),
	}, nil
}

