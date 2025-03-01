package app

import (
	"fmt"
	"log"

	"github.com/pkoukk/tiktoken-go"
)


func Count(words string)  {
	text := words
	encoding := "cl100k_base"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		log.Fatal(err)
	}

	token := tke.Encode(text, nil, nil)

	fmt.Println("Quantidade de tokens: ", len(token))
	fmt.Println("Tokens: ", token)
}