package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sawasaki-narumi/bank-code-clawler/driver/yamaguchi"
)

func displayResponse(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	//	fukuoka := fukuoka.Open("csv/fukuoka.csv")
	yamaguchi, err := yamaguchi.Open("csv/yamaguchi.csv")
	if err != nil {
		log.Fatal(err)
	}

	if err = yamaguchi.Exec(); err != nil {
		log.Fatal(err)
	}
}

// string to uint16
// numberStr = strings.TrimSpace(numberStr)
// number, err := strconv.Atoi(numberStr)
// uint16(number)
