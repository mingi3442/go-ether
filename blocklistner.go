package main

import (
  // "context"
  // "encoding/hex"
  // "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/ethclient"
)

func BlockListener() error {
  _, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/adf81b4ad1e04a768dc24131023d50de")
  // fmt.Println("aa")
  return err
}
