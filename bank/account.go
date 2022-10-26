package bank

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Account struct{
	Id int64 `json:"id"`
	Account_number string `json:"account_number"`
}

func (acc Account) SaveBankAccount() (int64, error) {
	db, err := sql.Open("sqlite3","database.sqlite")
	if err != nil {
		return 0, err
	}
	
	
	stmt, err := db.Prepare("INSERT INTO accounts ('account_number') VALUES ($1)")
	if err != nil{
		return 0, err
	}

	res, err := stmt.Exec(acc.Account_number)
	if err != nil{
		return 0, err
	}
	
	id, err := res.LastInsertId()

	if err != nil{
		return 0, err
	}

	defer db.Close()

	return id, nil
}