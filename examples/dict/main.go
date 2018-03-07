package main

import (
	"fmt"
	"log"

	"github.com/ck-schmidi/goqute"
)

func main() {
	qb, err := goqute.NewQuteBrowserInstance()
	if err != nil {
		log.Fatal(err)
	}
	qb.ExecuteCommand([]byte(fmt.Sprintf("open -t http://www.dict.cc/?s=%s\n", qb.Params.SelectedText)))
}
