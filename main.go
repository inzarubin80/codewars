package main
import "fmt"
import 	"github.com/inzarubin80/codewars/ProductOfConsecutiveFibNumbers"

func main() {
	res := ProductOfConsecutiveFibNumbers.ProductFib((uint64)(10000))
	fmt.Println(res)
}
