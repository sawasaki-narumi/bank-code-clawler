package driver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	. "github.com/sawasaki-narumi/bank-code-clawler/structs"

	"github.com/PuerkitoBio/goquery"
)

type ClawlDriver struct {
	Prefecture string
	Url        string
	OutputPath string
}

func Open() ClawlDriver {
	clawler := ClawlDriver{
		Prefecture: "fukuoka",
		Url:        "fukuoka",
		OutputPath: "fukuoka",
	}
	return clawler
}

func (v ClawlDriver) FetchDoc(url string) (*goquery.Document, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func (v ClawlDriver) SaveAsCsv(filename string, branches *[]*Branch) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString("店番号,店名\n")
	if err != nil {
		return err
	}

	for _, branch := range *branches {
		row := branch.Number + "," + branch.Name + "\n"
		_, err := file.WriteString(row)
		if err != nil {
			return err
		}
	}

	return nil
}

func (v ClawlDriver) ScrapeBranches(doc *goquery.Document) *[]*Branch {
	branches := make([]*Branch, 0, 10)
	doc.Find("div.table_style02_block").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(_ int, s *goquery.Selection) {
			span := s.Find("td span").First()
			if span.Length() != 1 {
				return
			}
			branchNumber := strings.Replace(span.Text(), "店番：", "", -1)
			branchNumber = strings.TrimSpace(branchNumber)
			anchor := s.Find("strong a").First()
			branchName := anchor.Text()
			branch := Branch{
				Name:   branchName,
				Number: branchNumber,
			}
			fmt.Println(branch)
			branches = append(branches, &branch)
		})
	})
	return &branches
}
