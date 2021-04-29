package main

import (
	"flag"
	"fmt"
	"math/big"
	"runtime"
	"time"
)

func main() {
	var n int
	var metodo string
	flag.IntVar(&n, "n", 2000000, "Numero Fibonacci a calcular")
	flag.StringVar(&metodo, "metodo", "matrix", "Método de cálculo 'matrix' o 'trasform'")
	flag.Parse()

	var result *big.Int
	start := time.Now()

	if metodo == "matrix" {
		result = fibmatrix(n)
	} else if metodo == "transform" {
		result = fibtransform(n)
	}

	elapsed := time.Since(start)
	fmt.Println(result)
	PrintMemUsage()
	fmt.Println("El cálculo tomó", elapsed, "con el método", metodo)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
