package graph

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	graphtypes "github.com/cybercongress/go-cyber/x/graph/types"

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
	switch graphMsg := msg.(type) {
	case *graphtypes.MsgCyberlink:
		return db.SaveCyberlink(
			graphMsg.Links,
			graphMsg.Address,
			tx.Timestamp,
			tx.Height,
			tx.TxHash,
		)
	}

	return nil
}
