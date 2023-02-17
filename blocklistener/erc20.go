package blocklistener

import (
  "log"
  "math/big"
  "strings"

  "github.com/ethereum/go-ethereum/accounts/abi/bind"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/liyue201/erc20-go/erc20"
)

//ERC20 Transaction 확인
func ERC20Transaction(data string) (string, string) {
  if len(data) != 136 {
    return "", "0"
  }
  methodId := data[:8]
  //To Address
  to := data[32:72]
  // Value
  value := data[37:136]
  //Erc20은 Method Id가 "a9059cbb"
  if methodId != "a9059cbb" {
    return "", "0"
  }
  i := new(big.Int)
  valueStr := strings.TrimLeft(value, "0")
  i.SetString(valueStr, 16)
  return to, i.String()
}

func GetContractInfo(client *ethclient.Client, toAddress *common.Address) (string, string, uint8) {
  // erc20.NewGGToken에는 이미 ERC20 abi파일이 있음
  instance, err := erc20.NewGGToken(*toAddress, client)
  if err != nil {
    log.Fatal(err)
  }
  name, err := instance.Name(&bind.CallOpts{})
  if err != nil {
    log.Fatal(err)
  }
  symbol, err := instance.Symbol(&bind.CallOpts{})
  if err != nil {
    log.Fatal(err)
  }
  decimals, err := instance.Decimals(&bind.CallOpts{})
  if err != nil {
    log.Fatal(err)
  }
  return symbol, name, decimals
}
