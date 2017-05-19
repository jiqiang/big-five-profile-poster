package utils

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

// BigFiveResultsPoster class
type BigFiveResultsPoster struct {
	data         []byte
	ResponseCode int
	Token        string
}

// Initialize hash and email
func (p *BigFiveResultsPoster) Initialize(resultsHash string, email string) {
	// add email to the hash
	resultsHash = resultsHash[:1] + `"EMAIL"=>"` + email + "\"," + resultsHash[1:]
	// reformat hash to be a valid json string
	resultsHash = strings.Replace(resultsHash, "=>", ":", -1)
	// set data to be hash byts array for post use
	p.data = []byte(resultsHash)
}

// Post to endpoint
func (p BigFiveResultsPoster) Post(url string) bool {
	// post to endpoint and set content type to be json
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(p.data))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// get status code
	p.ResponseCode = resp.StatusCode
	// read body from response
	body, _ := ioutil.ReadAll(resp.Body)
	// set token from body
	p.Token = string(body)
	if resp.StatusCode != 201 {
		return false
	}
	return true
}
