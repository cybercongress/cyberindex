package messages

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	graphtypes "github.com/cybercongress/go-cyber/v4/x/graph/types"
	gridtypes "github.com/cybercongress/go-cyber/v4/x/grid/types"
	resourcestypes "github.com/cybercongress/go-cyber/v4/x/resources/types"
)

// Deprecated: use CyberMessageAddressesParser
func CyberMessageAddressesParser(cyberMsg sdk.Msg) ([]string, error) {
	switch msg := cyberMsg.(type) {

	case *graphtypes.MsgCyberlink:
		return []string{msg.Neuron}, nil

	case *gridtypes.MsgCreateRoute:
		return []string{msg.Source}, nil

	case *gridtypes.MsgEditRoute:
		return []string{msg.Source}, nil

	case *gridtypes.MsgEditRouteName:
		return []string{msg.Source}, nil

	case *gridtypes.MsgDeleteRoute:
		return []string{msg.Source}, nil

	case *resourcestypes.MsgInvestmint:
		return []string{msg.Neuron}, nil
	}

	//return nil, return fmt.Errorf("message type not supported: %s", proto.MessageName(msg))
	return nil, nil
}
