# cyberindex

Index for Cyber's networks

Supported:
1. blocks and precommits
2. transactions and messages
3. validators and their uptime
4. accounts and their balances
5. cyberlinks and particles
6. resources and supply
7. grid routes
8. wasm contracts
9. advanced views for analytics

cyberindex run selected indexers from BDJuno, see schemas and registrar.go

Run:
```
cyberindex parse
```

Example configuration in sample_config.yaml
```
cp sample_config.yaml ~/.cyberindex/config.yaml
``` 

To run in Docker fill `.env`, then:

```
make docker
```

Thanks for supporting Cosmos's ecosystem development:
- Fission Labs [Juno](https://github.com/fissionlabsio/juno),
- Desmos Labs [Juno](https://github.com/desmos-labs/juno) 
- Forbole [BDJuno](https://github.com/forbole/bdjuno) 
- Juno Community [Wasmx](https://github.com/disperze/wasmx)