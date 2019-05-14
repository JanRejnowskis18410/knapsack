package main

const digits = "01"

// decToBin converts decimal int into its binary representation
// output:
// 49 in []byte represents 1
// 48 and 0 in []byte represents 0
func decToBin(x, size int) []byte {
	var a [64 + 1]byte
	i := len(a)
	for x >= 2 {
		i--
		a[i] = digits[x&1]
		x >>= 1
	}
	// x < 2
	i--
	a[i] = digits[x]
	return a[len(a)-size:]

	////base2 := strconv.FormatInt(int64(i), 2)
	////vector := strings.Repeat("0", itemsSize-len(base2)) + base2
	//binaryRepresentation := make([]int, size)
	//temp := x
	//for j := size - 1; j >= 0; j-- {
	//	binaryRepresentation[j] = int(math.Mod(float64(temp), 2))
	//	temp = int(math.Floor(float64(temp / 2)))
	//}
	//return binaryRepresentation
}
