package utils

import (
	"encoding/json"
	"os"

	"github.com/sei-protocol/golang-sdk/types"
)

func ParseTestFile(filename string) types.Test {
	pwd, _ := os.Getwd()
	file, _ := os.ReadFile(pwd + "/tests/" + filename + ".json")
	test := types.Test{}
	if err := json.Unmarshal([]byte(file), &test); err != nil {
		panic(err)
	}
	return test
}

func ParseMultiCollateralAccounts(raw json.RawMessage) types.MultiCollateralAccounts {
	accounts := types.MultiCollateralAccounts{}
	if err := json.Unmarshal(raw, &accounts); err != nil {
		panic(err)
	}
	return accounts
}

func ParseFundedOrder(raw json.RawMessage) types.FundedOrder {
	order := types.FundedOrder{}
	if err := json.Unmarshal(raw, &order); err != nil {
		panic(err)
	}
	return order
}

func ParseCancelOrder(raw json.RawMessage) types.CancelOrder {
	cancelOrder := types.CancelOrder{}
	if err := json.Unmarshal(raw, &cancelOrder); err != nil {
		panic(err)
	}
	return cancelOrder
}

func ParseDeposit(raw json.RawMessage) types.Deposit {
	deposit := types.Deposit{}
	if err := json.Unmarshal(raw, &deposit); err != nil {
		panic(err)
	}
	return deposit
}

func ParseStartingBalance(raw json.RawMessage) types.StartingBalance {
	startingBalance := types.StartingBalance{}
	if err := json.Unmarshal(raw, &startingBalance); err != nil {
		panic(err)
	}
	return startingBalance
}

func ParseOracleUpdate(raw json.RawMessage) types.OracleUpdate {
	oracleUpdate := types.OracleUpdate{}
	if err := json.Unmarshal(raw, &oracleUpdate); err != nil {
		panic(err)
	}
	return oracleUpdate
}

func ParseLiquidation(raw json.RawMessage) types.LiquidationRequest {
	liquidationRequest := types.LiquidationRequest{}
	if err := json.Unmarshal(raw, &liquidationRequest); err != nil {
		panic(err)
	}
	return liquidationRequest
}

func ParseSleep(raw json.RawMessage) types.Sleep {
	sleep := types.Sleep{}
	if err := json.Unmarshal(raw, &sleep); err != nil {
		panic(err)
	}
	return sleep
}

func ParseContractBalance(balanceBytes []byte) types.ContractBalance {
	var res types.ContractBalance
	if err := json.Unmarshal(balanceBytes, &res); err != nil {
		panic(err)
	}
	return res
}

func ParseContractPosition(positionBytes []byte) types.ContractPosition {
	var res types.ContractPosition
	if err := json.Unmarshal(positionBytes, &res); err != nil {
		panic(err)
	}
	return res
}

func ParseContractGetOrderResponse(orderBytes []byte) types.ContractGetOrderResponse {
	var res types.ContractGetOrderResponse
	if err := json.Unmarshal(orderBytes, &res); err != nil {
		panic(err)
	}
	return res
}

func ParsePortfolioSpecs(specsBytes []byte) types.PortfolioSpecs {
	var res types.PortfolioSpecs
	if err := json.Unmarshal(specsBytes, &res); err != nil {
		panic(err)
	}
	return res
}
