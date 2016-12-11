package main

import (
	"fmt"
	"os"
)

const ERR_NO_SUBCOMMAND = "（aknの使い方を説明する）"
const ERR_INVALID_SUBCOMMAND = "（aknの使い方を説明する）"

func main() {
	if len(os.Args) < 2 {
		fmt.Println(ERR_NO_SUBCOMMAND)
		return;
	}
	switch(os.Args[1]){
	case "test" :

		//engine.test(os.Args[2])
	default :
		fmt.Println(ERR_INVALID_SUBCOMMAND)
	}
}