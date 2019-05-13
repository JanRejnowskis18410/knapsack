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
	CharacteristicVector []bool
}

func main() {
	repo, err := repository.New("data/test3")
	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	perfectKnapsack := getPerfectKnapsack(repo.Items, repo.Capacity)
	duration := time.Since(startTime)
	fmt.Println("Time took:", duration)
	fmt.Println("Knapsack's capacity:", repo.Capacity)
	fmt.Println("Characteristic vector:", perfectKnapsack.CharacteristicVector)
	fmt.Println("Total value:", perfectKnapsack.TotalValue)
	fmt.Println("Total weight:", perfectKnapsack.TotalWeight)
	fmt.Println("Items:")
	for _, item := range perfectKnapsack.Items {
		fmt.Println("id:", item.Id, "weight:", item.Weight, "value:", item.Value)
	}
}

// getPerfectKnapsack finds a perfect knapsack from all possible solutions using brute force method.
// Perfect knapsack is the one that has the biggest total value of items and has weight
// lower or equal to the capacity of a required knapsack.
// getPerfectKnapsack finds perfect knapsack on the fly, without producing any helper matrices
func getPerfectKnapsack(items []repository.Item, knapsackCapacity int) (perfectKnapsack Knapsack) {
	totalKnapsacks := math.Pow(2, float64(len(items)))
	perfectKnapsack = Knapsack{}
	for i := 0.; i < totalKnapsacks; i++ {
		// create vector representation
		vector := make([]bool, len(items))
		temp := i
		for j := len(items) - 1; j >= 0; j-- {
			vector[j] = math.Mod(temp, 2) == 1
			temp = math.Floor(temp / 2)
		}
		// building knapsack with items
		currentKnapsack := Knapsack{}
		for i, v := range vector {
			if v {
				currentKnapsack.TotalValue += items[i].Value
				currentKnapsack.TotalWeight += items[i].Weight
				currentKnapsack.Items = append(currentKnapsack.Items, items[i])
				currentKnapsack.CharacteristicVector = vector
			}
		}
		// choosing max knapsack on the fly
		if perfectKnapsack.TotalValue < currentKnapsack.TotalValue && currentKnapsack.TotalWeight <= knapsackCapacity {
			perfectKnapsack = currentKnapsack
		}
	}
	return
}
