//qiang.ji@yahoo.com.au YRzqVEdSnDURzboSDswAXonq
package main

import (
	"fmt"

	"github.com/jiqiang/big-five-profile-poster/utils"
)

func main() {
	config := utils.GetConfig("./config")

	contentBytes := utils.GetFileContent(config.Source)

	serializer := utils.BigFiveResultsTextSerializer{}
	serializer.Read(string(contentBytes))

	result := serializer.Hash()

	poster := utils.BigFiveResultsPoster{}
	poster.Initialize(result, config.Email)

	poster.Post(config.Endpoint)
	fmt.Println(poster.ResponseCode)
	fmt.Println(poster.Token)
}
