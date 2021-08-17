package transaction

import (
	"encoding/json"
	"finances/model"
	"finances/util"
	"fmt"
	"io/ioutil"
	"net/http"
)

var transactions = model.Transactions{}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	transactions = append(
		transactions,
		model.Transaction{
			Title:     "Salário",
			Amount:    1200.0,
			Type:      0,
			CreatedAt: util.StringToTime("2020-04-05T11:45:26"),
		},
	)

	w.Header().Set("Content-type", "application/json")

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var body, _ = ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &transactions)

	fmt.Println(transactions)
}
