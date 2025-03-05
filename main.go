package main

import (
	"fmt"
	"llmcosts/llmcost"
	"os"
)

func main()  {
    fmt.Println("=======")
	fmt.Println("LLM Cost of Trainning and Inputs")
	fmt.Println("=======")

	fmt.Print("Qual o nome do seu modelo?\n")
	var text string
	fmt.Scanln(&text)
	fmt.Printf("Modelo => %s", text)
    fineTuningData, err := os.ReadFile("data.txt")
    if err != nil {
		panic(err)
	}

    trainningCost := llmcost.CalcTrainning(text, string(fineTuningData))
    fmt.Printf("=> O custo do treinamento deste modelo é USD %f.\n", trainningCost)
}