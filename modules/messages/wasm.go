package messages

import (
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/forbole/juno/v3/modules/messages"
)

func WasmMessageAddressesParser(cdc codec.Codec, wasmMsg sdk.Msg) ([]string, error) {
	switch msg := wasmMsg.(type) {
	case *wasmtypes.MsgStoreCode:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgInstantiateContract:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgExecuteContract:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgClearAdmin:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgUpdateAdmin:
		return []string{msg.Sender}, nil

	case *wasmtypes.MsgMigrateContract:
		return []string{msg.Sender}, nil
	}

	return nil, junomessages.MessageNotSupported(wasmMsg)
}
