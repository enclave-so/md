package main

type (
	Token struct {
		Chain     int      `json:"chain"`
		Address   string   `json:"address"`
		Symbol    string   `json:"symbol"`
		Name      string   `json:"name"`
		Decimals  int      `json:"decimals,omitempty"`
		Logo      string   `json:"logo,omitempty"`
		Coingecko string   `json:"coingecko,omitempty"`
		Sources   []string `json:"sources,omitempty"`
	}

	Tokenlist struct {
		Tokens []FullToken `json:"tokens"`
	}

	FullToken struct {
		ChainId  int    `json:"chainId"`
		Address  string `json:"address"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		LogoURI  string `json:"logoURI"`
	}

	CoingeckoToken struct {
		Id        string            `json:"id"`
		Symbol    string            `json:"symbol"`
		Name      string            `json:"name"`
		Platforms map[string]string `json:"platforms"`
	}

	CoingeckoPlatform struct {
		Id              string `json:"id"`
		ChainIdentifier int    `json:"chain_identifier"`
	}

	Chainlist struct {
		Chain  int    `json:"chain"`
		CoinId string `json:"coinId"`
	}

	TokenlistFile struct {
		URL  string `json:"url"`
		Name string `json:"name"`
	}
)
