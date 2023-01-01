package apiproject

import (
	"apiproject/constants"
	balance "apiproject/usecase/balance"
	"apiproject/usecase/topup"
	"encoding/json"
	"fmt"
	"net/http"
)

func getBalance(balanceService *balance.Balance) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		amount, err := balanceService.CheckBalance()
		if err != nil {
			resp, _ := json.Marshal(balance.Response{
				Status: false,
				Errors: []constants.GenericError{{Code: "900", Message: err.Error()}},
			})
			w.Write(resp)
		}
		resp, _ := json.Marshal(balance.Response{
			Status: true,
			Data:   balance.Data{Amount: amount},
			Errors: nil,
		})
		w.Write(resp)
	}
}
func topUp(topupService *topup.Topup) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var topupRequest topup.Request

		err := json.NewDecoder(r.Body).Decode(&topupRequest)
		if err != nil {
			fmt.Errorf("topup error:%v", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = topupService.DoTopup(topupRequest)
		if err != nil {
			resp, _ := json.Marshal(topup.Response{
				Status: false,
				Errors: []constants.GenericError{{Code: "900", Message: err.Error()}},
			})
			w.Write(resp)
			return
		}
		resp, _ := json.Marshal(topup.Response{
			Status: true,
			Errors: nil,
		})
		w.Write(resp)
	}
}

func routes() {
	balanceService := balance.NewBalance()
	topupService := topup.NewTopup()
	http.HandleFunc("/project/v1/balance", getBalance(balanceService))
	http.HandleFunc("/project/v1/topup", topUp(topupService))
	http.HandleFucn("/project/v1/topup", topUp(topupSevice))
}
