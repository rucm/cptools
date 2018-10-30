package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Session struct
type Session struct {
	kickID    string
	userID    string
	username  string
	issueTime string
	session   string
	privilege string
}

// Set session data
func (s *Session) Set(cookies []*http.Cookie) *Session {
	for _, c := range cookies {
		switch c.Name {
		case "_kick_id":
			s.kickID = c.Value
		case "_user_id":
			s.userID = c.Value
		case "_user_name":
			s.username = c.Value
		case "_issue_time":
			s.issueTime = c.Value
		case "_session":
			s.session = c.Value
		}
	}

	return s
}

func writeGob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}

func readGob(filePath string, object interface{}) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(object)
	}
	file.Close()
	return err
}

func main() {
	values := url.Values{}
	values.Add("name", "")
	values.Add("password", "")

	req, err := http.NewRequest(
		"POST",
		"https://practice.contest.atcoder.jp/login",
		strings.NewReader(values.Encode()),
	)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	cookies := res.Cookies()
	for _, c := range cookies {
		log.Println(c.Name, c.Value)
	}

	session := &Session{}
	session.Set(cookies)

	err = writeGob("./save.gob", session)
	if err != nil {
		log.Fatal(err)
	}

	session2 := &Session{}
	err = readGob("./save.gob", session2)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("_kick_id: %s", session2.kickID)
	log.Printf("_user_id: %s", session2.userID)
	log.Printf("_user_name: %s", session2.username)
	log.Printf("_issue_time: %s", session2.issueTime)
	log.Printf("_session: %s", session2.session)
	log.Printf("_privilege: %s", session2.privilege)

}
