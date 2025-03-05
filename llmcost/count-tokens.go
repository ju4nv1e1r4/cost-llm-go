package llmcost

import (
	"fmt"
	"log"

	"github.com/pkoukk/tiktoken-go"
)

func Count(words string) int {
	text := words
	encoding := "cl100k_base"

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		log.Fatal(err)
	}

	token := tke.Encode(text, nil, nil)
	nOfTokens := len(token)
	fmt.Println("O número de tokens é: ", nOfTokens)
	return nOfTokens
}

func CalcTrainning(model string, input string) float64 {
	var fineTuningCost float64

	nOfTokens := Count(input)

	gpt35turbo := 8.0 / 1_000_000.0
	gpt4omini20240718 := 3 / 1_000_000.0
	gpt4o20240806 := 25 / 1_000_000.0
	
	if model == "gpt-3.5-turbo" {
		fineTuningCost = float64(nOfTokens) * float64(gpt35turbo)
	} else if model == "gpt-4o-mini-2024-07-18" {
		fineTuningCost = float64(nOfTokens) * float64(gpt4omini20240718)
	} else if model == "gpt-4o-2024-08-06" {
		fineTuningCost = float64(nOfTokens) * float64(gpt4o20240806)
	} else {
		fmt.Println("-> O modelo não está na lista de custos. <-")
	}

	return fineTuningCost
}

func CalcInput(model string, input string) float64 {
	var inputCost float64

	nOfTokens := Count(input)

	gpt35turbo := 3.0 / 1_000_000.0
	gpt4omini20240718 := 0.3 / 1_000_000.0
	gpt4o20240806 := 3.75 / 1_000_000.0
	
	if model == "gpt-3.5-turbo" {
		inputCost = float64(nOfTokens) * float64(gpt35turbo)
	} else if model == "gpt-4o-mini-2024-07-18" {
		inputCost = float64(nOfTokens) * float64(gpt4omini20240718)
	} else if model == "gpt-4o-2024-08-06" {
		inputCost = float64(nOfTokens) * float64(gpt4o20240806)
	} else {
		fmt.Println("-> O modelo não está na lista de custos. <-")
	}

	return inputCost
}