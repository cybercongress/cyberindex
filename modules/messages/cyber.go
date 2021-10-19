package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	graphtypes "github.com/cybercongress/go-cyber/x/graph/types"
	gridtypes "github.com/cybercongress/go-cyber/x/grid/types"
	resourcestypes "github.com/cybercongress/go-cyber/x/resources/types"
	junomessages "github.com/forbole/juno/v2/modules/messages"
)

func CyberMessageAddressesParser(cdc codec.Codec, cyberMsg sdk.Msg) ([]string, error) {
	switch msg := cyberMsg.(type) {

	case *graphtypes.MsgCyberlink:
		return []string{msg.Neuron}, nil

	case *gridtypes.MsgCreateRoute:
		return []string{msg.Source}, nil

	case *gridtypes.MsgEditRoute:
		return []string{msg.Source}, nil

	case *gridtypes.MsgEditRouteAlias:
		return []string{msg.Source}, nil

	case *gridtypes.MsgDeleteRoute:
		return []string{msg.Source}, nil

	case *resourcestypes.MsgInvestmint:
		return []string{msg.Neuron}, nil
	}

	return nil, junomessages.MessageNotSupported(cyberMsg)
}
