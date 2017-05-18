package utils

import (
	"bytes"
	"fmt"
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
	resultsHash = resultsHash[:1] + `"EMAIL"=>` + email + "," + resultsHash[1:]
	resultsHash = strings.Replace(resultsHash, "=>", ":", -1)
	fmt.Println(resultsHash)
	p.data = []byte(resultsHash)
}

// Post to endpoint
func (p BigFiveResultsPoster) Post(url string) bool {
	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(p.data))

	//defer resp.Body.Close()

	p.ResponseCode = resp.StatusCode
	body, _ := ioutil.ReadAll(resp.Body)

	p.Token = string(body)
	fmt.Println(p.Token)
	if resp.StatusCode != 201 {
		return false
	}
	return true
}
