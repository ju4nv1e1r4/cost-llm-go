package main

import (
	"fmt"
	"llmcosts/app"
)

func main()  {
	fineTuningData := "iasudhfiosidofgposdkg aosidgjoisfdjgosdjfg juansosidjfosidjg llm iajsofisd osidjfosidjfosid osidjfosidjfosi osidjfosidjfo sodifjsodifjosij sodfi osdij osdif jsodi josdif jsodij osidfjosidj osidjfoisjdofi jsodifj osidfj osidjfoisjdofi jsodifj soidfj osidfjosidjfosidjfo sijdfoisjd ofisjdoifjsoidjf osidjfosidjfosijdfoisjdo ijfosi djfosi jdfoi jsdofi jsodifj soidfj osidjf osidjf oisd"

	trainningCOst := app.CalcTrainning("gpt-4o-2024-08-06", fineTuningData)
	inputCost := app.CalcInput("gpt-4o-2024-08-06", "Qual a sua função?")

	fmt.Printf("O custo do treinamento deste modelo é USD %f.\n", trainningCOst)
	fmt.Printf("O custo do prompt deste modelo é USD %f.", inputCost)
}