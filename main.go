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
	passport  string
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
		case "__passport":
			s.passport = c.Value
		case "__privilege":
			s.privilege = c.Value
		}
	}

	return s
}

func (s *Session) Write(filePath string) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(&s.kickID)
		encoder.Encode(&s.userID)
		encoder.Encode(&s.username)
		encoder.Encode(&s.issueTime)
		encoder.Encode(&s.session)
		encoder.Encode(&s.passport)
		encoder.Encode(&s.privilege)
	}
	file.Close()
	return err
}

func (s *Session) Read(filePath string) error {
	file, err := os.Open(filePath)
	if err == nil {
		decoder := gob.NewDecoder(file)
		decoder.Decode(&s.kickID)
		decoder.Decode(&s.userID)
		decoder.Decode(&s.username)
		decoder.Decode(&s.issueTime)
		decoder.Decode(&s.session)
		decoder.Decode(&s.passport)
		decoder.Decode(&s.privilege)
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

	err = session.Write("save.gob")
	if err != nil {
		log.Fatal(err)
	}

	session2 := &Session{}
	err = session2.Read("save.gob")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("_kick_id: %s", session2.kickID)
	log.Printf("_user_id: %s", session2.userID)
	log.Printf("_user_name: %s", session2.username)
	log.Printf("_issue_time: %s", session2.issueTime)
	log.Printf("_session: %s", session2.session)
	log.Printf("__passport: %s", session2.passport)
	log.Printf("__privilege: %s", session2.privilege)

}
