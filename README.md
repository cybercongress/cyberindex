# cyberindex

Bostrom edition, stargate support

Run:
```
cyberindex parse
```
Example configuration in sample_config, put to .cyberindex/config.toml in user's home dir 

To run in Docker fill `.env`, then:

```
make docker
```

cyberindex run selected indexers from BDJuno, see schemas and registrar.go

Thanks to Fission Labs [Juno](https://github.com/fissionlabsio/juno), Desmos Labs [Juno](https://github.com/desmos-labs/juno) and Forbole [BDJuno](https://github.com/forbole/bdjuno) for supporting Cosmos's  ecosystem development 