package main

import (
	"flag"
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
	CharacteristicVector string
}

func (knapsack *Knapsack) String() string {
	return fmt.Sprintf(
		"Characteristic vector: %s \n"+
			"Total value: %d \n"+
			"Total weight: %d",
		knapsack.CharacteristicVector,
		knapsack.TotalValue,
		knapsack.TotalWeight)
}

func main() {
	file := flag.String("file", "data/test3", "a string")
	fmt.Println("File:", *file)
	repo, err := repository.New(*file)
	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	perfectKnapsack := getPerfectKnapsack(repo.Items, repo.Capacity)
	duration := time.Since(startTime)
	fmt.Println("Time took:", duration)
	fmt.Println("Knapsack's capacity:", repo.Capacity)
	fmt.Println(perfectKnapsack.String())
	fmt.Println("Items:")
	for i, v := range perfectKnapsack.CharacteristicVector {
		// 49 represents 1, 48 represents 0
		if v == 49 {
			fmt.Println("id:", repo.Items[i].Id, "weight:", repo.Items[i].Weight, "value:", repo.Items[i].Value)
		}
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
		// build knapsack with items
		currentKnapsack := Knapsack{}
		for i, v := range vector {
			// 49 represents 1, 48 represents 0
			if v == 49 {
				currentKnapsack.TotalValue += items[i].Value
				currentKnapsack.TotalWeight += items[i].Weight
			}
			if currentKnapsack.TotalWeight > knapsackCapacity {
				continue BinaryIterator
			}
		}
		// choosing max knapsack on the fly
		if perfectKnapsack.TotalValue < currentKnapsack.TotalValue {
			perfectKnapsack = currentKnapsack
			perfectKnapsack.CharacteristicVector = vector
		}
	}
	return
}
