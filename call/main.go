package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dabankio/wallet-core/qa/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"liquidtor/Erc20"
	unitroller "liquidtor/Unitroller"
	"liquidtor/cToken"
	"liquidtor/db"
	"runtime"

	//liquitor "liquidtor"
	"liquidtor/liquitor"
	"math/big"
	"os"
	"strings"
	"sync"
	"time"
)

const fatory = "0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f"

var host = "wss://mainnet.infura.io/ws/v3/fcfaf99bc9f94b148a65108207306f9e"

const liqContract = "0x812843fe3dDeABd672121705ff090f49DbAbf22d" //
const comContract = "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b"
const cTokenContract = ""

var lk *sync.RWMutex
var clk *sync.RWMutex
var gol *sync.RWMutex
var minLk *sync.RWMutex
var lqLk *sync.RWMutex
var minBuf map[string]bool
var liqBuf map[string]bool
var index = 0
var SuggestGasPrice = big.NewInt(50 * 1000000000)

var isTest = false

type ContractNet struct {
	Liquitor *liquitor.Liquidtor
	//	Ctoken     *cToken.CToken
	Ctoken     map[string]*cToken.CToken
	Controller *unitroller.Unitroller
	Client     *ethclient.Client
}

var tokens map[string]*big.Int

type CTokenConfig struct {
	Contract string          `json:"contract"`
	Symbol   string          `json:"symbol"`
	Min      decimal.Decimal `json:"min"`
}
type Config struct {
	Compound []CTokenConfig `json:"compound"`
}

var golCtokens map[string]CTokenConfig

func InitNet(contractArray []string) *ContractNet {
	var rpcHost = host
	client, err := ethclient.Dial(rpcHost)
	if err != nil {
		fmt.Println("ethclient.Dial:", err)
		return nil
	}

	lqtor, _ := liquitor.NewLiquidtor(common.HexToAddress(liqContract), client)
	controller, _ := unitroller.NewUnitroller(common.HexToAddress(comContract), client)

	net := &ContractNet{
		Liquitor:   lqtor,
		Controller: controller,
		Ctoken:     make(map[string]*cToken.CToken, 0),
		Client:     client,
	}

	for _, con := range contractArray {
		oneCtoken, _ := cToken.NewCToken(common.HexToAddress(con), client)
		net.Ctoken[strings.ToUpper(con)] = oneCtoken
	}
	//进入市场列表
	//controller.GetAssetsIn()
	//通过ctoken 查询borrower的资产余额，就知道抵押物了
	return net
}
func LiquitorInit(lqtor *liquitor.Liquidtor) {
	var _comptrollerAddress, cETH, wETH, router, factory common.Address
	if isTest {
		_comptrollerAddress = common.HexToAddress("0xcfa7b0e37f5AC60f3ae25226F5e39ec59AD26152")
		cETH = common.HexToAddress("0x859e9d8a4edadfEDb5A2fF311243af80F85A91b8")
		wETH = common.HexToAddress("0xc778417e063141139fce010982780140aa0cd5ab")
		router = common.HexToAddress("0xe592427a0aece92de3edee1f18e0157c05861564") // uniswap
		//0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f
		factory = common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f") //uniswap factor
	} else {
		_comptrollerAddress = common.HexToAddress(comContract)
		cETH = common.HexToAddress("0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5")
		wETH = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
		router = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D") // uniswap
		//0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f
		factory = common.HexToAddress("0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f") //uniswap factor
	}

	a0 := eth.AddrInfo{
		PrivkHex: "e97e8f8f7ff2d1a10830838343f26f46abce6de9c586a02d0acbbb80002ffdc6",
		//PrivkHex: "225be05b88b876e7f49226b0528d3bf27de2c316e395b4bcfdaf2589b33d78fa",
	}
	fmt.Println(cETH, wETH, router, factory)
	auth := bind.NewKeyedTransactor(a0.ToECDSAKey())
	fmt.Println("addr:", auth.From.String())

	//fmt.Println(cETH, wETH, router, factory)
	tr, err := lqtor.SetComptroller(auth, _comptrollerAddress)
	if err != nil {
		fmt.Println("SetParam err:", err)
		return
	}
	fmt.Println("tr0", tr.Hash().String())

	//tr1, err := lqtor.SetParam(auth, cETH, wETH, router, factory)
	//if err != nil {
	//	fmt.Println("SetParam err:", err)
	//	return
	//}
	//

	//fmt.Println("tr1:", tr1.Hash().String())
}

//client, err := ethclient.Dial(host)
//if err != nil {
//log.Println(err)
//}
func loopLiquit(task chan db.Borrow, net *ContractNet, exeLiquit chan db.Borrow, ctx context.Context, ex chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("loopLiquit ctx.Done")
			return
		case t := <-task:
			var underToken string
			var balance *big.Int
			var decimal int32
			//fmt.Println("------loopLiquit-------",len(task))

			lk.RLock()
			if oneToken, ok := net.Ctoken[strings.ToUpper(t.ContractAddress)]; ok {
				lk.RUnlock()
				bal, err := oneToken.BorrowBalanceStored(&bind.CallOpts{}, common.HexToAddress(t.Address))
				if err != nil {
					fmt.Println("BorrowBalanceStored:", err)
					ex <- 1
					return
				}
				balance = bal
				token, _ := oneToken.Underlying(&bind.CallOpts{})
				//if err != nil {
				//	fmt.Println("Underlying:", err)
				//	ex <- 1
				//	return
				//}
				underToken = token.String()
				clk.RLock()
				if tk, ok := tokens[strings.ToUpper(underToken)]; ok {
					clk.RUnlock()
					decimal = int32(tk.Uint64())
				} else {
					clk.RUnlock()
					etoken, err := Erc20.NewErc20(token, net.Client)
					if err != nil {
						fmt.Println("Erc20.NewErc20:", err)
						ex <- 1
						return
					}
					dec, err := etoken.Decimals(&bind.CallOpts{})
					if err != nil {
						//	fmt.Println("etoken.Decimals:",err,token.String(),t.ContractAddress)
						//	continue
						dec = big.NewInt(18)
					}
					decimal = int32(dec.Uint64())
					clk.Lock()
					tokens[strings.ToUpper(underToken)] = dec
					clk.Unlock()
				}

			} else {
				lk.RUnlock()
			}

			ebal := ToEther(balance, decimal)
			gol.RLock()
			if cf, ok := golCtokens[strings.ToUpper(t.ContractAddress)]; ok {
				gol.RUnlock()
				if ebal.LessThan(cf.Min) {
					//fmt.Println("balance:", balance, "decimal:", decimal, "ebal:", ebal, "cf.Min:", cf.Min)
					//minLk.Lock()
					//key := strings.ToUpper(t.Address + "_" + t.ContractAddress)
					//minBuf[key] = true
					//minLk.Unlock()
					go db.UpdateLiquility("0", t.Address, t.ContractAddress)
					continue
				}

			} else {
				lk.RUnlock()
			}

			_, liq, _, err := net.Controller.GetAccountLiquidity(&bind.CallOpts{}, common.HexToAddress(t.Address))
			if err != nil {
				fmt.Println("GetAccountLiquidity:", err)
				ex <- 1
				return
			}

			fmt.Println("GetAccountLiquidity:addr,liq", t.Address, liq)

			go func() {

				gas, _ := net.Client.SuggestGasPrice(context.Background())
				if gas != nil {
					minLk.Lock()
					SuggestGasPrice = gas
					minLk.Unlock()
				}

			}()

			//fmt.Println("GetAccountLiquidity:",liq)
			if liq.Uint64() == 0 {
				exeLiquit <- t
				lqLk.Lock()
				key := strings.ToUpper(t.Address + "_" + t.ContractAddress)
				liqBuf[key] = true
				lqLk.Unlock()
			} else {
				task <- t
			}

		}
		time.Sleep(1 * time.Second)
	}
}

func LiquitorHandler(contract []string, hosts []string) {
	host = hosts[index]
	index += 1
	if index == len(hosts) {
		index = 0
	}
	net := InitNet(contract)
	if net == nil {
		LiquitorHandler(contract, hosts)
		return
	}
	lk = new(sync.RWMutex)
	minLk = new(sync.RWMutex)
	lqLk = new(sync.RWMutex)
	clk = new(sync.RWMutex)
	gol = new(sync.RWMutex)
	minBuf = make(map[string]bool, 0)
	liqBuf = make(map[string]bool, 0)

	task := make(chan db.Borrow, 200)
	exeLiquit := make(chan db.Borrow, 100)
	excption := make(chan int, 0)

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				return

			default:
				if len(task) <= 60 {
					st := 0
					data := db.GetRows(st)
					//fmt.Println("db.GetRows:",len(data))
					for i, d := range data {
						key := strings.ToUpper(d.Address + "_" + d.ContractAddress)
						lqLk.RLock()
						if _, ex := liqBuf[key]; !ex {
							lqLk.RUnlock()
							task <- d
						} else {
							lqLk.RUnlock()
						}
						fmt.Println(i)
					}
				}
				fmt.Println("len(task)", len(task))

			}
			time.Sleep(1 * time.Second)
		}

	}(ctx)

	for i := 0; i < 5; i++ {
		//get liquility
		go loopLiquit(task, net, exeLiquit, ctx, excption)
	}
	go func(ctx context.Context) {
		for {

			select {
			case <-ctx.Done():
				return
			case t := <-exeLiquit:
				txId := ExecLiquitor(net, t.Address, t.ContractAddress)
				if len(txId) == 0 {
					go db.UpdateLiquility("0", t.Address, t.ContractAddress)
					continue
				}
				go func(tx, addr, contract string, ctx context.Context) {
					for {
						select {
						case <-ctx.Done():
							return
						default:
							rec, err := net.Client.TransactionReceipt(context.Background(), common.HexToHash(tx))
							if err != nil {
								fmt.Println("TransactionReceipt:err,txId", err, tx)
								time.Sleep(2 * time.Second)
								continue
							}

							if rec != nil {
								if rec.Status > 0 {
									fmt.Println("succ:", tx)
								} else {
									fmt.Println("fail:", tx)
								}
								db.UpdateLiquility("0", addr, contract)
								return
							}

						}
					}

				}(txId, t.Address, t.ContractAddress, ctx)


			}
		}
	}(ctx)

	<-excption
	cancel()
	LiquitorHandler(contract, hosts)
}

func ExecLiquitor(net *ContractNet, borrower, repayToken string) string {
	cTokens, err := net.Controller.GetAssetsIn(&bind.CallOpts{}, common.HexToAddress(borrower))
	if err != nil {
		fmt.Println("GetAssetsIn", err)
		return ""
	}
	seizeToken, err := GetSnapShot(borrower, cTokens, net.Client)
	a0 := eth.AddrInfo{
		PrivkHex: "e97e8f8f7ff2d1a10830838343f26f46abce6de9c586a02d0acbbb80002ffdc6",
		//PrivkHex: "225be05b88b876e7f49226b0528d3bf27de2c316e395b4bcfdaf2589b33d78fa",
	}
	//price, err := net.Client.SuggestGasPrice(context.Background())

	auth := bind.NewKeyedTransactor(a0.ToECDSAKey())

	minLk.RLock()
	price := decimal.NewFromBigInt(SuggestGasPrice, 0).Mul(decimal.NewFromFloat(1.5)).BigInt()
	auth.GasPrice = price
	minLk.RUnlock()
	trans, err := net.Liquitor.LiquidateS(auth, common.HexToAddress(borrower), common.HexToAddress(repayToken), seizeToken)
	if err != nil {
		fmt.Println("LiquidateS", err)
		fmt.Println("LiquidateS failed: borrower,repayToken,seizeToken ", borrower, repayToken, seizeToken)
		return ""
	}
	return trans.Hash().String()
}

func GetBorrowAmount(borrower, cTokenAddress string, client *ethclient.Client) (*big.Int, error) {
	oneToken, _ := cToken.NewCToken(common.HexToAddress(cTokenAddress), client)
	borrowAmount, err := oneToken.BorrowBalanceStored(&bind.CallOpts{}, common.HexToAddress(borrower))
	fmt.Println(err)
	return borrowAmount, err
}
func GetSnapShot(borrower string, cTokenAddresses []common.Address, client *ethclient.Client) (suplyToken common.Address, err1 error) {
	//var maxSuply *big.Int
	maxSuply := big.NewInt(0)
	var seizeToken common.Address
	for _, cTokenAddress := range cTokenAddresses {
		oneToken, _ := cToken.NewCToken(cTokenAddress, client)
		_, suply, _, _, err := oneToken.GetAccountSnapshot(&bind.CallOpts{}, common.HexToAddress(borrower))
		if maxSuply.Cmp(suply) < 0 {
			maxSuply = suply
			seizeToken = cTokenAddress
		}
		err1 = err
	}

	return seizeToken, err1
}

//看那个抵押物余额最多
func GetMaxBalCToken(cTokens []common.Address, borrower string, client *ethclient.Client) common.Address {
	maxBal := decimal.Zero
	var maxAddr common.Address
	for _, token := range cTokens {
		oneToken, _ := cToken.NewCToken(token, client)
		decimal, _ := oneToken.Decimals(&bind.CallOpts{})
		bal, _ := oneToken.BalanceOf(&bind.CallOpts{}, common.HexToAddress(borrower))
		b := ToEther(bal, int32(decimal))
		if b.GreaterThan(maxBal) {
			maxAddr = token
			maxBal = b
		}
	}
	return maxAddr
}
func InitConfigFromJson() []string {
	// 打开文件
	file, _ := os.Open("config.json")
	// 关闭文件
	defer file.Close()
	//NewDecoder创建一个从file读取并解码json对象的*Decoder，解码器有自己的缓冲，并可能超前读取部分json数据。
	decoder := json.NewDecoder(file)
	//Decode从输入流读取下一个json编码值并保存在v指向的值里
	//var ctokens []CTokenConfig
	ctokens := make([]CTokenConfig, 0)
	cfg := Config{
		Compound: ctokens,
	}
	err := decoder.Decode(&cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	golCtokens = make(map[string]CTokenConfig, 0)
	contractArray := make([]string, 0)
	for _, v := range cfg.Compound {
		golCtokens[strings.ToUpper(v.Contract)] = v
		contractArray = append(contractArray, v.Contract)
	}
	return contractArray
}
func ToEther(input *big.Int, dec int32) decimal.Decimal {
	a := decimal.NewFromBigInt(input, 0)
	b := a.Div(decimal.New(1, dec))
	return b
}
func main() {
	//cEth := "0x4ddc2d193948926d02f9b1fe9e1daa0718270ed5"
	//cDai := "0x5d3a536e4d6dbd6114cc1ead35777bab948e3643"
	//cWbtc := "0xc11b1268c1a384e55c48c2391d8d480264a3a7f4"
	//cUsdc := "0x39aa39c021dfbae8fac545936693ac917d5e7563"
	//cUsdt := "0xf650c3d88d12db855b8bf7d11be6c55a4e07dcc9"
	//cBat := "0x6c8c6b02e7b2be14d4fa6022dfd6d75921d90e4e"
	//cComp := "0x70e36f6bf80a52b3b46b3af8e106cc0ed743e8e4"
	//cLink := "0xface851a4921ce59e912d19329929ce6da6eb0c7"
	//cRep := "0x158079ee67fce2f58472a96584a73c7ab9ac95c1"
	//cSai := "0xf5dce57282a584d2746faf1593d3121fcac444dc"
	//cTusd := "0x12392f67bdf24fae0af363c24ac620a2f67dad86"
	//cUni := "0x35a18000230da775cac24873d00ff85bccded550"
	//cZrx := "0xb3319f5d18bc0d84dd1b4825dcde5d5f7266d407"
	//cWbtc2 := "0xccf4429db6322d5c611ee964527d42e5d685dd6a"
	//contractArray := []string{cEth, cDai, cWbtc, cUsdc, cUsdt, cBat, cComp, cLink, cRep, cSai, cTusd, cUni, cZrx, cWbtc2}

	tokens = make(map[string]*big.Int, 0)

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)

	hostArray := []string{"wss://mainnet.infura.io/ws/v3/75c1d628b1d641ef9d9dac237bfe61a5", "wss://mainnet.infura.io/ws/v3/1c5430dcaba64f2a9d70432ef89aabe6", "wss://mainnet.infura.io/ws/v3/e3f31f4687294d2f904822855239035f", "wss://mainnet.infura.io/ws/v3/208d78944bbe48e393bccbf20ef494a7", "wss://mainnet.infura.io/ws/v3/630b59e0d28a4bcdafb1a0c52640f522", "wss://mainnet.infura.io/ws/v3/abc17edccebb4ff4a9318b848170faeb", "wss://mainnet.infura.io/ws/v3/fcfaf99bc9f94b148a65108207306f9e", "wss://mainnet.infura.io/ws/v3/85c51263825545bf8496006327bd98d1", "wss://mainnet.infura.io/ws/v3/2719102f50be441d89e32a525c902d85", "wss://mainnet.infura.io/ws/v3/4f685300a8d746209b81c96eb8b0e82d", "wss://mainnet.infura.io/ws/v3/bcb800e6ebc0432ea9aac3c68b15577f", "wss://mainnet.infura.io/ws/v3/6c3cec5a40674f7b9d4cdb293550469e", "wss://mainnet.infura.io/ws/v3/a8371dbc3fb244fb8abaf7126ab8f915"}
	contractArray := InitConfigFromJson()
	if contractArray == nil{
		return
	}
	go LiquitorHandler(contractArray, hostArray)
	//go cToken.FilOne(mainet, "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b", contractArray, 10234467)
	go cToken.GetEvent(host, "0x3d9819210a31b4961b30ef54be2aed79b9c9cd3b", contractArray, hostArray)
	fmt.Println("finish!")
	select {}
}
