package main

import (
	"fmt"
	"time"
)

func main() {

	for j := 0; j < 6; j += 2 {
		fmt.Println("Incrementando o J que tem o valor", j)
		time.Sleep(time.Second)
	}

	/*for -- 
	UsoQuando: preciso de REPETIDOR SEM AUTERAÇÕES -- COM CONTROLE DAS REPETIÇÕES

		Conceito :
		for interador := 0; <condicao<interador < 6; <faça interador ++ icremenatr 1 a cada interacao >>>
		Montagem for:
	for>enquanto varValor := 0; varValor < 6; varValor ++


			Declarando o valor da var interadora no escopo do for
			Obs essa variavel declarada no for só fica visivel ai neste escopo
			não da para acessa-la fora do escopo do for

			// incrementando de 1 em 1 usar no encrementador : ++
			// incrementando de 2 em 2 usar no encrementador : += 2



	*/

}
