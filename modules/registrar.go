package modules

import (
	"github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cybercongress/cyberindex/modules/energy"
	"github.com/cybercongress/cyberindex/modules/resources"
	"github.com/cybercongress/cyberindex/modules/utils"
	"github.com/desmos-labs/juno/client"
	"github.com/desmos-labs/juno/db"
	jmodules "github.com/desmos-labs/juno/modules"
	"github.com/desmos-labs/juno/modules/messages"
	"github.com/desmos-labs/juno/modules/registrar"
	juno "github.com/desmos-labs/juno/types"
	"github.com/forbole/bdjuno/database"
	"github.com/forbole/bdjuno/modules/auth"
	"github.com/forbole/bdjuno/modules/bank"
	bdmodules "github.com/forbole/bdjuno/modules/modules"

	cbd "github.com/cybercongress/cyberindex/database"
	"github.com/cybercongress/cyberindex/modules/graph"
	bdutils "github.com/forbole/bdjuno/modules/utils"
)

var (
	_ registrar.Registrar = &Registrar{}
)

type Registrar struct {}

func NewRegistrar() *Registrar {
	return &Registrar{}
}

func (r *Registrar) BuildModules(
	cfg juno.Config, encodingConfig *params.EncodingConfig, _ *sdk.Config, db db.Database, cp *client.Proxy,
) jmodules.Modules {

	parser := utils.AddressesParser
	bigDipperBd := database.Cast(db)
	cyberDb := &cbd.CyberDb{bigDipperBd}

	return []jmodules.Module{
		messages.NewModule(parser, encodingConfig.Marshaler, db),
		auth.NewModule(parser, encodingConfig, bdutils.MustCreateGrpcConnection(cfg), bigDipperBd),
		bank.NewModule(parser, encodingConfig, bdutils.MustCreateGrpcConnection(cfg), bigDipperBd),
		bdmodules.NewModule(cfg, bigDipperBd),
		graph.NewModule(parser, encodingConfig, bdutils.MustCreateGrpcConnection(cfg), cyberDb),
		energy.NewModule(parser, encodingConfig, bdutils.MustCreateGrpcConnection(cfg), cyberDb),
		resources.NewModule(parser, encodingConfig, bdutils.MustCreateGrpcConnection(cfg), cyberDb),
	}
}
