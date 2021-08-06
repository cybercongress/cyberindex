package modules

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cybercongress/cyberindex/modules/energy"
	"github.com/cybercongress/cyberindex/modules/resources"
	"github.com/desmos-labs/juno/client"
	"github.com/desmos-labs/juno/db"
	jmodules "github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	"github.com/desmos-labs/juno/modules/registrar"
	juno "github.com/desmos-labs/juno/types"
	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/modules"
	"github.com/forbole/bdjuno/modules/auth"
	"github.com/forbole/bdjuno/modules/bank"
	bdmodules "github.com/forbole/bdjuno/modules/modules"

	authttypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	cbd "github.com/cybercongress/cyberindex/database"
	"github.com/cybercongress/cyberindex/modules/graph"
	energytypes "github.com/cybercongress/go-cyber/x/energy/types"
	graphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	resourcestypes "github.com/cybercongress/go-cyber/x/resources/types"
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

func (r *Registrar) BuildModules(
	cfg juno.Config, encodingConfig *params.EncodingConfig, _ *sdk.Config, db db.Database, cp *client.Proxy,
) jmodules.Modules {
	bigDipperBd := database.Cast(db)
	cyberDb := &cbd.CyberDb{bigDipperBd}
	grpcConnection := client.MustCreateGrpcConnection(cfg)

	authClient := authttypes.NewQueryClient(grpcConnection)
	bankClient := banktypes.NewQueryClient(grpcConnection)
	graphClient := graphtypes.NewQueryClient(grpcConnection)
	energyClient := energytypes.NewQueryClient(grpcConnection)
	resourcesClient := resourcestypes.NewQueryClient(grpcConnection)

	return []jmodules.Module{
		messages.NewModule(r.parser, encodingConfig.Marshaler, db),
		auth.NewModule(r.parser, authClient, encodingConfig, bigDipperBd),
		bank.NewModule(r.parser, authClient, bankClient, encodingConfig, bigDipperBd),
		bdmodules.NewModule(cfg, bigDipperBd),
		graph.NewModule(r.parser, graphClient, encodingConfig, cyberDb),
		energy.NewModule(r.parser, energyClient, encodingConfig, cyberDb),
		resources.NewModule(r.parser, resourcesClient, encodingConfig, cyberDb),
	}
}
