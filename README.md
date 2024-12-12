# md - metadata for crypto tokens
Easily fetchable via frontend API

### Single Token
`/token/<chain_id>/<token_address>.json`

Example: Fetch USDC on Base network
```bash
curl -X GET https://raw.githubusercontent.com/enclave-so/md/refs/heads/main/token/8453/0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913.json
```

Response format:
```json
{
    "address":"0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
    "symbol":"USDC",
    "name":"USD Coin",
    "decimals":6,
    "logo":"https://ethereum-optimism.github.io/data/USDC/logo.png",
    "coingecko":"usd-coin",
    "sources":["coingecko","uniswap"]
}
```

### All Tokens
`/tokens/<chain_id>.json`

Example: Fetch all tokens on Base network
```bash
curl -X GET https://raw.githubusercontent.com/enclave-so/md/refs/heads/main/tokens/8453.json
```

Response format:
```json
{
    "tokens": [
        {
            "chain":8453,
            "address":"0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913",
            "symbol":"USDC",
            "name":"USD Coin",
            "decimals":6,
            "logo":"https://ethereum-optimism.github.io/data/USDC/logo.png",
            "coingecko":"usd-coin",
            "sources":["coingecko","uniswap"]
        }
    ]
}
```


## Tips

### Native Token Addresses
Native tokens (like ETH) use the zero address: `0x0000000000000000000000000000000000000000` (also known as `address(0)` in Solidity)

### All Tokens
Under chain `0`, you can find all tokens from all chains. `tokens/0.json` is a merged list of all tokens from all chains.

### Token Type
```ts
type Token = {
    chain: number
    address: string
    symbol: string
    name: string
    decimals?: number
    logo?: string
    coingecko?: string
    sources?: string[]
}
```

### EIP-55 Checksum Address Encoding
Convert addresses to their [EIP-55 checksum encoded](https://eips.ethereum.org/EIPS/eip-55) format using these libraries:

#### Using Viem (TypeScript)
```ts
import { getAddress } from 'viem'

getAddress('0x833589fcd6edb6e08f4c7c32d4f71b54bda02913') 
// '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913'
```

#### Using Ox (TypeScript)
```ts
import { Address } from 'ox'
 
const address = Address.from('0x833589fcd6edb6e08f4c7c32d4f71b54bda02913')
// '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913'
```

#### Using go-ethereum (Go)
```go
import (
    "github.com/ethereum/go-ethereum/common"
)

address := common.HexToAddress("0x833589fcd6edb6e08f4c7c32d4f71b54bda02913").Hex()
// '0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913'
```

### Price Data Integration
Fetch token prices from [CoinGecko](https://docs.coingecko.com/v3.0.1/reference/simple-price):

```ts
// Step 1: Fetch token metadata
const metadata = await fetch(
    'https://raw.githubusercontent.com/enclave_so/md/main/token/8453/0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913.json'
).then(res => res.json())

// Step 2: Fetch price using the CoinGecko ID from metadata
const price = await fetch(
    `https://api.coingecko.com/api/v3/simple/price?ids=${metadata.coingecko}&vs_currencies=usd`
).then(res => res.json())

console.log(`${metadata.symbol} price: $${price[metadata.coingecko].usd}`)
// Example output: "USDC price: $0.999328"
```

### Add your own tokenlist
Add your own tokenlist to `tokenlists.json` and create PR to merge it.
