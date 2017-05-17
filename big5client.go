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

	results := serializer.Hash()

	fmt.Println(results)
	fmt.Println(len(results))
}
