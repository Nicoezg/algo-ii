package primos

import "math/big"

const BASE = 20 // Para n tests, las chances de devolver true para un numero aleatorio no-n es (1/4)**n.
//Es recomendable y usual usar un n = 20; esto da una probabilidad de 0.000,000,000,001 de devolver un falso positivo.

func Siguiente(n int) int {
	//recibe por parametro un numero y devuelve el siguiente n una vez que se multiplica ese n por 3
	for !big.NewInt(int64(n)).ProbablyPrime(20) { //new int recibe el numero a checkear y probably prime recibe la cantidad de checkeos para que el % de certeza sea mayor
		n++
	}
	return n
}
