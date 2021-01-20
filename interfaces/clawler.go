package interfaces

import "github.com/PuerkitoBio/goquery"

type ClawlDriver interface {
	Hoge()
	FetchDoc(url string) (*goquery.Document, error)
	ScrapeBranches(doc *goquery.Document) *[]*Branch
	SaveAsCsv(filename string, branches *[]*Branch) error
}
