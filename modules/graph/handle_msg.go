package graph

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	graphtypes "github.com/cybercongress/go-cyber/v2/x/graph/types"

	"github.com/forbole/juno/v3/types"

	"github.com/cybercongress/cyberindex/v1/database"
)

func HandleMsg(
	tx *types.Tx,
	msg sdk.Msg,
	db *database.CyberDb,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}
	switch graphMsg := msg.(type) {
	case *graphtypes.MsgCyberlink:
		return db.SaveCyberlinks(
			graphMsg.Links,
			graphMsg.Neuron,
			tx.Timestamp,
			tx.Height,
			tx.TxHash,
		)
	}

	return nil
}
