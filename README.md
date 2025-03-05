# CostLLM - Cost Estimation for AI Models

## Overview

CostLLM is a Go package designed to calculate the cost of fine-tuning and inferencing language models (LLMs) based on the number of tokens processed. This package is useful for **MLOps** and **AIOps** routines and recommended for professionals who want to estimate the operational costs of their models before execution.

With **CostLLM**, you can:
- Calculate the cost of training (fine-tuning) a model based on the number of input tokens.
- Estimate the cost of executing prompts for different models.
- Easily integrate the package into other Go projects.

## Installation

To use **CostLLM** in your project, add the package to your Go project using:

```sh
go get github.com/ju4nv1e1r4/cost-llm-go
```
If you have problems adding, try forcing go to add our latest update:

```sh
go clean -modcache 
go get github.com/ju4nv1e1r4/cost-llm-go
go mod tidy
```
And import it into your code:

```go
import "github.com/ju4nv1e1r4/cost-llm-go/cost"
```

## Main Functions

### `Count`

This function receives a text and returns the number of tokens generated with the `cl100k_base` encoding.

**Usage:**

```go
package main

import (
"fmt"
"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
text := "This is an example text."
tokens := cost.Count(text)
fmt.Println("Number of tokens:", tokens)
}
```
### `CalcTrainning`

Calculates the cost of training (fine-tuning) a model based on the number of tokens provided.

**Parameters:**
- `model`: Name of the AI ​​model (example: `gpt-3.5-turbo`).
- `input`: Text used for training.

**Usage:**

```go
package main

import (
"fmt"
"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
cost := cost.CalcTrainning("gpt-3.5-turbo", "Text for training.")
fmt.Println("Cost of Fine-Tuning:", cost)
}
```

### `CalcInput`

Calculates the inference cost (prompt usage) for a given model.

**Parameters:**
- `model`: Name of the AI ​​model.
- `input`: Text of the prompt sent for inference.

**Usage:**

```go
package main

import (
"fmt"
"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
cost := cost.CalcInput("gpt-4o-2024-08-06", "How can you help me today?")
fmt.Println("Inference Cost:", cost)
}
```
## Integration with Other Projects

**LLM Cost** can be easily integrated into any Go project that needs to estimate LLM costs. This is especially useful for **MLOps** and **AIOps** engineers, as it allows for better cost management when running models in production.

If you want to monitor expenses, you can add logs or store the estimated costs in a database for financial control of inference and training.

**Example of use on a web server with Gin:**

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
	r := gin.Default()

	r.POST("/cost/train", func(c *gin.Context) {
		var request struct {
			Model string `json:"model"`
			Input string `json:"input"`
		}
		c.BindJSON(&request)

		cost := cost.CalcTrainning(request.Model, request.Input)
		c.JSON(200, gin.H{"fine_tuning_cost": cost})
	})

	r.Run(":8080")
}
```

## Contribution

If you want to contribute to **Cost LLM**, feel free to open a pull request or suggest improvements via _issues_ in the repository.

## License

This project is licensed under **Apache 2.0**, ensuring that improvements and modifications remain available to the community.

## Contact

For questions or collaborations:

📧 Email: [juanvieir4@gmail.com](mailto:juanvieir4@gmail.com)

📌 GitHub: [ju4nv1e1r4](https://github.com/ju4nv1e1r4/)
📌 LinkedIn: [juanvieira85](https://www.linkedin.com/in/juanvieira85/)

---

*portuguese version:* ~>

# CostLLM - Estimativa de Custos para Modelos de IA

## Visão Geral

CostLLM é um pacote em Go projetado para calcular o custo de fine-tuning e inferência de modelos de linguagem (LLMs) baseando-se no número de tokens processados. Este pacote é útil para rotinas de **MLOps** e **AIOps**  e recomendado para profissionais que desejam estimar os custos operacionais de seus modelos antes da execução.

Com **CostLLM**, você pode:
- Calcular o custo de treinamento (fine-tuning) de um modelo com base no número de tokens de entrada.
- Estimar o custo de execução de prompts para diferentes modelos.
- Integrar facilmente o pacote em outros projetos Go.

## Instalação

Para utilizar **CostLLM** em seu projeto, adicione o pacote ao seu projeto Go usando:

```sh
go get github.com/ju4nv1e1r4/cost-llm-go
```

Caso tenha problemas ao adicionar, tente forçar o go a adicionar nossa última atualização:

```sh
go clean -modcache 
go get github.com/ju4nv1e1r4/cost-llm-go
go mod tidy
```

E importe-o no seu código:

```go
import "github.com/ju4nv1e1r4/cost-llm-go/cost"
```

## Funções Principais

### `Count`

Esta função recebe um texto e retorna o número de tokens gerados com a codificação `cl100k_base`.

**Uso:**

```go
package main

import (
	"fmt"
	"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
	text := "Este é um exemplo de texto."
	tokens := cost.Count(text)
	fmt.Println("Número de tokens:", tokens)
}
```

---

### `CalcTrainning`

Calcula o custo de treinamento (fine-tuning) de um modelo com base no número de tokens fornecido.

**Parâmetros:**
- `model`: Nome do modelo de IA (exemplo: `gpt-3.5-turbo`).
- `input`: Texto usado para treinamento.

**Uso:**

```go
package main

import (
	"fmt"
	"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
	cost := cost.CalcTrainning("gpt-3.5-turbo", "Texto para treinamento.")
	fmt.Println("Custo de Fine-Tuning:", cost)
}
```

---

### `CalcInput`

Calcula o custo de inferência (uso de prompts) para um modelo específico.

**Parâmetros:**
- `model`: Nome do modelo de IA.
- `input`: Texto do prompt enviado para inferência.

**Uso:**

```go
package main

import (
	"fmt"
	"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
	cost := cost.CalcInput("gpt-4o-2024-08-06", "Como você pode me ajudar hoje?")
	fmt.Println("Custo da Inferência:", cost)
}
```

## Integração com Outros Projetos

**LLM Cost** pode ser integrado facilmente em qualquer projeto Go que precise estimar custos de LLMs. Isso é especialmente útil para engenheiros de **MLOps** e **AIOps**, pois permite um melhor gerenciamento de custos na execução de modelos em produção.

Se você deseja monitorar gastos, você pode adicionar logs ou armazenar os custos estimados em um banco de dados para controle financeiro de inferência e treinamento.

**Exemplo de uso em um servidor web com Gin:**

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ju4nv1e1r4/cost-llm-go/cost"
)

func main() {
	r := gin.Default()

	r.POST("/cost/train", func(c *gin.Context) {
		var request struct {
			Model string `json:"model"`
			Input string `json:"input"`
		}
		c.BindJSON(&request)

		cost := cost.CalcTrainning(request.Model, request.Input)
		c.JSON(200, gin.H{"fine_tuning_cost": cost})
	})

	r.Run(":8080")
}
```

## Contribuição

Se você deseja contribuir com **Cost LLM**, fique à vontade para abrir uma _pull request_ ou sugerir melhorias via _issues_ no repositório.

## Licença

Este projeto é licenciado sob a **Apache 2.0**, garantindo que melhorias e modificações permaneçam disponíveis para a comunidade.

## Contato

Para dúvidas ou colaborações:

📧 Email: [juanvieir4@gmail.com](mailto:juanvieir4@gmail.com)

📌 GitHub: [ju4nv1e1r4](https://github.com/ju4nv1e1r4/)
📌 LinkedIn: [juanvieira85](https://www.linkedin.com/in/juanvieira85/)
