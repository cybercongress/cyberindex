package resources

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	resourcestypes "github.com/cybercongress/go-cyber/v2/x/resources/types"

	"github.com/forbole/juno/v3/types"

	"github.com/cybercongress/cyberindex/v2/database"
)

func HandleMsg(
	tx *types.Tx,
	msg sdk.Msg,
	db *database.CyberDb,
) error {
	if len(tx.Logs) == 0 {
		return nil
	}
	switch resourcesMsg := msg.(type) {
	case *resourcestypes.MsgInvestmint:
		return db.SaveInvestmints(
			resourcesMsg.Neuron,
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
