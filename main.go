package main

import (
	"fmt"
	"llmcosts/app"
	"os"
)

func main()  {
	fineTuningData, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	
	trainningCost := app.CalcTrainning("gpt-3.5-turbo", string(fineTuningData))
	inputCost := app.CalcInput("gpt-3.5-turbo", "Qual a sua função?")

	fmt.Printf("=> O custo do treinamento deste modelo é USD %f.\n", trainningCost)
	fmt.Printf("=> O custo do prompt deste modelo é USD %f.", inputCost)
}