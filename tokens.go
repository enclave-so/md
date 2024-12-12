package main

type TokenMap map[string]Token

var (
	// map[chainId]TokenMap
	tokenMaps = map[int]TokenMap{}
)
