package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func save() {
	for _, tokenMap := range tokenMaps {
		for _, token := range tokenMap {
			saveTokenToFile(token)
		}
	}

	for chain, tokenMap := range tokenMaps {
		tokens := make([]Token, 0, len(tokenMap))
		for _, token := range tokenMap {
			tokens = append(tokens, token)
		}
		saveTokensToFile(tokens, chain)
	}

	var tokens []Token
	for _, tokenMap := range tokenMaps {
		for _, token := range tokenMap {
			tokens = append(tokens, token)
		}
	}
	saveTokensToFile(tokens, 0)
}

func saveTokenToFile(token Token) {
	filePath := fmt.Sprintf("token/%d/%s.json", token.Chain, token.Address)
	tokensJSON, err := json.MarshalIndent(token, "", "	")
	// tokensJSON, err := json.Marshal(token)
	if err != nil {
		return
	}
	os.WriteFile(filePath, tokensJSON, 0644)
}

func saveTokensToFile(tokens []Token, chain int) {
	filePath := fmt.Sprintf("tokens/%d.json", chain)
	tokensJSON, err := json.MarshalIndent(tokens, "", "	")
	// tokensJSON, err := json.Marshal(tokens)
	if err != nil {
		return
	}
	os.WriteFile(filePath, tokensJSON, 0644)
}
