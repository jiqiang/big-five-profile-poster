package main

import (
	"fmt"

	"github.com/jiqiang/big-five-profile-poster/utils"
)

func main() {
	config := utils.GetConfig("./config")

	contentBytes := utils.GetFileContent(config.Source)

	fmt.Println(string(contentBytes))
}
