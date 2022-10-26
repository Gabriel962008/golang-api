package bank

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


type Transfer struct{
	From   string  `json:"from"`
	To     string `json:"to"`
	Amount float64 `json:"amount"`
}

type TransferResponse struct{
	From_Account   string `json:"from_account"`
	From_Balance   string `json:"from_balance"`
	To_Account     string `json:"to_account"`
	To_Balance     string `json:"to_balance"`
}



func (tr Transfer) SaveBankTransfer() (TransferResponse, error){
	var trResponse TransferResponse
	db, err := sql.Open("sqlite3","database.sqlite")
	if err != nil {
		return trResponse, err
	}
	defer db.Close()
	
	stmt, err := db.Prepare("INSERT INTO transfer ('from', 'to', 'amount') VALUES ($1, $2, $3)")
	if err != nil{
		return trResponse, err
	}

	_, err = stmt.Exec(tr.From, tr.To, tr.Amount)
	if err != nil{
		return trResponse, err
	}

	trResponse.From_Account = tr.From
	err = db.QueryRow("SELECT (SUM(t.amount) - (SELECT SUM(t.amount) FROM transfer t  WHERE t.'from'= ?) ) AS BALANCE FROM transfer t  WHERE t.'to'= ?", tr.From, tr.From).Scan(&trResponse.From_Balance)
	if err != nil {
		return trResponse, err
	}

	trResponse.To_Account = tr.To
	err = db.QueryRow("SELECT (SUM(t.amount) - (SELECT SUM(t.amount) FROM transfer t  WHERE t.'from'= ?) ) AS BALANCE FROM transfer t  WHERE t.'to'= ?", tr.To, tr.To).Scan(&trResponse.To_Balance)
	if err != nil {
		return trResponse, err
	}

	return trResponse, nil
}