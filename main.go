package main

import (
	"fmt"
	"github.com/sidletsky/knapsack/repository"
	"math"
	"strconv"
	"strings"
	"time"
)

type Knapsack struct {
	TotalValue           int
	TotalWeight          int
	Items                []repository.Item
	CharacteristicVector string
}

func main() {
	repo, err := repository.New("data/7")
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
	itemsSize := len(items)
	totalKnapsacks := int(math.Pow(2, float64(itemsSize)))
	perfectKnapsack = Knapsack{}
BinaryIterator:
	for i := 0; i < totalKnapsacks; i++ {
		// create vector representation
		base2 := strconv.FormatInt(int64(i), 2)
		vector := strings.Repeat("0", itemsSize-len(base2)) + base2
		// building knapsack with items
		currentKnapsack := Knapsack{}
		for i, v := range vector {
			if currentKnapsack.TotalWeight > knapsackCapacity {
				continue BinaryIterator
			}
			// 49 represents 1, 48 represents 0
			if v == 49 {
				currentKnapsack.TotalValue += items[i].Value
				currentKnapsack.TotalWeight += items[i].Weight
				currentKnapsack.Items = append(currentKnapsack.Items, items[i])
				currentKnapsack.CharacteristicVector = vector
			}
		}
		// choosing max knapsack on the fly
		if knapsackCapacity > currentKnapsack.TotalWeight && perfectKnapsack.TotalValue < currentKnapsack.TotalValue {
			perfectKnapsack = currentKnapsack
		}
	}
	return
}
