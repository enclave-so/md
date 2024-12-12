package main

import (
	"fmt"
)

func collect() {
	collectFile("default.json")
	collectTokenlists()
	collectSushiswap()
	collectCoingecko()
}

func collectFile(path string) {
	tokens, err := retrieveFile[[]Token](path)
	if err != nil {
		fmt.Printf("Failed to parse tokenlist: %v", err)
		return
	}

	for _, token := range tokens {
		mergeToken(token)
	}
}

func collectTokenlists() {
	tokenlists, _ := retrieveFile[[]TokenlistFile]("tokenlists.json")
	for _, tokenlist := range tokenlists {
		collectTokenlist(tokenlist)
	}
}

func collectTokenlist(tokenlistFile TokenlistFile) {
	tokenlist, err := retrieveURL[Tokenlist](tokenlistFile.URL)
	if err != nil {
		fmt.Printf("Failed to parse tokenlist: %v", err)
		return
	}

	for _, tokenFromList := range tokenlist.Tokens {
		mergeToken(toToken(tokenFromList, tokenlistFile.Name))
	}
}

func collectCoingecko() {
	coingeckoTokens, err := retrieveURL[[]CoingeckoToken]("https://api.coingecko.com/api/v3/coins/list?include_platform=true")
	if err != nil {
		fmt.Printf("Failed to fetch coingecko: %v", err)
		return
	}

	coingeckoPlatforms := coingeckoPlatforms()

	for _, coingeckoToken := range coingeckoTokens {
		for platform, platformAddress := range coingeckoToken.Platforms {
			if chainId, ok := coingeckoPlatforms[platform]; ok {
				mergeToken(Token{
					Address:   platformAddress,
					Chain:     chainId,
					Symbol:    coingeckoToken.Name,
					Name:      coingeckoToken.Name,
					Coingecko: coingeckoToken.Id,
					Sources:   []string{"coingecko"},
				})
			}
		}
	}
}

func coingeckoPlatforms() map[string]int {
	platforms := map[string]int{}
	coingeckoPlatforms, err := retrieveURL[[]CoingeckoPlatform]("https://api.coingecko.com/api/v3/asset_platforms")
	if err != nil {
		fmt.Printf("Failed to fetch coingecko: %v", err)
		return platforms
	}

	for _, platform := range coingeckoPlatforms {
		if platform.ChainIdentifier != 0 {
			platforms[platform.Id] = platform.ChainIdentifier
		}
	}

	return platforms
}

// func collectTrustWalletChains() {
// 	body, err := retrieveURL("https://raw.githubusercontent.com/trustwallet/assets/refs/heads/master/blockchains/ethereum/chainlist.json")
// 	if err != nil {
// 		fmt.Printf("Failed to fetch trustwallet: %v", err)
// 		return
// 	}

// 	chainlist, err := parseJSON[[]Chainlist](body)
// 	if err != nil {
// 		fmt.Printf("Failed to parse trustwallet: %v", err)
// 		return
// 	}

// 	for _, chain := range chainlist {

// 	}
// }

// func collectTokenlists(url string, source string) {
// 	body, _ := retrieveURL(url)

// 	tokenlist, err := parseJSON[Tokenlist](body)
// 	if err != nil {
// 		fmt.Printf("Failed to parse tokenlist: %v", err)
// 		return
// 	}
// }

func collectSushiswap() {
	chains := []string{
		"ethereum",
		"ropsten",
		"rinkeby",
		"goerli",
		"kovan",
		"fantom",
		"fantom-testnet",
		"polygon",
		"polygon-testnet",
		"xdai",
		"bsc",
		"bsc-testnet",
		"moonbase",
		"avalanche",
		"fuji",
		"heco",
		"heco-testnet",
		"harmony",
		"harmony-testnet",
		"okex",
		"okex-testnet",
		"arbitrum",
		"celo",
		"palm",
		"moonriver",
		"fuse",
		"telos",
		"moonbeam",
		"optimism",
		"kava",
		"metis",
		"arbitrum-nova",
		"boba-avax",
		"boba",
		"bttc",
		"boba-bnb",
		"thundercore",
		"polygon-zkevm",
		"core",
		"haqq",
		"zksync-era",
		"linea",
		"scroll",
		"filecoin",
		"cronos",
		"zetachain",
		"blast",
		"skale-europa",
		"rootstock",
	}

	for _, chain := range chains {
		tokens, err := retrieveURL[[]FullToken](fmt.Sprintf("https://raw.githubusercontent.com/sushiswap/list/refs/heads/master/lists/token-lists/default-token-list/tokens/%s.json", chain))
		if err != nil {
			fmt.Printf("Failed to parse tokenlist: %v", err)
			return
		}

		for _, token := range tokens {
			mergeToken(toToken(token, "sushiswap"))
		}
	}
}
