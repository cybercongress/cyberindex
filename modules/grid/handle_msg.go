package grid

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	gridtypes "github.com/cybercongress/go-cyber/x/grid/types"

	"github.com/forbole/juno/v2/types"

	"github.com/cybercongress/cyberindex/database"
)

func HandleMsg(
	tx *types.Tx,
	msg sdk.Msg,
	db *database.CyberDb,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}
	switch energyMsg := msg.(type) {
	case *gridtypes.MsgCreateRoute:
		return db.SaveRoute(
			energyMsg.Source,
			energyMsg.Destination,
			energyMsg.Alias,
			tx.Timestamp,
			tx.Height,
			tx.TxHash,
		)
	case *gridtypes.MsgEditRoute:
		return db.UpdateRouteValue(
			energyMsg.Source,
			energyMsg.Destination,
			energyMsg.Value,
		)
	case *gridtypes.MsgEditRouteAlias:
		return db.UpdateRouteAlias(
			energyMsg.Source,
			energyMsg.Destination,
			energyMsg.Alias,
		)
	case *gridtypes.MsgDeleteRoute:
		return db.DeleteRoute(
			energyMsg.Source,
			energyMsg.Destination,
		)
	}

	return nil
}
