package main

import (
	"fmt"
	"github.com/ju4nv1e1r4/cost-llm-go/cost"
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

    trainningCost :=  cost.CalcTrainning(text, string(fineTuningData))
    fmt.Printf("=> O custo do treinamento deste modelo é USD %f.\n", trainningCost)
}