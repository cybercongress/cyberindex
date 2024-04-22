package messages

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	liquiditytypes "github.com/cybercongress/go-cyber/v4/x/liquidity/types"
	junomessages "github.com/forbole/juno/v5/modules/messages"
)

func LiquidityMessageAddressesParser(cdc codec.Codec, liquidityMsg sdk.Msg) ([]string, error) {
	switch msg := liquidityMsg.(type) {
	case *liquiditytypes.MsgCreatePool:
		return []string{msg.PoolCreatorAddress}, nil

	case *liquiditytypes.MsgDepositWithinBatch:
		return []string{msg.DepositorAddress}, nil

	case *liquiditytypes.MsgSwapWithinBatch:
		return []string{msg.SwapRequesterAddress}, nil

	case *liquiditytypes.MsgWithdrawWithinBatch:
		return []string{msg.WithdrawerAddress}, nil
	}

	return nil, junomessages.MessageNotSupported(liquidityMsg)
}
