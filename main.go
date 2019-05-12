package main

import (
	"fmt"
	"github.com/sidletsky/knapsack/repository"
	"math"
	"time"
)

type Knapsack struct {
	TotalValue           int
	TotalWeight          int
	Items                []repository.Item
	CharacteristicVector []int
}

func main() {
	repo, err := repository.New("data/test2")
	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	totalKnapsacks := math.Pow(2, float64(repo.Size))
	maxKnapsack := Knapsack{}
	for i := 0.; i < totalKnapsacks; i++ {
		// create vector representation
		vector := make([]int, repo.Size)
		temp := i
		for j := repo.Size - 1; j >= 0; j-- {
			vector[j] = int(math.Mod(temp, 2))
			temp = math.Floor(temp / 2)
		}
		// building knapsack with items
		currentKnapsack := Knapsack{}
		for i, v := range vector {
			if v == 1 {
				currentKnapsack.TotalValue += repo.Items[i].Value
				currentKnapsack.TotalWeight += repo.Items[i].Weight
				currentKnapsack.Items = append(currentKnapsack.Items, repo.Items[i])
				currentKnapsack.CharacteristicVector = vector
			}
		}
		// choosing max knapsack on the fly
		if maxKnapsack.TotalValue < currentKnapsack.TotalValue && currentKnapsack.TotalWeight <= repo.Capacity {
			maxKnapsack = currentKnapsack
		}
	}
	// perfect knapsack output
	fmt.Println("Knapsack's capacity:", repo.Capacity)
	fmt.Println("Characteristic vector:", maxKnapsack.CharacteristicVector)
	fmt.Println("Total value:", maxKnapsack.TotalValue)
	fmt.Println("Total weight:", maxKnapsack.TotalWeight)
	fmt.Println("Items:")
	for _, item := range maxKnapsack.Items {
		fmt.Println("id:", item.Id, "weight:", item.Weight, "value:", item.Value)
	}
	duration := time.Since(startTime)
	fmt.Println("Time took:", duration)
}
