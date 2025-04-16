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

	fmt.Print("What model are you using?\n")
	var text string
	scanText, err := fmt.Scanln(&text)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Model => %s", text)
	fmt.Println("=======")
	_ = scanText
    fineTuningData, err := os.ReadFile("data.txt") // your file path here
    if err != nil {
		panic(err)
	}

    cost.CalcTrainning(text, string(fineTuningData))
	cost.CalcInput(text, string(fineTuningData))
}