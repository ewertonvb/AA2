package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var n_threads = 32
var parciais = make([]float64, n_threads)

func main() {
	serie := 1000000

	qtd := serie / n_threads
	var wg sync.WaitGroup
	wg.Add(n_threads)

	start := time.Now()
	for i := 0; i < n_threads; i++ {
		inicio := i * qtd
		fim := (i + 1) * qtd
		go calcPi(i, inicio, fim, &wg)
	}
	wg.Wait()

	var pi float64
	for _, parcial := range parciais {
		pi += parcial
	}

	fmt.Println("Valor de pi:", pi)
	// Code to measure
	print("\n")
	duration := time.Since(start)
	fmt.Println("Duração: ", duration.Milliseconds())
}

func calcPi(tid int, inicio int, fim int, wg *sync.WaitGroup) {
	var soma float64
	for i := inicio; i < fim; i++ {
		soma += (math.Pow(-1, float64(i)) * 4.) / (2.*float64(i) + 1.)
	}
	parciais[tid] = soma
	wg.Done()
}