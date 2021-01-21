package fukuoka

import (
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	. "github.com/sawasaki-narumi/bank-code-clawler/structs"
	"github.com/sawasaki-narumi/bank-code-clawler/utils"

	"github.com/PuerkitoBio/goquery"
)

type ClawlDriver struct {
	outputPath string
}

const URL = "http://www.fukuokabank.co.jp/atmsearch/?prefix="

func Open(outputPath string) ClawlDriver {
	clawler := ClawlDriver{
		outputPath: outputPath,
	}
	return clawler
}

func (v ClawlDriver) saveAsCsv(filename string, branches *[]*Branch) error {
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

func (v ClawlDriver) scrapeBranches(doc *goquery.Document) *[]*Branch {
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

func (v ClawlDriver) Exec() error {
	prefixes := "あいうえおかきくけこさしすせそたちつてとなにぬねのはひふへほまみむめもやゆよらりるれろわをん"

	branches := make([]*Branch, 0, 30)
	for _, prefix := range prefixes {
		prefix := url.QueryEscape(string(prefix))
		url := URL + prefix
		doc, err := utils.FetchDoc(url)
		if err != nil {
			return err
		}
		branches = append(branches, *v.scrapeBranches(doc)...)
		time.Sleep(1500 * time.Millisecond)
	}

	if err := v.saveAsCsv(v.outputPath, &branches); err != nil {
		return err
	}
	return nil
}
