package main

import (
	"fmt"
	repo "github.com/sidletsky/knapsack/repository"
	"math"
	"sort"
	"time"
)

type Knapsack struct {
	TotalValue           int
	TotalWeight          int
	Items                []repo.Item
	CharacteristicVector []int
}

func main() {
	repository, err := repo.New("data/test1")
	if err != nil {
		panic(err)
	}
	startTime := time.Now()
	characteristicVectors := createBinaryMatrix(repository.Size)
	// building bags of items
	knapsacks := make([]Knapsack, int(math.Pow(2, float64(repository.Size))))
	for i, iv := range characteristicVectors {
		knapsack := Knapsack{}
		for i, jv := range iv {
			if jv == 1 {
				knapsack.TotalValue += repository.Items[i].Value
				knapsack.TotalWeight += repository.Items[i].Weight
				knapsack.Items = append(knapsack.Items, repository.Items[i])
				knapsack.CharacteristicVector = iv
			}
		}
		knapsacks[i] = knapsack
	}
	// sorting from the biggest to the smallest
	sort.Slice(knapsacks, func(i, j int) bool {
		return knapsacks[i].TotalValue > knapsacks[j].TotalValue
	})

	// searching for the perfect knapsack
	fmt.Println("Knapsack's capacity:", repository.Capacity)
	for _, v := range knapsacks {
		if v.TotalWeight <= repository.Capacity {
			fmt.Println("Characteristic vector:", v.CharacteristicVector)
			fmt.Println("Total value:", v.TotalValue)
			fmt.Println("Total weight:", v.TotalWeight)
			fmt.Println("Items:")
			for _, item := range v.Items {
				fmt.Println("id:", item.Id, "weight:", item.Weight, "value:", item.Value)
			}
			break
		}
	}
	duration := time.Since(startTime)
	fmt.Println("Time took:", duration)
}

func createBinaryMatrix(size int) [][]int {
	dimension := math.Pow(2, float64(size))
	//dimension := int(math.Pow(float64(size), 2))
	matrix := make([][]int, int(dimension))
	for i := 0.; i < dimension; i++ {
		vector := make([]int, size)
		temp := i
		for j := size - 1; j >= 0; j-- {
			vector[j] = int(math.Mod(temp, 2))
			temp = math.Floor(temp / 2)
		}
		matrix[int(i)] = vector
	}
	return matrix
}
