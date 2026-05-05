package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ExchangeData struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Erro: número incorreto de argumentos")
		fmt.Println("Uso: go run . [valor_em_brl] [moeda_destino]")
		fmt.Println("Exemplo: go run . 10 USD")
		os.Exit(1)
	}

	valueStr := os.Args[1]
	targetCurrency := strings.ToUpper(os.Args[2])

	data, err := os.ReadFile("currency.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo currency.json:", err)
		os.Exit(1)
	}

	var exchangeData ExchangeData
	err = json.Unmarshal(data, &exchangeData)
	if err != nil {
		fmt.Println("Erro ao processar JSON:", err)
		os.Exit(1)
	}

	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		fmt.Println("Erro: valor em BRL inválido")
		os.Exit(1)
	}

	rate, exists := exchangeData.Rates[targetCurrency]
	if !exists {
		fmt.Println("Erro: moeda não encontrada")
		os.Exit(1)
	}

	result := value * rate

	fmt.Printf("%.2f\n", result)
}
