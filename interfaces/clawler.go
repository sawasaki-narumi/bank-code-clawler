package interfaces

import "github.com/PuerkitoBio/goquery"

type ClawlDriver interface {
	scrapeBranches(doc *goquery.Document) *[]*Branch
	saveAsCsv(filename string, branches *[]*Branch) error
	Exec() error
}
