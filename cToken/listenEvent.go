package cToken

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	unitroller "liquidtor/Unitroller"
	"liquidtor/db"
	"log"
	"strings"
)

var index = 0

func GetEvent(host, comtroller string, conAddr []string, hosts []string) {
	exit := false
	defer func() {
		if exit {
			GetEvent(host, comtroller, conAddr, hosts)
		}
	}()
	host = hosts[index]
	index += 1
	if index == len(hosts) {
		index = 0
	}
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Println("ethclient.Dial", err)
		exit = true
		return
	}

	addresses := make([]common.Address, 0)
	for _, addr := range conAddr {
		contractAddress := common.HexToAddress(addr)
		addresses = append(addresses, contractAddress)
	}
	borrowEvent := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")
	query := ethereum.FilterQuery{
		Addresses: addresses,
		Topics:    [][]common.Hash{{borrowEvent}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Println("SubscribeFilterLogs", err)
		exit = true
		return
	}
	contractAbi, err := abi.JSON(strings.NewReader(CTokenABI))
	if err != nil {
		log.Println("abi.JSON", err)
		exit = true
		return
	}

	controller, err := unitroller.NewUnitroller(common.HexToAddress(comtroller), client)
	if err != nil {
		log.Println("unitroller.NewUnitroller", err)
		exit = true
		return
	}
	for {
		select {
		case err := <-sub.Err():
			log.Println("sub.Err()", err)
			exit = true
			return
		case vLog := <-logs:
			//fmt.Println(vLog) // pointer to event log
			//vLog.Topics

			for _, topic := range vLog.Topics {
				//topics[i] = vLog.Topics[i].Hex()
				if topic == borrowEvent {
					event := CTokenBorrow{}
					err := contractAbi.UnpackIntoInterface(&event, "Borrow", vLog.Data)
					if err != nil {
						log.Println("contractAbi.Unpack", err)
						continue
					}
					fmt.Println("listen event contract address:", vLog.Address.String())
					fmt.Println("borrower:", event.Borrower.String())
					fmt.Println("borrowAmount:", event.BorrowAmount.Uint64())

					_, liq, _, err := controller.GetAccountLiquidity(&bind.CallOpts{}, common.HexToAddress(event.Borrower.String()))

					db.AddData(event.Borrower.String(), vLog.Address.String(), vLog.BlockHash.String(), event.AccountBorrows.String(), event.TotalBorrows.String(), liq.String(), vLog.BlockNumber)

				}
			}

		}
	}
}
