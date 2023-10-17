package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main(){

	// balance
	var bankBalance int
	
	var balance sync.Mutex

	fmt.Printf("Initial account bal: %d", bankBalance)
	fmt.Println()

	incomes := []Income{
		{
			Source: "Main Job",
			Amount: 100,
		},
		{
			Source: "Part Time Job",
			Amount: 200,
		},
		{
			Source: "another source 1",
			Amount: 10,
		},
		{
			Source: "Investments",
			Amount: 20,
		},
	}

	wg.Add(len(incomes))

	for i, income := range incomes{
		go func(i int , income Income){
			defer wg.Done()
			for week := 1;  week<=52 ; week++{
				balance.Lock()
				temp := bankBalance
				temp+=income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("Amount Earned from %s in week %d: %d\n", income.Source, week, income.Amount)
			}

		}(i, income)
	}
	wg.Wait()

	fmt.Printf("Final account bal: %d", bankBalance)
	
}