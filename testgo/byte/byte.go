package byte

import "fmt"

func main() {
	bs := []byte{1, 3}
	for _, n := range bs {
		fmt.Printf("% 08b", n) // prints 00000000 11111101
	}
}
