package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"slices"

	"github.com/ethereum/go-ethereum/common"
)

// func parseJSON[T any](body []byte) (T, error) {
// 	var v T
// 	err := json.Unmarshal(body, &v)
// 	return v, err
// }

func mergeToken(newToken Token) {
	if newToken.Chain == 0 || newToken.Address == "" {
		return
	}

	newToken.Address = common.HexToAddress(newToken.Address).Hex()

	tokenMap, ok := tokenMaps[newToken.Chain]
	if !ok {
		tokenMap = TokenMap{}
		tokenMaps[newToken.Chain] = tokenMap
	}

	token := tokenMap[newToken.Address]

	if token.Address == "" {
		token.Address = newToken.Address
	}

	if token.Chain == 0 {
		token.Chain = newToken.Chain
	}

	if token.Name == "" {
		token.Name = newToken.Name
	}

	if token.Symbol == "" {
		token.Symbol = newToken.Symbol
	}

	if token.Decimals == 0 {
		token.Decimals = newToken.Decimals
	}

	if token.Logo == "" {
		token.Logo = newToken.Logo
	}

	if token.Coingecko == "" {
		token.Coingecko = newToken.Coingecko
	}

	token.Sources = append(token.Sources, newToken.Sources...)
	slices.Sort(token.Sources)
	token.Sources = slices.Compact(token.Sources)
	tokenMap[newToken.Address] = token
}

// Delete previous tokens from token/ directory
func deletePrevious() {
	os.RemoveAll("token")
	os.RemoveAll("tokens")
}

// Create directories for tokens
func createDirectories() {
	os.Mkdir("token", 0755)
	os.Mkdir("tokens", 0755)

	for chainId := range tokenMaps {
		os.Mkdir(fmt.Sprintf("token/%d", chainId), 0755)
	}
}

func retrieveURL[T any](url string) (T, error) {
	var v T

	resp, err := http.Get(url)
	if err != nil {
		return v, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(body, &v)
	return v, err
}

func retrieveFile[T any](path string) (T, error) {
	var v T

	body, err := os.ReadFile(path)
	if err != nil {
		return v, err
	}

	err = json.Unmarshal(body, &v)
	return v, err
}

func toToken(fulltoken FullToken, source string) Token {
	return Token{
		Address:  fulltoken.Address,
		Chain:    fulltoken.ChainId,
		Symbol:   fulltoken.Symbol,
		Name:     fulltoken.Name,
		Decimals: fulltoken.Decimals,
		Logo:     fulltoken.LogoURI,
		Sources:  []string{source},
	}
}
