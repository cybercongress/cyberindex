module github.com/cybercongress/cyberindex

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.5
	github.com/cybercongress/go-cyber v0.2.0-alpha1.0.20210607083035-3971386c95b9
	github.com/desmos-labs/juno v0.0.0-20210513082948-fad7f160e2cd
	github.com/forbole/bdjuno v0.0.0-20210603080009-7adc4516ebe7
	github.com/lib/pq v1.10.2
	github.com/tendermint/tendermint v0.34.10
	google.golang.org/grpc v1.37.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2

replace github.com/forbole/bdjuno => github.com/litvintech/bdjuno v0.0.0-20210607150853-22ce4779b220
