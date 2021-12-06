package cToken

import (
	"context"
	"fmt"
	unitroller "liquidtor/Unitroller"
	"liquidtor/db"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func FilOne(host, comtroller string, conAddr []string,i int64) {
//9601359
	//12508458
	client, err := ethclient.Dial(host)
	if err != nil {
		log.Println(err)
		return
	}
	for ; i >= 9601359; {
		FilLog(host, comtroller, conAddr, i,client)
		i -= 2
	}
}

func FilLog(host, comtroller string, conAddr []string, height int64,client *ethclient.Client) {
	ex:=false
    defer func() {
    	if ex{
			FilOne(host,comtroller,conAddr,height)
		}

	}()


	addresses := make([]common.Address, 0)
	for _, addr := range conAddr {
		contractAddress := common.HexToAddress(addr)
		addresses = append(addresses, contractAddress)
	}
	//contractAddress := common.HexToAddress(conAddress)
	borrowEvent := common.HexToHash("0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(height),
		ToBlock:   big.NewInt(height + 1),
		Addresses: addresses,
		Topics:    [][]common.Hash{{borrowEvent}},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
		ex = true
		return
	}

	contractAbi, err := abi.JSON(strings.NewReader(CTokenABI))
	if err != nil {
		log.Println(err)
		ex = true
		return
	}
	controller, _ := unitroller.NewUnitroller(common.HexToAddress(comtroller), client)

	for _, vLog := range logs {
		fmt.Println("height:", vLog.BlockNumber)
		for _, topic := range vLog.Topics {
			//topics[i] = vLog.Topics[i].Hex()
			if topic == borrowEvent {
				event := CTokenBorrow{}
				err := contractAbi.UnpackIntoInterface(&event, "Borrow", vLog.Data)
				if err != nil {
					log.Println(err)
					ex = true
					return
				}
				fmt.Println("borrower:", event.Borrower.String())
				fmt.Println("borrowAmount:", event.BorrowAmount.Uint64())
				_, liq, _, err := controller.GetAccountLiquidity(&bind.CallOpts{}, common.HexToAddress(event.Borrower.String()))
				db.AddData(event.Borrower.String(), vLog.Address.String(),vLog.BlockHash.String(), event.AccountBorrows.String(), event.TotalBorrows.String(), liq.String(),vLog.BlockNumber)

			}
		}

		//fmt.Println(topics[0]) //
	}

	//	fmt.Println(hash.Hex()) // 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80

}
