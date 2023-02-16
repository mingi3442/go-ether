package main

import (
  "context"
  // "encoding/hex"
  "fmt"
  "github.com/ethereum/go-ethereum/core/types"
  "github.com/ethereum/go-ethereum/ethclient"
  "log"
)

func BlockListener() error {
  // Infura에서 websocket URL을 통해 연결
  client, err := ethclient.Dial("wss://mainnet.infura.io/ws/v3/adf81b4ad1e04a768dc24131023d50de")
  if err != nil {
    log.Fatal(err)
  }
  //Block의 Header를 가져올 Channel 생성
  headers := make(chan *types.Header)
  //만든 Channel을 통해 연결
  sub, err := client.SubscribeNewHead(context.Background(), headers)
  if err != nil {
    log.Fatal(err)
  }
  //무한 Loop로 작동
  for {
    select {
    //Err를 받아올 경우
    case err := <-sub.Err():
      log.Fatal(err)
      //정상적으로 Header값을 가져오는 경우
    case header := <-headers:
      //header값의 Hash값을 통해 다시 연결
      block, err := client.BlockByHash(context.Background(), header.Hash())
      if err != nil {
        log.Fatal(err)
      }
      fmt.Println("----------------------------------------------------------------------")
      fmt.Println("Block Hash : " + block.Hash().Hex())
      fmt.Println("Block Number : ", block.Number().Uint64())
      fmt.Println("Block Timestamp : ", block.Time())
      fmt.Println("Block Nonce : ", block.Nonce())
      fmt.Println("Total Transactions : ", len(block.Transactions()))
      fmt.Println("----------------------------------------------------------------------")
    }
  }

}
