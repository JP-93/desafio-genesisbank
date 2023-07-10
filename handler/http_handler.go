package handler

import (
	"desafio-genesisbank/database"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Output struct {
	ValorConvertido float64 `json:"valorConvertido"`
	SimboloMoeda    string  `json:"SimboloMoeda"`
}

type Input struct {
	exchange float64
	Simbol   string
}

func ExchangeValue(w http.ResponseWriter, r *http.Request) {
	var data database.Output
	rs, err := filterParam(r)
	if err != nil {
		fmt.Println("Error on filter")
		w.WriteHeader(http.StatusInternalServerError)
	}
	data.ValorConvertido = rs.exchange
	data.SimboloMoeda = rs.Simbol
	database.DB.Create(&data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)

}

func GetAllValues(w http.ResponseWriter, r *http.Request) {
	var data []database.Output
	database.DB.Find(&data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// filterParm filtra os valores que estão no path da request
func filterParam(r *http.Request) (Input, error) {
	amout := chi.URLParam(r, "amout")
	from := chi.URLParam(r, "from")
	rate := chi.URLParam(r, "rate")
	to := chi.URLParam(r, "to")

	var s string
	var result float64

	if from == "USD" && to == "BRL" {
		result = exchangeUSD(amout, rate)
		s = "R$"
	} else if from == "EURO" && to == "BRL" {
		result = exchangeEuro(amout, rate)
		s = "R$"
	} else if from == "BTC" {
		resBTC, simbolBTC := exchangeBTC(amout, rate, to)
		result = resBTC
		s = simbolBTC
	} else {
		res, simbol := exchengeBr(amout, rate, to)
		result = res
		s = simbol
	}

	return Input{
		exchange: result,
		Simbol:   s,
	}, nil

}

// exchangeBr converte o real para dolar e euro
func exchengeBr(valueAmout, valueRate, valueTo string) (float64, string) {
	var res float64
	var s string
	if valueTo == "USD" {
		convAmout, err := strconv.ParseFloat(valueAmout, 8)
		if err != nil {
			panic("conversion error")
		}
		convRate, err := strconv.ParseFloat(valueRate, 8)
		if err != nil {
			panic("conversion error")
		}

		res = math.Round(convAmout / convRate)
		s = "$"
	} else if valueTo == "EURO" {
		convAmout, err := strconv.ParseFloat(valueAmout, 8)
		if err != nil {
			panic("conversion error")
		}
		convRate, err := strconv.ParseFloat(valueRate, 8)
		if err != nil {
			panic("conversion error")
		}
		res = math.Round(convAmout / convRate)
		s = "€"
	}
	return res, s
}

// exchangeUSD converte dolar para real
func exchangeUSD(valueAmout, valueRate string) float64 {
	convAmout, err := strconv.ParseFloat(valueAmout, 8)
	if err != nil {
		panic("conversion error")
	}
	convRate, err := strconv.ParseFloat(valueRate, 8)
	if err != nil {
		panic("conversion error")
	}
	res := convAmout * convRate
	return res
}

// exchangeEuro converte euro para real
func exchangeEuro(valueAmout, valueRate string) float64 {
	convAmout, err := strconv.ParseFloat(valueAmout, 8)
	if err != nil {
		panic("conversion error")
	}
	convRate, err := strconv.ParseFloat(valueRate, 8)
	if err != nil {
		panic("conversion error")
	}
	res := convAmout * convRate
	return res
}

// exchanteBTC converte bitcoins para real e dolar
func exchangeBTC(valueAmout, valueRate, valueTo string) (float64, string) {
	var res float64
	var s string

	if valueTo == "BRL" {
		convAmout, err := strconv.ParseFloat(valueAmout, 8)
		if err != nil {
			panic("conversion error")
		}
		convRate, err := strconv.ParseFloat(valueRate, 8)
		if err != nil {
			panic("conversion error")
		}
		res = convAmout * convRate
		s = "R$"
	} else if valueTo == "USD" {
		convAmout, err := strconv.ParseFloat(valueAmout, 8)
		if err != nil {
			panic("conversion error")
		}
		convRate, err := strconv.ParseFloat(valueRate, 8)
		if err != nil {
			panic("conversion error")
		}
		res = convAmout * convRate
		s = "$"
	}

	return res, s
}
