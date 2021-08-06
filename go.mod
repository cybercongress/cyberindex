module github.com/cybercongress/cyberindex

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.7
	github.com/cybercongress/go-cyber v0.2.0-beta5
	github.com/desmos-labs/juno v0.0.0-20210726090239-dc0f7b55ac70
	github.com/forbole/bdjuno v0.0.0-20210805064553-5ed89b378186
	github.com/lib/pq v1.10.2
	github.com/tendermint/tendermint v0.34.11
	google.golang.org/grpc v1.37.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace google.golang.org/grpc => google.golang.org/grpc v1.33.2
