package main

import (
	"github.com/labstack/echo/v4"
	bank "apigo/bank"
)

var accounts []bank.Account
var transfers []bank.Transfer

func  createBankAccount(c echo.Context) error {
	acc := new(bank.Account)
	if err := c.Bind(acc); err != nil{
		return err
	}

	idInserted, err := acc.SaveBankAccount()
	if err == nil{
		acc.Id = idInserted
	}
	
	return c.JSON(201, acc)
}


func createBankAccountsTransfer(c echo.Context) error{
	tr := new(bank.Transfer)
	if err := c.Bind(tr); err != nil{
		return err
	}

	//transfers = append(transfers, *tr)
	resp, err := tr.SaveBankTransfer()
	if err != nil{
		return c.JSON(500, err)
	}

	return c.JSON(200, resp)
}

func main(){
	e := echo.New()
	e.POST("/bank-accounts", createBankAccount)
	e.POST("/bank-accounts/transfer", createBankAccountsTransfer)
	e.Logger.Fatal(e.Start(":8000"))
}
