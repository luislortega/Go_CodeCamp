package main

//We can import multiple packages
import (
	"fmt"
	"helloworld/03_packages/strucutil"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(strucutil.Reverse("hola"))
}
