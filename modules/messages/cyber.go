package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	energytypes "github.com/cybercongress/go-cyber/x/energy/types"
	graphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	resourcestypes "github.com/cybercongress/go-cyber/x/resources/types"
	junomessages "github.com/desmos-labs/juno/modules/messages"
)

func CyberMessageAddressesParser(cdc codec.Marshaler, cyberMsg sdk.Msg) ([]string, error) {
	switch msg := cyberMsg.(type) {

	case *graphtypes.MsgCyberlink:
		return []string{msg.Address}, nil

	case *energytypes.MsgCreateRoute:
		return []string{msg.Source}, nil

	case *energytypes.MsgEditRoute:
		return []string{msg.Source}, nil

	case *energytypes.MsgEditRouteAlias:
		return []string{msg.Source}, nil

	case *energytypes.MsgDeleteRoute:
		return []string{msg.Source}, nil

	case *resourcestypes.MsgInvestmint:
		return []string{msg.Agent}, nil
	}

	return nil, junomessages.MessageNotSupported(cyberMsg)
}
