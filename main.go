package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/sawasaki-narumi/bank-code-clawler/driver"
	. "github.com/sawasaki-narumi/bank-code-clawler/structs"
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
	fuk := driver.Open()
	//	prefixes := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん"
	prefixes := "あ"

	branches := make([]*Branch, 0, 30)
	for _, prefix := range prefixes {
		prefix := url.QueryEscape(string(prefix))
		url := "http://www.fukuokabank.co.jp/atmsearch/?prefix=" + prefix
		doc, err := fuk.FetchDoc(url)
		if err != nil {
			log.Fatal(err)
		}
		branches = append(branches, *fuk.ScrapeBranches(doc)...)
		time.Sleep(1500 * time.Millisecond)
	}

	if err := fuk.SaveAsCsv("fukuoka.csv", &branches); err != nil {
		log.Fatal(err)
	}
}

// string to uint16
// numberStr = strings.TrimSpace(numberStr)
// number, err := strconv.Atoi(numberStr)
// uint16(number)
