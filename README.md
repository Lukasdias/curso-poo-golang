# Sistema de Gerenciamento de Contas Correntes

<p align="center">
  <img src="https://raw.githubusercontent.com/golang-samples/gopher-vector/master/gopher.png" alt="Gopher" width="200" height="200">
</p>

Bem-vindo à documentação do Sistema de Gerenciamento de Contas Correntes em Go! Nesta documentação, você encontrará informações e exemplos sobre o uso de structs e Programação Orientada a Objetos (POO) em Go para implementar um sistema de gerenciamento de contas correntes.

## Introdução

O Sistema de Gerenciamento de Contas Correntes é um aplicativo simples que permite criar e manipular contas correntes. Cada conta corrente é representada por uma struct chamada `CheckingAccount`, que possui os seguintes campos:

- `Holder` (clients.Holder): O titular da conta corrente.
- `NumAgency` (int): O número da agência da conta corrente.
- `NumAccount` (int): O número da conta corrente.
- `Funds` (float64): O saldo disponível na conta corrente.

A struct `Holder` é usada para representar as informações do titular da conta, incluindo `Name` (nome), `Cpf` (CPF) e `Profission` (profissão).

Nesta documentação, vamos explorar como criar, manipular e gerenciar contas correntes usando essas structs e os métodos implementados.

## Instalação

Antes de começar, certifique-se de ter o Go instalado em seu sistema. Você pode obter a versão mais recente do Go em [https://golang.org/dl/](https://golang.org/dl/).

Após a instalação do Go, você pode clonar o repositório do Sistema de Gerenciamento de Contas Correntes em seu diretório de trabalho usando o seguinte comando:

```bash
git clone https://github.com/Lukasdias/curso-poo-golang-alura
```

## Uso Básico

Para começar a usar o Sistema de Gerenciamento de Contas Correntes, siga as etapas abaixo:

1. Importe os pacotes necessários em seu código Go:

```go
import (
	"caminho/para/o/pacote/accounts"
	"caminho/para/o/pacote/clients"
)
```

2. Crie uma nova conta corrente usando a função `acc.NewCheckingAccount()`:

```go
holder := clients.Holder{
	Name:       "John Doe",
	Cpf:        "123.456.789-00",
	Profission: "Developer",
}

account := acc.NewCheckingAccount(holder, 589, 123456, 1000.0)
```

3. Acesse os campos da conta corrente e do titular:

```go
fmt.Println("Titular:", account.Holder.Name)
fmt.Println("CPF do Titular:", account.Holder.Cpf)
fmt.Println("Profissão do Titular:", account.Holder.Profission)
fmt.Println("Número da Agência:", account.NumAgency)
fmt.Println("Número da Conta:", account.NumAccount)
fmt.Println("Saldo:", account.Funds)
```

4. Realize operações com a conta corrente, como depósitos e saques:

```go
account.DepositFunds(500.0)
account.WithdrawFunds(200.0)
```

5. Obtenha informações atualizadas sobre a conta corrente:

```go
fmt.Println("Saldo Atualizado:", account.Funds)
```

6. Realize transferências de fundos entre contas correntes:

```go
destinyAccount := acc.NewCheckingAccount(clients.Holder{
	Name:       "Jane Smith",
	Cpf:        "987.654.321-00",
	Profission: "Designer",
}, 589, 654321, 0.0)

account.TransferFunds(200.0, destinyAccount)

fmt.Println("Saldo da Conta Origem:", account.Funds)
fmt.Println("Saldo da Conta Destino:", destinyAccount.Funds)
```

## Considerações Finais

O Sistema de Gerenciamento de Contas Correntes é uma introdução prática à programação orientada a objetos em Go, usando structs para representar contas correntes, titulares e operações comuns de gerenciamento. Esteja à vontade para explorar o código-fonte do projeto e expandir as funcionalidades de acordo com suas necessidades.

Divirta-se explorando a programação orientada a objetos em Go!

---
