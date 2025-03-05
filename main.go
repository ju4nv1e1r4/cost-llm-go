package main

import (
	"fmt"
	"llmcosts/app"
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

	fmt.Println("\n---------")
	fmt.Print("Digite um prompt:\n")
	var input string
	fmt.Scanln(&input)
	fmt.Printf("Prompt => %s\n", input)

	fineTuningData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	
	trainningCost := app.CalcTrainning(text, string(fineTuningData))
	inputCost := app.CalcInput(text, input)

	fmt.Printf("=> O custo do treinamento deste modelo é USD %f.\n", trainningCost)
	fmt.Printf("=> O custo do prompt deste modelo é USD %f.", inputCost)
}