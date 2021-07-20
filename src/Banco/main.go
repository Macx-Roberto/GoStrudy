package main

import "fmt"

type ContaCorrente struct{
	ctitular string
 	nAgencia int
	nconta int // minusculo visibilidade em pacote maiusculo total
	nsaldo float64
}
//func (c *ContaCorrente) Sacar(valorSaque ...float64) string { "..." permite passsar um numero indeterminado de parametros
func (c *ContaCorrente) Sacar(valorSaque float64) string {
	if (valorSaque <= c.nsaldo){
		c.nsaldo -= valorSaque
		return "Saque concluido com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	contaSecundaria := ContaCorrente{}
	contaSecundaria.ctitular = "segundaconta"
	contaSecundaria.nAgencia = 1
	contaSecundaria.nconta   = 2 
	contaSecundaria.nsaldo   = 345.66
	fmt.Println(contaSecundaria)
	fmt.Println(ContaCorrente{ctitular: "terceiraconta", nAgencia: 1})
	fmt.Println(ContaCorrente{"Guilherme", 589, 589, 123.55})

	var contaDaCris *ContaCorrente // usa * ao utilizar o new 
	contaDaCris = new(ContaCorrente)
	contaDaCris.ctitular = "Cris"
	fmt.Println(contaDaCris)
	fmt.Println(*contaDaCris)

	contaA := ContaCorrente{"Guilherme", 589, 589, 1230.55}
	contaB := ContaCorrente{"Guilherme", 589, 589, 1230.55}
	println(contaA == contaB)
	
	contaB.ctitular = "zeze"
	println(contaA == contaB)
	// ao comparar em um if tbm seria necessario o * para utilização com new

	// contaB.nsaldo = contaB.nsaldo - nValorSaque
	fmt.Println(contaSecundaria.Sacar(500))
	fmt.Println(contaB.Sacar(500))
}