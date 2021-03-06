package main

import (
	"fmt"
)

func main() {
	x1 := 10
	x2 := 20.5
	x3 := "Rei"
	x4 := "R"
	x5 := true

	fmt.Println("=============== TIPOS ===============")

	fmt.Printf("NUMERO    -   %T  valor: %v \n", x1, x1)
	fmt.Printf("DECIMAL   -   %T  valor: %v \n", x2, x2)
	fmt.Printf("TEXTO     -   %T  valor: %v \n", x3, x3)
	fmt.Printf("CARACTERE -   %T  valor: %v \n", x4, x4)
	fmt.Printf("LOGICO    -   %T  valor: %v \n", x5, x5)

	fmt.Println("========== NUMEROS INTEIROS POR BYTES ===============")

	// # alias apelidos para
	// int32 >> alias >> rune
	var numeroInt32 rune = -12456
	fmt.Println(numeroInt32)

	// uint8 >> alias >> byte
	var numeroUint8 byte = 123
	fmt.Println(numeroUint8)

	// ===== RESUMO : ================================================
	//OU VC DECLARA O TIPO ::: tipoVisibilidade nomevar tipo = "valor"
	var nomeVar1 string = "Reinaldo"
	fmt.Println(nomeVar1)

	//OU FAZ POR INFERENCIA O GO INTERFERE CONFORME O VALOR ATRIBUIDO
	nomeVar2 := 43
	fmt.Println(nomeVar2)

}

/*
	TIPOS NUMERICOS INTEIROS
	-- 4 tipos de tipos inteiros separados  por quantidade de bytes que ocupam

	#INT
	--- int  == aceita numeros negativos / int sozinho ele usa de acordo com sua arquitetura , o GO infere de acordo com sua maquina
	int8 >> ate o numero 123 >> alias >>
	int16
	int32 >> ate o numero 123 >> alias >> rune
	int64

	# UINT
	--- uint == sem sinal -- nao funciona para numeros com sinal antes do numero exemplo numeros negativos -1 , -1000 --usa-se para evitar retorno negativo
	--- uint8, uint16, uint32, uint64

	# DECIMAIS
	- FLOAT -- vc nao pode declarar só float tem que espeficar float32 ou float64 , mas para deixar por inferencia declarar com operador curto :=


*/

// :=  Operador que Declara e Atribui só funciona dentro do escopo/ blockcode
// Printf -- mostra a mensagem formatada
// Coloca as frases dentro de aspas duplas
// %v  para saber valor , SEPARAR ASPAS DA VAR COM VIRGULA e coloca a variavel depois das aspas
// %T para mostrar o tipo e a var que quer saber o tipo depois das aspas tbm
// iMPORTANTE ::  SEPARAR ASPAS DA MENSAGEM COM AS VARIAVEIS UTILIZADAS COM VIRGULA
// \n -- para pular linha.

/*
# OPERADOR CURTO
- := parece uma marmota (gopher) ou o punisher.
- Uso:
    - Tipagem automática
    - Só pode repetir se houverem variáveis novas
    - != do assignment operator (operador de atribuição)
    - Só funciona dentro de codeblocks
- Terminologia:
    - keywords (palavras-chave) são termos reservados
    - operadores, operandos
    - statement (declaração, afirmação) → uma linha de código, uma instrução que forma uma ação, formada de expressões
    - expressão → qualquer coisa que "produz um resultado"
    - scope (abrangência)
        - package-level scope
- Lição principal:
    - := utilizado pra criar novas variáveis, dentro de code blocks
		- = para atribuir valores a variáveis já existentes
======================================
- O tipo pode ser deduzido pelo compilador:
    - x := 10
    - var y = "a tia do batima"
- Ou pode ser declarado especificamente:
    - var w string = "isso é uma string"
    - var z int = 15
    - na declaração var z int com package scope, atribuição z = 15 no codeblock (somente)
- Tipos de dados primitivos: disponíveis na linguagem nativamente como blocos básicos de construção
    - int, string, bool
- Tipos de dados compostos: são tipos compostos de tipos primitivos, e criados pelo usuário
    - slice, array, struct, map
- O ato de definir, criar, estruturar tipos compostos chama-se composição.



*/
