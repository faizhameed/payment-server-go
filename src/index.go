package main

func squareRoot(x float64) float64 {

	for i := float64(1); i < float64(x/2); i += 1 {
		if i*i == x {
			return i
		}
	}
	return float64(0)
}

func main() {
	n := squareRoot(4)
	println(n)
}
