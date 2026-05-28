# Currency Converter

Um conversor de moedas em Go que realiza conversões de BRL (Real Brasileiro) para outras moedas utilizando taxas de câmbio pré-configuradas.

## O que faz

Este programa permite converter valores em Reais (BRL) para outras moedas através de um comando simples no terminal. Ele:

- Lê o valor em BRL a ser convertido
- Lê o código da moeda de destino (ex: USD, EUR, JPY)
- Carrega as taxas de câmbio do arquivo `currency.json`
- Realiza o cálculo de conversão
- Exibe o resultado com duas casas decimais

## Como compilar

```bash
go build -o convert main.go
```

Isso gera um executável chamado `convert`.

## Como usar

```bash
./convert [valor_em_brl] [moeda_destino]
```

### Exemplos

```bash
# Converter 10 BRL para USD
./convert 10 USD
# Resultado: 1.51

# Converter 12 BRL para JPY
./convert 12 JPY
# Resultado: 195.48

# Converter 100 BRL para EUR
./convert 100 EUR
# Resultado: 13.70
```

## Moedas suportadas

As moedas disponíveis estão definidas no arquivo `currency.json`. Atualmente suporta:

- USD (Dólar Americano)
- EUR (Euro)
- JPY (Iene Japonês)
- GBP (Libra Esterlina)
- CHF (Franco Suíço)
- AUD (Dólar Australiano)
- CAD (Dólar Canadense)
- CNY (Yuan Chinês)
- HKD (Dólar de Hong Kong)
- NZD (Dólar Neozelandês)
- E outras moedas disponíveis no arquivo

## Testes

### Teste 1: Conversão básica para USD
```bash
./convert 10 USD
# Esperado: 1.51
```

### Teste 2: Conversão para JPY
```bash
./convert 12 JPY
# Esperado: 195.48
```

### Teste 3: Sensibilidade a maiúsculas/minúsculas
```bash
./convert 10 usd
# Esperado: 1.51 (funciona com letras minúsculas)
```

### Teste 4: Valor maior
```bash
./convert 1000 EUR
# Esperado: 137.00
```

### Teste 5: Moeda inválida (deve retornar erro)
```bash
./convert 10 XYZ
# Esperado: Erro: moeda não encontrada
```

### Teste 6: Número de argumentos incorreto (deve retornar erro)
```bash
./convert 10
# Esperado: Erro: número incorreto de argumentos
```

### Teste 7: Valor inválido (deve retornar erro)
```bash
./convert abc USD
# Esperado: Erro: valor em BRL inválido
```

## Tratamento de erros

O programa valida e trata os seguintes erros:

- ❌ Número incorreto de argumentos
- ❌ Valor em BRL não é um número
- ❌ Código de moeda não existe nas taxas de câmbio
- ❌ Arquivo `currency.json` não encontrado
- ❌ JSON malformado

## Estrutura do projeto

```
currency-converter/
├── main.go           # Código principal da aplicação
├── currency.json     # Arquivo com as taxas de câmbio
├── go.mod           # Módulo Go
└── README.md        # Este arquivo
```

## Conceitos praticados

Este projeto pratica os seguintes conceitos em Go:

- Argumentos de linha de comando (`os.Args`)
- Leitura de arquivos (`os.ReadFile`)
- Processamento de JSON (`json.Unmarshal`)
- Estruturas de dados (`struct`)
- Maps para armazenar dados
- Conversão de tipos (`strconv.ParseFloat`)
- Manipulação de strings (`strings.ToUpper`)
- Lógica condicional
- Tratamento de erros
- Formatação de saída (`fmt.Printf`)
