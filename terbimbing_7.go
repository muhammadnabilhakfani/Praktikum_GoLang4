package main
import "fmt"

func main() {
	var h, batas, tekanan int
	const p = 1025.0
	const g = 9.8
	fmt.Scan(&h, &batas)
	tekanan = int(p * g * float64(h))
	fmt.Println(tekanan)
	fmt.Print(tekanan >= batas)
}
