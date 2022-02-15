# Run Undeployed Bytecode -- Go Version
- The typescript (ethers) version can be found [here](https://github.com/a16z/virtual-eth-contract-ethers)
- If you'd like to run some on-chain code in an `eth_call` without paying deployment costs, this is likely the most performant way.
- `go run cmd/main.go`
- Example is for a BalanceGetter, but easily adjustable with an ABI and some bytecode.