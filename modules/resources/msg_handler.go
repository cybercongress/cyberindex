package resources

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	resourcestypes "github.com/cybercongress/go-cyber/x/resources/types"

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
	switch resourcesMsg := msg.(type) {
	case *resourcestypes.MsgInvestmint:
		return db.SaveInvestmints(
			resourcesMsg.Agent,
			resourcesMsg.Amount,
			resourcesMsg.Resource,
			resourcesMsg.Length,
			tx.Timestamp,
			tx.Height,
			tx.TxHash,
		)
	}

	return nil
}
