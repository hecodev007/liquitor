package main

import (
	"context"
	"fmt"
	"liquidtor/liquitor"

	"github.com/dabankio/wallet-core/qa/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"io/ioutil"
	unitroller "liquidtor/Unitroller"
	"math/big"
	"net/http"
	"strings"
	"testing"
)

func TestLi(t *testing.T) {
	mainet := "wss://mainnet.infura.io/ws/v3/fcfaf99bc9f94b148a65108207306f9e"
	client, err := ethclient.Dial(mainet)
	if err != nil {
		fmt.Println(err)
		return
	}
	controller, _ := unitroller.NewUnitroller(common.HexToAddress("0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b"), client)
	_, liq, shortAll, err := controller.GetAccountLiquidity(&bind.CallOpts{}, common.HexToAddress("0x59fCc41D9b705bAA430a10570f2f02B124dee27d"))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(liq)
	fmt.Println(shortAll.String())
}

func TestLog(t *testing.T) {
	cEth := "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"
	cDai := "0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"
	cWbtc := "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"
	cUsdc := "0x39aa39c021dfbae8fac545936693ac917d5e7563"
	cUsdt := "0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9"
	cBat := "0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"
	cComp := "0x70e36f6bf80a52b3b46b3af8e106cc0ed743e8e4"
	cLink := "0xface851a4921ce59e912d19329929ce6da6eb0c7"
	cRep := "0x158079ee67fce2f58472a96584a73c7ab9ac95c1"
	cSai := "0xf5dce57282a584d2746faf1593d3121fcac444dc"
	cTusd := "0x12392f67bdf24fae0af363c24ac620a2f67dad86"
	cUni := "0x35a18000230da775cac24873d00ff85bccded550"
	cZrx := "0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"
	cWbtc2 := "0xccf4429db6322d5c611ee964527d42e5d685dd6a"
	contractArray := []string{cEth, cDai, cWbtc, cUsdc, cUsdt, cBat, cComp, cLink, cRep, cSai, cTusd, cUni, cZrx, cWbtc2}
	fmt.Println(contractArray)
	//mainet := "wss://mainnet.infura.io/ws/v3/fcfaf99bc9f94b148a65108207306f9e"
	client, err := ethclient.Dial(host)
	if err != nil {
		fmt.Println(err)
		return
	}

	liq, err := liquitor.NewLiquidtor(common.HexToAddress(liqContract), client)
	LiquitorInit(liq)

	//	cToken.FilOne(mainet, "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b", contractArray, 9)
	//	cToken.GetEvent(mainet, "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b", contractArray)
	fmt.Println(1)
	select {}

	//	listen event contract address: 0x39AA39c021dfbaE8faC545936693aC917d5E7563
	//borrower: 0x337699a15e9CB2236045D9D29427f17eD1C4773f
	//borrowAmount: 500000000
}

func TestToEther(t *testing.T) {
	InitConfigFromJson()

	//r := ToEther(new(big.Int).Setint64(1200000000), 9)
	//fmt.Println(r)
}

func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return "", err
	}
	return string(robots), nil
}

//0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8
func TestGetLiqBorrow(t *testing.T) {
	//https://ropsten.etherscan.io/txs?a=0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8&p=29
	for i := 0; i < 59; i++ {
		url := fmt.Sprintf("https://ropsten.etherscan.io/txs?a=%s&p=%d", "0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8", i)
		ret, err := Get(url)
		if err != nil {
			fmt.Println(err)
			if strings.Contains(err.Error(), "connected host") {
				continue
			}
			break
		}
		fmt.Println(1)
		if strings.Contains(ret, "liquidateBorrow") {
			fmt.Println(ret)
		}
	}
}
func toWei(amount decimal.Decimal, dec int8) decimal.Decimal {
	return amount.Mul(decimal.New(1, int32(dec)))
}
func mul(amount decimal.Decimal, a *big.Int) *big.Int {
	return amount.Mul(decimal.NewFromBigInt(a, 0)).BigInt()

}

func GetAmountOut(amountIn int) int {
	reservA := 10000
	reservB := 10000
	//amountIn:=500
	inWithFee := amountIn * 997
	nuerator := inWithFee * reservB
	//	demon:=reservA*1000
	demon := reservA*1000 + inWithFee
	amountOut := nuerator / demon
	return amountOut
}

func TestUniCal(t *testing.T) {
	reservA := 10000
	reservB := 10000

	amountIn := 500
	inWithFee := amountIn * 997
	nuerator := inWithFee * reservB

	demon := reservA*1000 + inWithFee
	amountOut := nuerator / demon
	p1 := decimal.NewFromInt(int64(amountOut)).Div(decimal.NewFromInt(int64(amountIn)))
	fmt.Println(p1)

	reservA += amountIn
	reservB -= amountOut
	p2 := decimal.NewFromInt(int64(reservB)).Div(decimal.NewFromInt(int64(reservA)))
	fmt.Println(p2)

}

func TestToWei(t *testing.T) {
	//fmt.Println(mul(decimal.NewFromFloat(0.1),big.NewInt(10000233)))
	//0x52bbe01d3ef0af3adb06175020ee0b2f4c99864a900f59058d27bcfe265443fb
	client, err := ethclient.Dial(host)
	if err != nil {
		fmt.Println(err)
		return
	}

	rec, err := client.TransactionReceipt(context.Background(), common.HexToHash("0x52bbe01d3ef0af3adb06175020ee0b2f4c99864a900f59058d27bcfe265443fb"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rec.Status)
}

func TestLiqBorrow(t *testing.T) {
	a0 := decimal.NewFromFloat(1.0)
	a1 := decimal.NewFromFloat(0)
	for i := 1; i <= 10; i++ {
		r := decimal.NewFromFloat(1).Add(decimal.NewFromFloat(110).Sub(decimal.New(int64(i), 0).Mul(decimal.NewFromFloat(10))).Div(decimal.NewFromFloat(100)))

		//fmt.Println(r)
		a1 = a0.Mul(r)
		fmt.Println(a1)
		a0 = a1
	}

	sum := 0
	for i := 1; i <= 10; i++ {
		an := 300 + (i-1)*155
		sum += an
		fmt.Println(an)
	}
	fmt.Println(sum)

	//ctoken(cEth) mint
	//comptroller enter markets (ctoken = cEth)
	//ctoken(usdt) borrow

	//const host = "https://ropsten.infura.io/v3/fcfaf99bc9f94b148a65108207306f9e"
	/*	const host = "wss://ropsten.infura.io/ws/v3/bcb800e6ebc0432ea9aac3c68b15577f"
		//const liqContract = "0x923eca9a79bf9e8a06dc5d160af3a007068f2a7d"
		const comContract = "0xcfa7b0e37f5AC60f3ae25226F5e39ec59AD26152"
		const cUsdt = "0xF6958Cf3127e62d3EB26c79F4f45d3F3b2CcdeD4"
		const cEth = "0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8"
		const OraclePrice = "0xb90c96607b45f9bb7509861a1ce77cb8a72edfb2"
		client, err := ethclient.Dial(host)
		if err != nil {
			fmt.Println(err)
			return
		}

		a0 := eth.AddrInfo{
			//PrivkHex: "02e35ef9fd76747520ead76ddc8523d75a996a70508012a0b1153a8c580664e1",
			PrivkHex: "225be05b88b876e7f49226b0528d3bf27de2c316e395b4bcfdaf2589b33d78fa",
		}
		comtroller, err := unitroller.NewUnitroller(common.HexToAddress(comContract), client)
		if err != nil {
			fmt.Println(err)
			return
		}
		auth := bind.NewKeyedTransactor(a0.ToECDSAKey())

		cusdt, _ := cToken.NewCToken(common.HexToAddress(cUsdt), client)
		//dec, err := cusdt.Decimals(&bind.CallOpts{})
		//if err != nil {
		//	fmt.Println(err)
		//	return
		//}
		//0xb90c96607b45f9bB7509861A1CE77Cb8a72EdFB2
		borrowRate, err := cusdt.BorrowRatePerBlock(&bind.CallOpts{})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("borrowRate:", borrowRate)
		fmt.Println("from:", auth.From.String())
		_, liquility, shortFall, err := comtroller.GetAccountLiquidity(&bind.CallOpts{}, auth.From)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("liquility:", liquility)
		fmt.Println("shortFall:", shortFall)
		//borAmount:=big.NewInt(1)
		priceFeed, err := oracle.NewOracle(common.HexToAddress(OraclePrice), client)
		underlyPrice, err := priceFeed.GetUnderlyingPrice(&bind.CallOpts{}, common.HexToAddress(cUsdt))
		borrowAccount := toWei(decimal.NewFromBigInt(liquility, 0).Div(decimal.NewFromBigInt(underlyPrice, 0)), 6)
		fmt.Println("borrow Account:", borrowAccount)
		//borrowTrans, err := cusdt.Borrow(auth, mul(decimal.NewFromFloat(0.1),borrowRate))
		borrowTrans, err := cusdt.Borrow(auth, borrowAccount.BigInt())
		if err != nil {
			fmt.Println(err)
			return
		}
		time.Sleep(20 * time.Second)
		recBow, err := client.TransactionReceipt(context.Background(), borrowTrans.Hash())
		if err != nil {
			fmt.Println(err)
			return
		}
		if recBow.Status == 1 {
			fmt.Println("borrow finished:", recBow.TxHash.String())
		}

		_, liquility1, _, err := comtroller.GetAccountLiquidity(&bind.CallOpts{}, auth.From)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("liquility1:", liquility1)*/
}
func TestLiqtor(t *testing.T) {
	//ctoken(cEth) mint

	client, err := ethclient.Dial("https://ropsten.infura.io/v3/85c51263825545bf8496006327bd98d1")
	if err != nil {
		fmt.Println(err)
		return
	}

	a0 := eth.AddrInfo{
		//PrivkHex: "e97e8f8f7ff2d1a10830838343f26f46abce6de9c586a02d0acbbb80002ffdc6",
		PrivkHex: "225be05b88b876e7f49226b0528d3bf27de2c316e395b4bcfdaf2589b33d78fa",
	}
	auth := bind.NewKeyedTransactor(a0.ToECDSAKey())
	bal, _ := client.BalanceAt(context.Background(), auth.From, nil)
	fmt.Println(bal)

	comtroller, err := unitroller.NewUnitroller(common.HexToAddress(comContract), client)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("from:", auth.From)
	auth.GasLimit = 800000
	borrower := common.HexToAddress("0x9D41635F16D100fC0A77C942E1e3010B9962A5Ad")
	lqtor, _ := liquitor.NewLiquidtor(common.HexToAddress(liqContract), client)
	_, liq, _, err := comtroller.GetAccountLiquidity(&bind.CallOpts{}, borrower)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("liquility:", liq)
	//if liq.int64() != int64(0) {
	//	return
	//}

	trans, err := lqtor.LiquidateS(auth, borrower, common.HexToAddress("0x39AA39c021dfbaE8faC545936693aC917d5E7563"), common.HexToAddress("0x39AA39c021dfbaE8faC545936693aC917d5E7563"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(trans.Hash().String())
}

func TestWithDraw(t *testing.T) {
	//ctoken(cEth) mint
	//comptroller enter markets (ctoken = cEth)
	//ctoken(usdt) borrow

	//const host = "https://ropsten.infura.io/v3/fcfaf99bc9f94b148a65108207306f9e"
	const host = "wss://ropsten.infura.io/ws/v3/bcb800e6ebc0432ea9aac3c68b15577f"
	//const liqContract = "0x923eca9a79bf9e8a06dc5d160af3a007068f2a7d"
	const comContract = "0xcfa7b0e37f5AC60f3ae25226F5e39ec59AD26152"
	const cUsdt = "0xF6958Cf3127e62d3EB26c79F4f45d3F3b2CcdeD4"
	const cEth = "0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8"
	const OraclePrice = "0xb90c96607b45f9bb7509861a1ce77cb8a72edfb2"
	//const liqContract = "0x923eca9a79bf9e8a06dc5d160af3a007068f2a7d"
	const liqContract = "0x1d4E93b4C6b7C533E4d5e3F560B4F8D4e278319b"

	client, err := ethclient.Dial(host)
	if err != nil {
		fmt.Println(err)
		return
	}

	a0 := eth.AddrInfo{
		//PrivkHex: "fa243c65a9ec44ce2b822f87a0aaa2888f2fb01d365098346bfe87f7c7821b2a",
		PrivkHex: "225be05b88b876e7f49226b0528d3bf27de2c316e395b4bcfdaf2589b33d78fa",
	}
	auth := bind.NewKeyedTransactor(a0.ToECDSAKey())
	auth.GasPrice = big.NewInt(30 * 1000000000)
	auth.GasLimit = 1000000
	lqtor, _ := liquitor.NewLiquidtor(common.HexToAddress(liqContract), client)
	trans, err := lqtor.Withdraw(auth, common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab"))
	fmt.Println(trans.Hash().String())
}
