package main
import "fmt"

const hargaKopi = 20000
const diskon = 1.11
func main() {
	var jumlah, total int
	fmt.Scan(&jumlah)
	total = jumlah * hargaKopi
	var hargaDiskon float64
	hargaDiskon =float64(total) * float64(diskon)
	fmt.Printf("%.0f", hargaDiskon)
}
