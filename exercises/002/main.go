package main

import ("fmt")


func multiplicar(numeros ...int) int{ // Em python não seria diferente eu só trocaria o "...int" por *args
	if len(numeros) == 0{
		return 0
	}

	total := 1

	for _, numero := range(numeros){
		total *= numero
	}
	return total
}

func parOuImpar(numero int) string {
    if numero%2 == 0 {
        return "Par"
    }
    return "Impar" 
}

func main(){
	resultado := multiplicar(2, 4, 5)
    fmt.Println("O total da multiplicação é:", resultado)
	conclusao := parOuImpar(resultado)
	fmt.Println("O resultado obtido da operação é: ", conclusao)
}