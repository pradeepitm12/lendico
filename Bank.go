package main

import (
	"net/http"
	"strconv"

	"github.com/pradeepitm12/lendico/utils"
)

type Input struct {
	rate     float64
	amount   float64
	duration float64
	date     string
}

func main() {
	http.HandleFunc("/annuity", BankLoan)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

/**
* Request Handler
 */
func BankLoan(w http.ResponseWriter, req *http.Request) {
	var input = Input{}
	query := req.URL.Query()
	input.date = query.Get("date")
	input.duration, _ = strconv.ParseFloat(query.Get("duration"), 64)
	input.amount, _ = strconv.ParseFloat(query.Get("amount"), 64)
	input.rate, _ = strconv.ParseFloat(query.Get("rate"), 64)
	list := utils.GenerateRepaymentPlan(input.rate, input.amount, input.duration, input.date)
	for _, elm := range list {

		w.Write([]byte("{ \n borrowerPaymentAmount : " + strconv.FormatFloat(elm.BorrowerPaymentAmount, 'f', 2, 64) + ",\n" +
			"date : " + elm.Date + ",\n" +
			"initialOutstandingPrincipal :  " + strconv.FormatFloat(elm.InitialOutstandingPrincipal, 'f', 2, 64) + ",\n" +
			"interest : " + strconv.FormatFloat(elm.Interest, 'f', 2, 64) + ",\n" +
			"principal : " + strconv.FormatFloat(elm.Principal, 'f', 2, 64) + ",\n" +
			"remainingOutstandingPrincipal : " + strconv.FormatFloat(elm.RemainingOutstandingPrincipal, 'f', 2, 64) + "\n}\n"))
	}

	//fmt.Println(utils.GenerateRepaymentPlan(input.rate, input.amount, input.duration, input.date))
}
