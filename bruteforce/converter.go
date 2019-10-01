package bruteforce

const digits = "01"

// DecToBin converts fast decimal int into its binary representation
// output:
// 49 in []byte represents 1
// 48 and 0 in []byte represents 0
func DecToBin(x, size int) []byte {
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
}
