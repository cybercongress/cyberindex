package utils

import (
	junomessages "github.com/desmos-labs/juno/modules/messages"
)

var AddressesParser = junomessages.JoinMessageParsers(
	junomessages.CosmosMessageAddressesParser,
	cyberMessageAddressesParser,
)