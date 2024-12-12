package main

func filter() {
	// for _, tokenMap := range tokenMaps {
	// 	for address, token := range tokenMap {
	// 		if token.Decimals == 0 { //|| token.Logo == "" {
	// 			delete(tokenMap, address)
	// 		}
	// 	}
	// }

	for chainId, tokenMap := range tokenMaps {
		if len(tokenMap) == 0 {
			delete(tokenMaps, chainId)
		}
	}
}
