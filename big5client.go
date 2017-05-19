//qiang.ji@yahoo.com.au YRzqVEdSnDURzboSDswAXonq
package main

import (
	"fmt"

	"github.com/jiqiang/big-five-profile-poster/utils"
)

func main() {
	// load configuration from file
	config := utils.GetConfig("./config")

	contentBytes := utils.GetFileContent(config.Source)
	serializer := utils.BigFiveResultsTextSerializer{}
	// read content
	serializer.Initialize(string(contentBytes))
	// build hash
	hash := serializer.Hash()

	poster := utils.BigFiveResultsPoster{}

	// read hash and add email to hash
	poster.Initialize(hash, config.Email)
	// post to endpoint
	poster.Post(config.Endpoint)

	// print out status code and token
	fmt.Println(poster.ResponseCode)
	fmt.Println(poster.Token)
}
