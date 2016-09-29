package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"strconv"
)

func FloatToString(input_num float64) string {
    return strconv.FormatFloat(input_num, 'f', 2, 64)
}

func main() {

	//базовая валюта в журнале
	basecurrency := "RUB"
	//ее название в журнале
	basecurrencyname := "руб"
	//имя создаваемого файла с ценами
	pricedbfilename := "prices.dat"
	//коды валют для загрузки
	currcode := "USD,EUR"
	//адрес api сайта
	url := "https://api.fixer.io/latest?base=%s&symbols=%s"
	//время для записи в файл, указано стандартное. todo: установить время загрузки цен
	defaulttime := "10:00:00"

	resp, err := http.Get(fmt.Sprintf(url, basecurrency, currcode))
	if err != nil {
		fmt.Println("http get error")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading body")
	}

	var apiObj map[string]*json.RawMessage
	
	var rates map[string]float64
	
	var currdate string

	err = json.Unmarshal(body, &apiObj)

	err = json.Unmarshal(*apiObj["date"], &currdate)

	currdate = strings.Replace(currdate, "-", "/", -1)

	err = json.Unmarshal(*apiObj["rates"], &rates)
	
	p := ""
	//для каждой валюты создаем строку
	for code, rate := range rates {
		p += fmt.Sprintf("P %v %v %v %v %v\n", currdate, defaulttime, code, FloatToString(1/rate), basecurrencyname)
	}
	//выведем в консоль
	fmt.Println(p)
	//запишем в файл
    err = ioutil.WriteFile(pricedbfilename, []byte(p), 0644)
    if err != nil {
        panic(err)
    }

	return
}