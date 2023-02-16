package main

import (
  "context"
  "encoding/hex"
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

      if len(block.Uncles()) > 0 {
        for _, uncle := range block.Uncles() {
          uncleFee := float64((uncle.Number.Uint64()+8-block.Number().Uint64())*2) / 8.0 // Uncle Block Fee 계산
          fmt.Println("****************************************************************")
          fmt.Println("Uncle Block Length : ", len(block.Uncles()))
          fmt.Println("Uncle Miner Address : ", uncle.Coinbase.Hex())
          fmt.Println("Uncle Block Number : ", uncle.Number.Uint64())
          fmt.Println("Uncle Block Reward : ", uncleFee)
          fmt.Println("****************************************************************")
        }
      }
      for _, tx := range block.Transactions() {
        fmt.Println("#####################################################################")
        fmt.Println("Transaction Hash : ", tx.Hash().Hex())
        fmt.Println("Transfer Values : ", tx.Value().String(), "wei") // Transaction의 value값은 wei단위로 나옴
        if tx.To() != nil {
          //Transcation에 To가 있을 경우
          fmt.Println("To Address : ", tx.To())
        } else {
          //Transcation에 To가 없을 경우 CA
          fmt.Println("Create Contract ")
        }
        fmt.Println("Transaction Nonce : ", tx.Nonce())
        fmt.Println("Transaction Gas Limit : ", tx.Gas()) // 예상 Gas Limit
        fmt.Println("Transaction GasFeeCap : ", tx.GasFeeCap().Uint64())
        fmt.Println("Transaction GasTipCap : ", tx.GasTipCap().Uint64())
        fmt.Println("Transcation Input Data : ", hex.EncodeToString(tx.Data())) // Input Data는 인코딩이 필요
        fmt.Println("#####################################################################")
      }
    }
  }

}
