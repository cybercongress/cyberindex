package x

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/desmos-labs/juno/client"
	"github.com/desmos-labs/juno/config"
	"github.com/desmos-labs/juno/db"
	"github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/x/auth"
	"github.com/forbole/bdjuno/x/bank"
	bmodules "github.com/forbole/bdjuno/x/modules"
	"github.com/forbole/bdjuno/x/utils"

	cbd "github.com/cybercongress/cyberindex/database"
	"github.com/cybercongress/cyberindex/x/graph"
)

type ModulesRegistrar struct {
	parser messages.MessageAddressesParser
}

func NewModulesRegistrar(parser messages.MessageAddressesParser) *ModulesRegistrar {
	return &ModulesRegistrar{
		parser: parser,
	}
}

func (r *ModulesRegistrar) BuildModules(
	cfg *config.Config, encodingConfig *params.EncodingConfig, _ *sdk.Config, db db.Database, cp *client.Proxy,
) modules.Modules {

	bigDipperBd := database.Cast(db)
	cyberDb := &cbd.CyberDb{bigDipperBd}

	return []modules.Module{
		messages.NewModule(r.parser, encodingConfig.Marshaler, db),
		auth.NewModule(encodingConfig, utils.MustCreateGrpcConnection(cfg), bigDipperBd),
		bank.NewModule(encodingConfig, utils.MustCreateGrpcConnection(cfg), bigDipperBd),
		bmodules.NewModule(cfg, bigDipperBd),
		graph.NewModule(encodingConfig, utils.MustCreateGrpcConnection(cfg), cyberDb),
	}
}
