/*
		Copyright 2025 Juan Lucas Lopes Vieira

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

 		http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cost

import (
	"fmt"
	"log"

	"github.com/pkoukk/tiktoken-go"
)

// Count recebe um texto como entrada e retorna o número de tokens gerados usando a codificação "cl100k_base".
// Ele utiliza a biblioteca tiktoken para tokenização.
// Se ocorrer um erro ao obter a codificação, o programa encerrará com um log.Fatal.
func Count(words string) int {
	text := words
	encoding := "cl100k_base" // A maioria dos modelos usa este encoder

	tke, err := tiktoken.GetEncoding(encoding)
	if err != nil {
		log.Fatal(err)
	}

	token := tke.Encode(text, nil, nil)
	nOfTokens := len(token)
	
	fmt.Println("O número de tokens é: ", nOfTokens)
	return nOfTokens // devemos retornar um inteiro, pois servirá no cálculo em outras funções
}

// CalcTrainning estima o custo de fine-tuning de um modelo específico com base no número de tokens do input.
// Ele recebe o nome do modelo e o texto de entrada, calcula o número de tokens e aplica a taxa de fine-tuning correspondente.
// Se o modelo não estiver na lista de custos, uma mensagem de aviso é exibida.
// Retorna o custo formatado como uma string em dólares.
func CalcTrainning(model string, input string) string {
	var fineTuningCost float64

	nOfTokens := Count(input)

	// preços
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

	resultOf := fmt.Sprintf("USD: %f", fineTuningCost)
	return resultOf
}

// CalcInput estima o custo de processamento de entrada para um modelo específico com base no número de tokens do input.
// Ele recebe o nome do modelo e o texto de entrada, calcula os tokens e aplica a taxa correspondente para processamento de entrada.
// Se o modelo não estiver na lista de custos, uma mensagem de aviso é exibida.
// Retorna o custo formatado como uma string em dólares.
func CalcInput(model string, input string) string {
	var inputCost float64

	nOfTokens := Count(input)

	// preços
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

	resultOf := fmt.Sprintf("USD: %f", inputCost)
	return resultOf
}