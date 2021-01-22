package yamaguchi

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	. "github.com/sawasaki-narumi/bank-code-clawler/structs"
	"github.com/sawasaki-narumi/bank-code-clawler/utils"
)

type ClawlDriver struct {
	outputPath string
}

func Open(outputPath string) (*ClawlDriver, error) {
	v := &ClawlDriver{outputPath: outputPath}
	return v, nil
}

func (v ClawlDriver) Exec() error {
	urls := []string{
		"https://www.yamaguchibank.co.jp/outline/store/shimonoseki.html",
	}

	for _, url := range urls {
		doc, err := utils.FetchDoc(url)
		if err != nil {
			log.Fatal(err)
		}

		v.scrapeBranches(doc)
	}
	return nil
}

func (v ClawlDriver) scrapeBranches(doc *goquery.Document) *[]*Branch {
	branches := make([]*Branch, 0, 10)
	doc.Find("div#pageTop div.l_content div.m_address_table table").Each(func(i int, s *goquery.Selection) {
		th := s.Find("tr th.table_tit").First()
		re := regexp.MustCompile(`（店番：\d+）`)

		var foundString, branchNumber string
		if foundString = re.FindString(th.Text()); len(foundString) == 0 {
			return
		}
		foundString = strings.Replace(foundString, "（店番：", "", -1)
		branchNumber = strings.Replace(foundString, "）", "", -1)

		branchName := re.ReplaceAllString(th.Text(), "")
		branch := Branch{
			Name:   branchName,
			Number: branchNumber,
		}
		fmt.Println(branch)
		branches = append(branches, &branch)
	})
	return nil
	return &branches
}
