package main

import "fmt"

func main() {
	var total, diskon, potongan, totalAkhir int

	fmt.Scan(&total)
	fmt.Scan(&diskon)

	potongan =total * diskon / 100
	totalAkhir = total - potongan

	fmt.Println(totalAkhir)
}
