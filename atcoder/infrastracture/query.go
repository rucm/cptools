package atcoder

import (
	"log"
	"strings"

	"github.com/rucm/cptools/config"

	"github.com/PuerkitoBio/goquery"

	"github.com/rucm/cptools/domain"
)

// Query :
type Query struct {
	Session       *domain.Session
	primitiveData *goquery.Document
}

// Download :
func (query *Query) Download(contestID string) {

	URL := strings.Replace(config.Config.AtCoder.ContestBaseURL, "<contestID>", contestID, 1)
	req := &RequestHandler{
		URL:         URL,
		Method:      "POST",
		ContentType: "application/x-www-form-urlencoded",
		Session:     query.Session,
	}

	res := req.Execute(nil)
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res.Body))

	if err != nil {
		log.Fatal(err)
	}

	query.primitiveData = doc
}

// Parse :
func (query *Query) Parse() {

}

// Get :
func (query *Query) Get() *domain.Contest {

	return nil
}
