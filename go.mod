module github.com/cybercongress/cyberindex

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.42.4
	github.com/cybercongress/go-cyber v0.2.0-alpha1.0.20210419093921-20902c023deb
	github.com/desmos-labs/juno v0.0.0-20210415095314-04a79a4e1908
	github.com/forbole/bdjuno v0.0.0-20210415101149-b0b01eee64e5
	github.com/go-co-op/gocron v0.3.3 // indirect
	github.com/rs/zerolog v1.20.0 // indirect
	github.com/tendermint/tendermint v0.34.9
	google.golang.org/grpc v1.36.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
