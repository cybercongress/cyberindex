package energy

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	energytypes "github.com/cybercongress/go-cyber/x/energy/types"

	"github.com/desmos-labs/juno/types"

	"github.com/cybercongress/cyberindex/database"
)

func HandleMsg(
	tx *types.Tx,
	msg sdk.Msg,
	cdc codec.Marshaler,
	db *database.CyberDb,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}
	switch energyMsg := msg.(type) {
	case *energytypes.MsgCreateRoute:
		return db.SaveRoute(
			energyMsg.Source,
			energyMsg.Destination,
			energyMsg.Alias,
			tx.Timestamp,
			tx.Height,
			tx.TxHash,
		)
	case *energytypes.MsgEditRoute:
		return db.UpdateRouteValue(
			energyMsg.Source,
			energyMsg.Destination,
			energyMsg.Value,
		)
	case *energytypes.MsgEditRouteAlias:
		return db.UpdateRouteAlias(
			energyMsg.Source,
			energyMsg.Destination,
			energyMsg.Alias,
		)
	case *energytypes.MsgDeleteRoute:
		return db.DeleteRoute(
			energyMsg.Source,
			energyMsg.Destination,
		)
	}

	return nil
}
