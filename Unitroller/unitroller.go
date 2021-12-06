// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unitroller

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// UnitrollerABI is the input ABI used to generate the binding from.
const UnitrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CompGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"CompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"ContributorCompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPausedAlis\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_grantComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPausedAlias\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"compSpeed\",\"type\":\"uint256\"}],\"name\":\"_setCompSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"compSpeed\",\"type\":\"uint256\"}],\"name\":\"_setContributorCompSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compContributorSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastContributorBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isComped\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"contributor\",\"type\":\"address\"}],\"name\":\"updateContributorRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Unitroller is an auto generated Go binding around an Ethereum contract.
type Unitroller struct {
	UnitrollerCaller     // Read-only binding to the contract
	UnitrollerTransactor // Write-only binding to the contract
	UnitrollerFilterer   // Log filterer for contract events
}

// UnitrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnitrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnitrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnitrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnitrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnitrollerSession struct {
	Contract     *Unitroller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnitrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnitrollerCallerSession struct {
	Contract *UnitrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnitrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnitrollerTransactorSession struct {
	Contract     *UnitrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnitrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnitrollerRaw struct {
	Contract *Unitroller // Generic contract binding to access the raw methods on
}

// UnitrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnitrollerCallerRaw struct {
	Contract *UnitrollerCaller // Generic read-only contract binding to access the raw methods on
}

// UnitrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnitrollerTransactorRaw struct {
	Contract *UnitrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnitroller creates a new instance of Unitroller, bound to a specific deployed contract.
func NewUnitroller(address common.Address, backend bind.ContractBackend) (*Unitroller, error) {
	contract, err := bindUnitroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unitroller{UnitrollerCaller: UnitrollerCaller{contract: contract}, UnitrollerTransactor: UnitrollerTransactor{contract: contract}, UnitrollerFilterer: UnitrollerFilterer{contract: contract}}, nil
}

// NewUnitrollerCaller creates a new read-only instance of Unitroller, bound to a specific deployed contract.
func NewUnitrollerCaller(address common.Address, caller bind.ContractCaller) (*UnitrollerCaller, error) {
	contract, err := bindUnitroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnitrollerCaller{contract: contract}, nil
}

// NewUnitrollerTransactor creates a new write-only instance of Unitroller, bound to a specific deployed contract.
func NewUnitrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*UnitrollerTransactor, error) {
	contract, err := bindUnitroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnitrollerTransactor{contract: contract}, nil
}

// NewUnitrollerFilterer creates a new log filterer instance of Unitroller, bound to a specific deployed contract.
func NewUnitrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*UnitrollerFilterer, error) {
	contract, err := bindUnitroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnitrollerFilterer{contract: contract}, nil
}

// bindUnitroller binds a generic wrapper to an already deployed contract.
func bindUnitroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnitrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unitroller *UnitrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unitroller.Contract.UnitrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unitroller *UnitrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unitroller.Contract.UnitrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unitroller *UnitrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unitroller.Contract.UnitrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unitroller *UnitrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unitroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unitroller *UnitrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unitroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unitroller *UnitrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unitroller.Contract.contract.Transact(opts, method, params...)
}

// BorrowGuardianPausedAlis is a free data retrieval call binding the contract method 0x74b1913e.
//
// Solidity: function _borrowGuardianPausedAlis() view returns(bool)
func (_Unitroller *UnitrollerCaller) BorrowGuardianPausedAlis(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "_borrowGuardianPausedAlis")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPausedAlis is a free data retrieval call binding the contract method 0x74b1913e.
//
// Solidity: function _borrowGuardianPausedAlis() view returns(bool)
func (_Unitroller *UnitrollerSession) BorrowGuardianPausedAlis() (bool, error) {
	return _Unitroller.Contract.BorrowGuardianPausedAlis(&_Unitroller.CallOpts)
}

// BorrowGuardianPausedAlis is a free data retrieval call binding the contract method 0x74b1913e.
//
// Solidity: function _borrowGuardianPausedAlis() view returns(bool)
func (_Unitroller *UnitrollerCallerSession) BorrowGuardianPausedAlis() (bool, error) {
	return _Unitroller.Contract.BorrowGuardianPausedAlis(&_Unitroller.CallOpts)
}

// MintGuardianPausedAlias is a free data retrieval call binding the contract method 0x451dee99.
//
// Solidity: function _mintGuardianPausedAlias() view returns(bool)
func (_Unitroller *UnitrollerCaller) MintGuardianPausedAlias(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "_mintGuardianPausedAlias")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPausedAlias is a free data retrieval call binding the contract method 0x451dee99.
//
// Solidity: function _mintGuardianPausedAlias() view returns(bool)
func (_Unitroller *UnitrollerSession) MintGuardianPausedAlias() (bool, error) {
	return _Unitroller.Contract.MintGuardianPausedAlias(&_Unitroller.CallOpts)
}

// MintGuardianPausedAlias is a free data retrieval call binding the contract method 0x451dee99.
//
// Solidity: function _mintGuardianPausedAlias() view returns(bool)
func (_Unitroller *UnitrollerCallerSession) MintGuardianPausedAlias() (bool, error) {
	return _Unitroller.Contract.MintGuardianPausedAlias(&_Unitroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Unitroller *UnitrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Unitroller *UnitrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Unitroller.Contract.AccountAssets(&_Unitroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Unitroller *UnitrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Unitroller.Contract.AccountAssets(&_Unitroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Unitroller *UnitrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Unitroller *UnitrollerSession) Admin() (common.Address, error) {
	return _Unitroller.Contract.Admin(&_Unitroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Unitroller *UnitrollerCallerSession) Admin() (common.Address, error) {
	return _Unitroller.Contract.Admin(&_Unitroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Unitroller *UnitrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Unitroller *UnitrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Unitroller.Contract.AllMarkets(&_Unitroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Unitroller *UnitrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Unitroller.Contract.AllMarkets(&_Unitroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Unitroller *UnitrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Unitroller *UnitrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _Unitroller.Contract.BorrowCapGuardian(&_Unitroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Unitroller *UnitrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Unitroller.Contract.BorrowCapGuardian(&_Unitroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.BorrowCaps(&_Unitroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.BorrowCaps(&_Unitroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Unitroller *UnitrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Unitroller *UnitrollerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Unitroller.Contract.BorrowGuardianPaused(&_Unitroller.CallOpts, arg0)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Unitroller *UnitrollerCallerSession) BorrowGuardianPaused(arg0 common.Address) (bool, error) {
	return _Unitroller.Contract.BorrowGuardianPaused(&_Unitroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Unitroller *UnitrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, cToken common.Address) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "checkMembership", account, cToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Unitroller *UnitrollerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Unitroller.Contract.CheckMembership(&_Unitroller.CallOpts, account, cToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Unitroller *UnitrollerCallerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Unitroller.Contract.CheckMembership(&_Unitroller.CallOpts, account, cToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Unitroller *UnitrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Unitroller *UnitrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Unitroller.Contract.CloseFactorMantissa(&_Unitroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Unitroller.Contract.CloseFactorMantissa(&_Unitroller.CallOpts)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) CompAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompAccrued(&_Unitroller.CallOpts, arg0)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompAccrued(&_Unitroller.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Unitroller *UnitrollerCaller) CompBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Unitroller *UnitrollerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Unitroller.Contract.CompBorrowState(&_Unitroller.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Unitroller *UnitrollerCallerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Unitroller.Contract.CompBorrowState(&_Unitroller.CallOpts, arg0)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) CompBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompBorrowerIndex(&_Unitroller.CallOpts, arg0, arg1)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompBorrowerIndex(&_Unitroller.CallOpts, arg0, arg1)
}

// CompContributorSpeeds is a free data retrieval call binding the contract method 0x986ab838.
//
// Solidity: function compContributorSpeeds(address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) CompContributorSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compContributorSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompContributorSpeeds is a free data retrieval call binding the contract method 0x986ab838.
//
// Solidity: function compContributorSpeeds(address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) CompContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompContributorSpeeds(&_Unitroller.CallOpts, arg0)
}

// CompContributorSpeeds is a free data retrieval call binding the contract method 0x986ab838.
//
// Solidity: function compContributorSpeeds(address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CompContributorSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompContributorSpeeds(&_Unitroller.CallOpts, arg0)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Unitroller *UnitrollerCaller) CompInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Unitroller *UnitrollerSession) CompInitialIndex() (*big.Int, error) {
	return _Unitroller.Contract.CompInitialIndex(&_Unitroller.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Unitroller *UnitrollerCallerSession) CompInitialIndex() (*big.Int, error) {
	return _Unitroller.Contract.CompInitialIndex(&_Unitroller.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Unitroller *UnitrollerCaller) CompRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Unitroller *UnitrollerSession) CompRate() (*big.Int, error) {
	return _Unitroller.Contract.CompRate(&_Unitroller.CallOpts)
}

// CompRate is a free data retrieval call binding the contract method 0xaa900754.
//
// Solidity: function compRate() view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CompRate() (*big.Int, error) {
	return _Unitroller.Contract.CompRate(&_Unitroller.CallOpts)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) CompSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompSpeeds(&_Unitroller.CallOpts, arg0)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompSpeeds(&_Unitroller.CallOpts, arg0)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) CompSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompSupplierIndex(&_Unitroller.CallOpts, arg0, arg1)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.CompSupplierIndex(&_Unitroller.CallOpts, arg0, arg1)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Unitroller *UnitrollerCaller) CompSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "compSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Block = *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return *outstruct, err

}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Unitroller *UnitrollerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Unitroller.Contract.CompSupplyState(&_Unitroller.CallOpts, arg0)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Unitroller *UnitrollerCallerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Unitroller.Contract.CompSupplyState(&_Unitroller.CallOpts, arg0)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.ComptrollerImplementation(&_Unitroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.ComptrollerImplementation(&_Unitroller.CallOpts)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Unitroller *UnitrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Unitroller *UnitrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Unitroller.Contract.GetAccountLiquidity(&_Unitroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Unitroller *UnitrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Unitroller.Contract.GetAccountLiquidity(&_Unitroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Unitroller *UnitrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Unitroller *UnitrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _Unitroller.Contract.GetAllMarkets(&_Unitroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Unitroller *UnitrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Unitroller.Contract.GetAllMarkets(&_Unitroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Unitroller *UnitrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Unitroller *UnitrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Unitroller.Contract.GetAssetsIn(&_Unitroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Unitroller *UnitrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Unitroller.Contract.GetAssetsIn(&_Unitroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Unitroller *UnitrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Unitroller *UnitrollerSession) GetBlockNumber() (*big.Int, error) {
	return _Unitroller.Contract.GetBlockNumber(&_Unitroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Unitroller.Contract.GetBlockNumber(&_Unitroller.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Unitroller *UnitrollerCaller) GetCompAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "getCompAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Unitroller *UnitrollerSession) GetCompAddress() (common.Address, error) {
	return _Unitroller.Contract.GetCompAddress(&_Unitroller.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Unitroller *UnitrollerCallerSession) GetCompAddress() (common.Address, error) {
	return _Unitroller.Contract.GetCompAddress(&_Unitroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Unitroller *UnitrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, cTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Unitroller *UnitrollerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Unitroller.Contract.GetHypotheticalAccountLiquidity(&_Unitroller.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Unitroller *UnitrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Unitroller.Contract.GetHypotheticalAccountLiquidity(&_Unitroller.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Unitroller *UnitrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Unitroller *UnitrollerSession) IsComptroller() (bool, error) {
	return _Unitroller.Contract.IsComptroller(&_Unitroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Unitroller *UnitrollerCallerSession) IsComptroller() (bool, error) {
	return _Unitroller.Contract.IsComptroller(&_Unitroller.CallOpts)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Unitroller *UnitrollerCaller) LastContributorBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "lastContributorBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Unitroller *UnitrollerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.LastContributorBlock(&_Unitroller.CallOpts, arg0)
}

// LastContributorBlock is a free data retrieval call binding the contract method 0xbea6b8b8.
//
// Solidity: function lastContributorBlock(address ) view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) LastContributorBlock(arg0 common.Address) (*big.Int, error) {
	return _Unitroller.Contract.LastContributorBlock(&_Unitroller.CallOpts, arg0)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Unitroller *UnitrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Unitroller *UnitrollerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Unitroller.Contract.LiquidateCalculateSeizeTokens(&_Unitroller.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Unitroller *UnitrollerCallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Unitroller.Contract.LiquidateCalculateSeizeTokens(&_Unitroller.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Unitroller *UnitrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Unitroller *UnitrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Unitroller.Contract.LiquidationIncentiveMantissa(&_Unitroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Unitroller.Contract.LiquidationIncentiveMantissa(&_Unitroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Unitroller *UnitrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
		IsComped                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsListed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.CollateralFactorMantissa = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsComped = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Unitroller *UnitrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _Unitroller.Contract.Markets(&_Unitroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa, bool isComped)
func (_Unitroller *UnitrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
	IsComped                 bool
}, error) {
	return _Unitroller.Contract.Markets(&_Unitroller.CallOpts, arg0)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Unitroller *UnitrollerCaller) MaxAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "maxAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Unitroller *UnitrollerSession) MaxAssets() (*big.Int, error) {
	return _Unitroller.Contract.MaxAssets(&_Unitroller.CallOpts)
}

// MaxAssets is a free data retrieval call binding the contract method 0x94b2294b.
//
// Solidity: function maxAssets() view returns(uint256)
func (_Unitroller *UnitrollerCallerSession) MaxAssets() (*big.Int, error) {
	return _Unitroller.Contract.MaxAssets(&_Unitroller.CallOpts)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Unitroller *UnitrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Unitroller *UnitrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Unitroller.Contract.MintGuardianPaused(&_Unitroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Unitroller *UnitrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Unitroller.Contract.MintGuardianPaused(&_Unitroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Unitroller *UnitrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Unitroller *UnitrollerSession) Oracle() (common.Address, error) {
	return _Unitroller.Contract.Oracle(&_Unitroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Unitroller *UnitrollerCallerSession) Oracle() (common.Address, error) {
	return _Unitroller.Contract.Oracle(&_Unitroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Unitroller *UnitrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Unitroller *UnitrollerSession) PauseGuardian() (common.Address, error) {
	return _Unitroller.Contract.PauseGuardian(&_Unitroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Unitroller *UnitrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Unitroller.Contract.PauseGuardian(&_Unitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Unitroller *UnitrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Unitroller *UnitrollerSession) PendingAdmin() (common.Address, error) {
	return _Unitroller.Contract.PendingAdmin(&_Unitroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Unitroller *UnitrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Unitroller.Contract.PendingAdmin(&_Unitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.PendingComptrollerImplementation(&_Unitroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Unitroller *UnitrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Unitroller.Contract.PendingComptrollerImplementation(&_Unitroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Unitroller *UnitrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Unitroller *UnitrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Unitroller.Contract.SeizeGuardianPaused(&_Unitroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Unitroller *UnitrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Unitroller.Contract.SeizeGuardianPaused(&_Unitroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Unitroller *UnitrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Unitroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Unitroller *UnitrollerSession) TransferGuardianPaused() (bool, error) {
	return _Unitroller.Contract.TransferGuardianPaused(&_Unitroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Unitroller *UnitrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Unitroller.Contract.TransferGuardianPaused(&_Unitroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Unitroller *UnitrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Unitroller *UnitrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.Become(&_Unitroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Unitroller *UnitrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.Become(&_Unitroller.TransactOpts, unitroller)
}

// GrantComp is a paid mutator transaction binding the contract method 0x27efe3cb.
//
// Solidity: function _grantComp(address recipient, uint256 amount) returns()
func (_Unitroller *UnitrollerTransactor) GrantComp(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_grantComp", recipient, amount)
}

// GrantComp is a paid mutator transaction binding the contract method 0x27efe3cb.
//
// Solidity: function _grantComp(address recipient, uint256 amount) returns()
func (_Unitroller *UnitrollerSession) GrantComp(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.GrantComp(&_Unitroller.TransactOpts, recipient, amount)
}

// GrantComp is a paid mutator transaction binding the contract method 0x27efe3cb.
//
// Solidity: function _grantComp(address recipient, uint256 amount) returns()
func (_Unitroller *UnitrollerTransactorSession) GrantComp(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.GrantComp(&_Unitroller.TransactOpts, recipient, amount)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Unitroller *UnitrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Unitroller *UnitrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetBorrowCapGuardian(&_Unitroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Unitroller *UnitrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetBorrowCapGuardian(&_Unitroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Unitroller *UnitrollerTransactor) SetBorrowPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setBorrowPaused", cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Unitroller *UnitrollerSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetBorrowPaused(&_Unitroller.TransactOpts, cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Unitroller *UnitrollerTransactorSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetBorrowPaused(&_Unitroller.TransactOpts, cToken, state)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Unitroller *UnitrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetCloseFactor(&_Unitroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetCloseFactor(&_Unitroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setCollateralFactor", cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Unitroller *UnitrollerSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetCollateralFactor(&_Unitroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetCollateralFactor(&_Unitroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCompSpeed is a paid mutator transaction binding the contract method 0x434caf25.
//
// Solidity: function _setCompSpeed(address cToken, uint256 compSpeed) returns()
func (_Unitroller *UnitrollerTransactor) SetCompSpeed(opts *bind.TransactOpts, cToken common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setCompSpeed", cToken, compSpeed)
}

// SetCompSpeed is a paid mutator transaction binding the contract method 0x434caf25.
//
// Solidity: function _setCompSpeed(address cToken, uint256 compSpeed) returns()
func (_Unitroller *UnitrollerSession) SetCompSpeed(cToken common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetCompSpeed(&_Unitroller.TransactOpts, cToken, compSpeed)
}

// SetCompSpeed is a paid mutator transaction binding the contract method 0x434caf25.
//
// Solidity: function _setCompSpeed(address cToken, uint256 compSpeed) returns()
func (_Unitroller *UnitrollerTransactorSession) SetCompSpeed(cToken common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetCompSpeed(&_Unitroller.TransactOpts, cToken, compSpeed)
}

// SetContributorCompSpeed is a paid mutator transaction binding the contract method 0x598ee1cb.
//
// Solidity: function _setContributorCompSpeed(address contributor, uint256 compSpeed) returns()
func (_Unitroller *UnitrollerTransactor) SetContributorCompSpeed(opts *bind.TransactOpts, contributor common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setContributorCompSpeed", contributor, compSpeed)
}

// SetContributorCompSpeed is a paid mutator transaction binding the contract method 0x598ee1cb.
//
// Solidity: function _setContributorCompSpeed(address contributor, uint256 compSpeed) returns()
func (_Unitroller *UnitrollerSession) SetContributorCompSpeed(contributor common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetContributorCompSpeed(&_Unitroller.TransactOpts, contributor, compSpeed)
}

// SetContributorCompSpeed is a paid mutator transaction binding the contract method 0x598ee1cb.
//
// Solidity: function _setContributorCompSpeed(address contributor, uint256 compSpeed) returns()
func (_Unitroller *UnitrollerTransactorSession) SetContributorCompSpeed(contributor common.Address, compSpeed *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetContributorCompSpeed(&_Unitroller.TransactOpts, contributor, compSpeed)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Unitroller *UnitrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetLiquidationIncentive(&_Unitroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetLiquidationIncentive(&_Unitroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Unitroller *UnitrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setMarketBorrowCaps", cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Unitroller *UnitrollerSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetMarketBorrowCaps(&_Unitroller.TransactOpts, cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Unitroller *UnitrollerTransactorSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SetMarketBorrowCaps(&_Unitroller.TransactOpts, cTokens, newBorrowCaps)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Unitroller *UnitrollerTransactor) SetMintPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setMintPaused", cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Unitroller *UnitrollerSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetMintPaused(&_Unitroller.TransactOpts, cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Unitroller *UnitrollerTransactorSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetMintPaused(&_Unitroller.TransactOpts, cToken, state)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Unitroller *UnitrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPauseGuardian(&_Unitroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPauseGuardian(&_Unitroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Unitroller *UnitrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPriceOracle(&_Unitroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SetPriceOracle(&_Unitroller.TransactOpts, newOracle)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Unitroller *UnitrollerTransactor) SetSeizePaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setSeizePaused", state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Unitroller *UnitrollerSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetSeizePaused(&_Unitroller.TransactOpts, state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Unitroller *UnitrollerTransactorSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetSeizePaused(&_Unitroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Unitroller *UnitrollerTransactor) SetTransferPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_setTransferPaused", state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Unitroller *UnitrollerSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetTransferPaused(&_Unitroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Unitroller *UnitrollerTransactorSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Unitroller.Contract.SetTransferPaused(&_Unitroller.TransactOpts, state)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SupportMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "_supportMarket", cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Unitroller *UnitrollerSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SupportMarket(&_Unitroller.TransactOpts, cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.SupportMarket(&_Unitroller.TransactOpts, cToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Unitroller *UnitrollerSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.BorrowAllowed(&_Unitroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.BorrowAllowed(&_Unitroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Unitroller *UnitrollerTransactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Unitroller *UnitrollerSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.BorrowVerify(&_Unitroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Unitroller *UnitrollerTransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.BorrowVerify(&_Unitroller.TransactOpts, cToken, borrower, borrowAmount)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Unitroller *UnitrollerTransactor) ClaimComp(opts *bind.TransactOpts, holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "claimComp", holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Unitroller *UnitrollerSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.ClaimComp(&_Unitroller.TransactOpts, holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Unitroller *UnitrollerTransactorSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.ClaimComp(&_Unitroller.TransactOpts, holder, cTokens)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Unitroller *UnitrollerTransactor) ClaimComp0(opts *bind.TransactOpts, holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "claimComp0", holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Unitroller *UnitrollerSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Unitroller.Contract.ClaimComp0(&_Unitroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Unitroller *UnitrollerTransactorSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Unitroller.Contract.ClaimComp0(&_Unitroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Unitroller *UnitrollerTransactor) ClaimComp1(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "claimComp1", holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Unitroller *UnitrollerSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.ClaimComp1(&_Unitroller.TransactOpts, holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Unitroller *UnitrollerTransactorSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.ClaimComp1(&_Unitroller.TransactOpts, holder)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Unitroller *UnitrollerTransactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Unitroller *UnitrollerSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.EnterMarkets(&_Unitroller.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Unitroller *UnitrollerTransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.EnterMarkets(&_Unitroller.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Unitroller *UnitrollerTransactor) ExitMarket(opts *bind.TransactOpts, cTokenAddress common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "exitMarket", cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Unitroller *UnitrollerSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.ExitMarket(&_Unitroller.TransactOpts, cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.ExitMarket(&_Unitroller.TransactOpts, cTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Unitroller *UnitrollerSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.LiquidateBorrowAllowed(&_Unitroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.LiquidateBorrowAllowed(&_Unitroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Unitroller *UnitrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Unitroller *UnitrollerSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.LiquidateBorrowVerify(&_Unitroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Unitroller *UnitrollerTransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.LiquidateBorrowVerify(&_Unitroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Unitroller *UnitrollerSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.MintAllowed(&_Unitroller.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.MintAllowed(&_Unitroller.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Unitroller *UnitrollerTransactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "mintVerify", cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Unitroller *UnitrollerSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.MintVerify(&_Unitroller.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Unitroller *UnitrollerTransactorSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.MintVerify(&_Unitroller.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Unitroller *UnitrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Unitroller *UnitrollerSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RedeemAllowed(&_Unitroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RedeemAllowed(&_Unitroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Unitroller *UnitrollerTransactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Unitroller *UnitrollerSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RedeemVerify(&_Unitroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Unitroller *UnitrollerTransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RedeemVerify(&_Unitroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Unitroller *UnitrollerSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RepayBorrowAllowed(&_Unitroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RepayBorrowAllowed(&_Unitroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Unitroller *UnitrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Unitroller *UnitrollerSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RepayBorrowVerify(&_Unitroller.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Unitroller *UnitrollerTransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.RepayBorrowVerify(&_Unitroller.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Unitroller *UnitrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Unitroller *UnitrollerSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SeizeAllowed(&_Unitroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SeizeAllowed(&_Unitroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Unitroller *UnitrollerTransactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Unitroller *UnitrollerSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SeizeVerify(&_Unitroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Unitroller *UnitrollerTransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.SeizeVerify(&_Unitroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Unitroller *UnitrollerTransactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Unitroller *UnitrollerSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.TransferAllowed(&_Unitroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Unitroller *UnitrollerTransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.TransferAllowed(&_Unitroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Unitroller *UnitrollerTransactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Unitroller *UnitrollerSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.TransferVerify(&_Unitroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Unitroller *UnitrollerTransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Unitroller.Contract.TransferVerify(&_Unitroller.TransactOpts, cToken, src, dst, transferTokens)
}

// UpdateContributorRewards is a paid mutator transaction binding the contract method 0x741b2525.
//
// Solidity: function updateContributorRewards(address contributor) returns()
func (_Unitroller *UnitrollerTransactor) UpdateContributorRewards(opts *bind.TransactOpts, contributor common.Address) (*types.Transaction, error) {
	return _Unitroller.contract.Transact(opts, "updateContributorRewards", contributor)
}

// UpdateContributorRewards is a paid mutator transaction binding the contract method 0x741b2525.
//
// Solidity: function updateContributorRewards(address contributor) returns()
func (_Unitroller *UnitrollerSession) UpdateContributorRewards(contributor common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.UpdateContributorRewards(&_Unitroller.TransactOpts, contributor)
}

// UpdateContributorRewards is a paid mutator transaction binding the contract method 0x741b2525.
//
// Solidity: function updateContributorRewards(address contributor) returns()
func (_Unitroller *UnitrollerTransactorSession) UpdateContributorRewards(contributor common.Address) (*types.Transaction, error) {
	return _Unitroller.Contract.UpdateContributorRewards(&_Unitroller.TransactOpts, contributor)
}

// UnitrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Unitroller contract.
type UnitrollerActionPausedIterator struct {
	Event *UnitrollerActionPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerActionPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerActionPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerActionPaused represents a ActionPaused event raised by the Unitroller contract.
type UnitrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Unitroller *UnitrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*UnitrollerActionPausedIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &UnitrollerActionPausedIterator{contract: _Unitroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Unitroller *UnitrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *UnitrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerActionPaused)
				if err := _Unitroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Unitroller *UnitrollerFilterer) ParseActionPaused(log types.Log) (*UnitrollerActionPaused, error) {
	event := new(UnitrollerActionPaused)
	if err := _Unitroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the Unitroller contract.
type UnitrollerActionPaused0Iterator struct {
	Event *UnitrollerActionPaused0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerActionPaused0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerActionPaused0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerActionPaused0 represents a ActionPaused0 event raised by the Unitroller contract.
type UnitrollerActionPaused0 struct {
	CToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Unitroller *UnitrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*UnitrollerActionPaused0Iterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &UnitrollerActionPaused0Iterator{contract: _Unitroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Unitroller *UnitrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *UnitrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerActionPaused0)
				if err := _Unitroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Unitroller *UnitrollerFilterer) ParseActionPaused0(log types.Log) (*UnitrollerActionPaused0, error) {
	event := new(UnitrollerActionPaused0)
	if err := _Unitroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerCompGrantedIterator is returned from FilterCompGranted and is used to iterate over the raw logs and unpacked data for CompGranted events raised by the Unitroller contract.
type UnitrollerCompGrantedIterator struct {
	Event *UnitrollerCompGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerCompGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerCompGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerCompGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerCompGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerCompGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerCompGranted represents a CompGranted event raised by the Unitroller contract.
type UnitrollerCompGranted struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompGranted is a free log retrieval operation binding the contract event 0x98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c.
//
// Solidity: event CompGranted(address recipient, uint256 amount)
func (_Unitroller *UnitrollerFilterer) FilterCompGranted(opts *bind.FilterOpts) (*UnitrollerCompGrantedIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "CompGranted")
	if err != nil {
		return nil, err
	}
	return &UnitrollerCompGrantedIterator{contract: _Unitroller.contract, event: "CompGranted", logs: logs, sub: sub}, nil
}

// WatchCompGranted is a free log subscription operation binding the contract event 0x98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c.
//
// Solidity: event CompGranted(address recipient, uint256 amount)
func (_Unitroller *UnitrollerFilterer) WatchCompGranted(opts *bind.WatchOpts, sink chan<- *UnitrollerCompGranted) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "CompGranted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerCompGranted)
				if err := _Unitroller.contract.UnpackLog(event, "CompGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompGranted is a log parse operation binding the contract event 0x98b2f82a3a07f223a0be64b3d0f47711c64dccd1feafb94aa28156b38cd9695c.
//
// Solidity: event CompGranted(address recipient, uint256 amount)
func (_Unitroller *UnitrollerFilterer) ParseCompGranted(log types.Log) (*UnitrollerCompGranted, error) {
	event := new(UnitrollerCompGranted)
	if err := _Unitroller.contract.UnpackLog(event, "CompGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerCompSpeedUpdatedIterator is returned from FilterCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for CompSpeedUpdated events raised by the Unitroller contract.
type UnitrollerCompSpeedUpdatedIterator struct {
	Event *UnitrollerCompSpeedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerCompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerCompSpeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerCompSpeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerCompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerCompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerCompSpeedUpdated represents a CompSpeedUpdated event raised by the Unitroller contract.
type UnitrollerCompSpeedUpdated struct {
	CToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompSpeedUpdated is a free log retrieval operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Unitroller *UnitrollerFilterer) FilterCompSpeedUpdated(opts *bind.FilterOpts, cToken []common.Address) (*UnitrollerCompSpeedUpdatedIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &UnitrollerCompSpeedUpdatedIterator{contract: _Unitroller.contract, event: "CompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchCompSpeedUpdated is a free log subscription operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Unitroller *UnitrollerFilterer) WatchCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *UnitrollerCompSpeedUpdated, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerCompSpeedUpdated)
				if err := _Unitroller.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCompSpeedUpdated is a log parse operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Unitroller *UnitrollerFilterer) ParseCompSpeedUpdated(log types.Log) (*UnitrollerCompSpeedUpdated, error) {
	event := new(UnitrollerCompSpeedUpdated)
	if err := _Unitroller.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerContributorCompSpeedUpdatedIterator is returned from FilterContributorCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for ContributorCompSpeedUpdated events raised by the Unitroller contract.
type UnitrollerContributorCompSpeedUpdatedIterator struct {
	Event *UnitrollerContributorCompSpeedUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerContributorCompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerContributorCompSpeedUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerContributorCompSpeedUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerContributorCompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerContributorCompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerContributorCompSpeedUpdated represents a ContributorCompSpeedUpdated event raised by the Unitroller contract.
type UnitrollerContributorCompSpeedUpdated struct {
	Contributor common.Address
	NewSpeed    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContributorCompSpeedUpdated is a free log retrieval operation binding the contract event 0x386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b3.
//
// Solidity: event ContributorCompSpeedUpdated(address indexed contributor, uint256 newSpeed)
func (_Unitroller *UnitrollerFilterer) FilterContributorCompSpeedUpdated(opts *bind.FilterOpts, contributor []common.Address) (*UnitrollerContributorCompSpeedUpdatedIterator, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "ContributorCompSpeedUpdated", contributorRule)
	if err != nil {
		return nil, err
	}
	return &UnitrollerContributorCompSpeedUpdatedIterator{contract: _Unitroller.contract, event: "ContributorCompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchContributorCompSpeedUpdated is a free log subscription operation binding the contract event 0x386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b3.
//
// Solidity: event ContributorCompSpeedUpdated(address indexed contributor, uint256 newSpeed)
func (_Unitroller *UnitrollerFilterer) WatchContributorCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *UnitrollerContributorCompSpeedUpdated, contributor []common.Address) (event.Subscription, error) {

	var contributorRule []interface{}
	for _, contributorItem := range contributor {
		contributorRule = append(contributorRule, contributorItem)
	}

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "ContributorCompSpeedUpdated", contributorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerContributorCompSpeedUpdated)
				if err := _Unitroller.contract.UnpackLog(event, "ContributorCompSpeedUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseContributorCompSpeedUpdated is a log parse operation binding the contract event 0x386537fa92edc3319af95f1f904dcf1900021e4f3f4e08169a577a09076e66b3.
//
// Solidity: event ContributorCompSpeedUpdated(address indexed contributor, uint256 newSpeed)
func (_Unitroller *UnitrollerFilterer) ParseContributorCompSpeedUpdated(log types.Log) (*UnitrollerContributorCompSpeedUpdated, error) {
	event := new(UnitrollerContributorCompSpeedUpdated)
	if err := _Unitroller.contract.UnpackLog(event, "ContributorCompSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerDistributedBorrowerCompIterator is returned from FilterDistributedBorrowerComp and is used to iterate over the raw logs and unpacked data for DistributedBorrowerComp events raised by the Unitroller contract.
type UnitrollerDistributedBorrowerCompIterator struct {
	Event *UnitrollerDistributedBorrowerComp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerDistributedBorrowerCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerDistributedBorrowerComp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerDistributedBorrowerComp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerDistributedBorrowerCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerDistributedBorrowerCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerDistributedBorrowerComp represents a DistributedBorrowerComp event raised by the Unitroller contract.
type UnitrollerDistributedBorrowerComp struct {
	CToken          common.Address
	Borrower        common.Address
	CompDelta       *big.Int
	CompBorrowIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerComp is a free log retrieval operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Unitroller *UnitrollerFilterer) FilterDistributedBorrowerComp(opts *bind.FilterOpts, cToken []common.Address, borrower []common.Address) (*UnitrollerDistributedBorrowerCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &UnitrollerDistributedBorrowerCompIterator{contract: _Unitroller.contract, event: "DistributedBorrowerComp", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerComp is a free log subscription operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Unitroller *UnitrollerFilterer) WatchDistributedBorrowerComp(opts *bind.WatchOpts, sink chan<- *UnitrollerDistributedBorrowerComp, cToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerDistributedBorrowerComp)
				if err := _Unitroller.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedBorrowerComp is a log parse operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Unitroller *UnitrollerFilterer) ParseDistributedBorrowerComp(log types.Log) (*UnitrollerDistributedBorrowerComp, error) {
	event := new(UnitrollerDistributedBorrowerComp)
	if err := _Unitroller.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerDistributedSupplierCompIterator is returned from FilterDistributedSupplierComp and is used to iterate over the raw logs and unpacked data for DistributedSupplierComp events raised by the Unitroller contract.
type UnitrollerDistributedSupplierCompIterator struct {
	Event *UnitrollerDistributedSupplierComp // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerDistributedSupplierCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerDistributedSupplierComp)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerDistributedSupplierComp)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerDistributedSupplierCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerDistributedSupplierCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerDistributedSupplierComp represents a DistributedSupplierComp event raised by the Unitroller contract.
type UnitrollerDistributedSupplierComp struct {
	CToken          common.Address
	Supplier        common.Address
	CompDelta       *big.Int
	CompSupplyIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierComp is a free log retrieval operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Unitroller *UnitrollerFilterer) FilterDistributedSupplierComp(opts *bind.FilterOpts, cToken []common.Address, supplier []common.Address) (*UnitrollerDistributedSupplierCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &UnitrollerDistributedSupplierCompIterator{contract: _Unitroller.contract, event: "DistributedSupplierComp", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierComp is a free log subscription operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Unitroller *UnitrollerFilterer) WatchDistributedSupplierComp(opts *bind.WatchOpts, sink chan<- *UnitrollerDistributedSupplierComp, cToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerDistributedSupplierComp)
				if err := _Unitroller.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDistributedSupplierComp is a log parse operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Unitroller *UnitrollerFilterer) ParseDistributedSupplierComp(log types.Log) (*UnitrollerDistributedSupplierComp, error) {
	event := new(UnitrollerDistributedSupplierComp)
	if err := _Unitroller.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Unitroller contract.
type UnitrollerFailureIterator struct {
	Event *UnitrollerFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerFailure represents a Failure event raised by the Unitroller contract.
type UnitrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Unitroller *UnitrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*UnitrollerFailureIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &UnitrollerFailureIterator{contract: _Unitroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Unitroller *UnitrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *UnitrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerFailure)
				if err := _Unitroller.contract.UnpackLog(event, "Failure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Unitroller *UnitrollerFilterer) ParseFailure(log types.Log) (*UnitrollerFailure, error) {
	event := new(UnitrollerFailure)
	if err := _Unitroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Unitroller contract.
type UnitrollerMarketEnteredIterator struct {
	Event *UnitrollerMarketEntered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerMarketEntered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerMarketEntered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerMarketEntered represents a MarketEntered event raised by the Unitroller contract.
type UnitrollerMarketEntered struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Unitroller *UnitrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*UnitrollerMarketEnteredIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &UnitrollerMarketEnteredIterator{contract: _Unitroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Unitroller *UnitrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *UnitrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerMarketEntered)
				if err := _Unitroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Unitroller *UnitrollerFilterer) ParseMarketEntered(log types.Log) (*UnitrollerMarketEntered, error) {
	event := new(UnitrollerMarketEntered)
	if err := _Unitroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Unitroller contract.
type UnitrollerMarketExitedIterator struct {
	Event *UnitrollerMarketExited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerMarketExited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerMarketExited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerMarketExited represents a MarketExited event raised by the Unitroller contract.
type UnitrollerMarketExited struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Unitroller *UnitrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*UnitrollerMarketExitedIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &UnitrollerMarketExitedIterator{contract: _Unitroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Unitroller *UnitrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *UnitrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerMarketExited)
				if err := _Unitroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Unitroller *UnitrollerFilterer) ParseMarketExited(log types.Log) (*UnitrollerMarketExited, error) {
	event := new(UnitrollerMarketExited)
	if err := _Unitroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Unitroller contract.
type UnitrollerMarketListedIterator struct {
	Event *UnitrollerMarketListed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerMarketListed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerMarketListed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerMarketListed represents a MarketListed event raised by the Unitroller contract.
type UnitrollerMarketListed struct {
	CToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Unitroller *UnitrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*UnitrollerMarketListedIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &UnitrollerMarketListedIterator{contract: _Unitroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Unitroller *UnitrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *UnitrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerMarketListed)
				if err := _Unitroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Unitroller *UnitrollerFilterer) ParseMarketListed(log types.Log) (*UnitrollerMarketListed, error) {
	event := new(UnitrollerMarketListed)
	if err := _Unitroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Unitroller contract.
type UnitrollerNewBorrowCapIterator struct {
	Event *UnitrollerNewBorrowCap // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewBorrowCap)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewBorrowCap)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewBorrowCap represents a NewBorrowCap event raised by the Unitroller contract.
type UnitrollerNewBorrowCap struct {
	CToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Unitroller *UnitrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, cToken []common.Address) (*UnitrollerNewBorrowCapIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewBorrowCapIterator{contract: _Unitroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Unitroller *UnitrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *UnitrollerNewBorrowCap, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewBorrowCap)
				if err := _Unitroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Unitroller *UnitrollerFilterer) ParseNewBorrowCap(log types.Log) (*UnitrollerNewBorrowCap, error) {
	event := new(UnitrollerNewBorrowCap)
	if err := _Unitroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Unitroller contract.
type UnitrollerNewBorrowCapGuardianIterator struct {
	Event *UnitrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewBorrowCapGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewBorrowCapGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Unitroller contract.
type UnitrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Unitroller *UnitrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*UnitrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewBorrowCapGuardianIterator{contract: _Unitroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Unitroller *UnitrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *UnitrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewBorrowCapGuardian)
				if err := _Unitroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Unitroller *UnitrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*UnitrollerNewBorrowCapGuardian, error) {
	event := new(UnitrollerNewBorrowCapGuardian)
	if err := _Unitroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Unitroller contract.
type UnitrollerNewCloseFactorIterator struct {
	Event *UnitrollerNewCloseFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewCloseFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewCloseFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewCloseFactor represents a NewCloseFactor event raised by the Unitroller contract.
type UnitrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Unitroller *UnitrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*UnitrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewCloseFactorIterator{contract: _Unitroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Unitroller *UnitrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *UnitrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewCloseFactor)
				if err := _Unitroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Unitroller *UnitrollerFilterer) ParseNewCloseFactor(log types.Log) (*UnitrollerNewCloseFactor, error) {
	event := new(UnitrollerNewCloseFactor)
	if err := _Unitroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Unitroller contract.
type UnitrollerNewCollateralFactorIterator struct {
	Event *UnitrollerNewCollateralFactor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewCollateralFactor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewCollateralFactor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Unitroller contract.
type UnitrollerNewCollateralFactor struct {
	CToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Unitroller *UnitrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*UnitrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewCollateralFactorIterator{contract: _Unitroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Unitroller *UnitrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *UnitrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewCollateralFactor)
				if err := _Unitroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Unitroller *UnitrollerFilterer) ParseNewCollateralFactor(log types.Log) (*UnitrollerNewCollateralFactor, error) {
	event := new(UnitrollerNewCollateralFactor)
	if err := _Unitroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Unitroller contract.
type UnitrollerNewLiquidationIncentiveIterator struct {
	Event *UnitrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewLiquidationIncentive)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewLiquidationIncentive)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Unitroller contract.
type UnitrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Unitroller *UnitrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*UnitrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewLiquidationIncentiveIterator{contract: _Unitroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Unitroller *UnitrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *UnitrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewLiquidationIncentive)
				if err := _Unitroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Unitroller *UnitrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*UnitrollerNewLiquidationIncentive, error) {
	event := new(UnitrollerNewLiquidationIncentive)
	if err := _Unitroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Unitroller contract.
type UnitrollerNewPauseGuardianIterator struct {
	Event *UnitrollerNewPauseGuardian // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewPauseGuardian)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewPauseGuardian)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Unitroller contract.
type UnitrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Unitroller *UnitrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*UnitrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewPauseGuardianIterator{contract: _Unitroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Unitroller *UnitrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *UnitrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewPauseGuardian)
				if err := _Unitroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Unitroller *UnitrollerFilterer) ParseNewPauseGuardian(log types.Log) (*UnitrollerNewPauseGuardian, error) {
	event := new(UnitrollerNewPauseGuardian)
	if err := _Unitroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnitrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Unitroller contract.
type UnitrollerNewPriceOracleIterator struct {
	Event *UnitrollerNewPriceOracle // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *UnitrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnitrollerNewPriceOracle)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(UnitrollerNewPriceOracle)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *UnitrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnitrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnitrollerNewPriceOracle represents a NewPriceOracle event raised by the Unitroller contract.
type UnitrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Unitroller *UnitrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*UnitrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Unitroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &UnitrollerNewPriceOracleIterator{contract: _Unitroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Unitroller *UnitrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *UnitrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Unitroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnitrollerNewPriceOracle)
				if err := _Unitroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Unitroller *UnitrollerFilterer) ParseNewPriceOracle(log types.Log) (*UnitrollerNewPriceOracle, error) {
	event := new(UnitrollerNewPriceOracle)
	if err := _Unitroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
