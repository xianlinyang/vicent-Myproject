package hex_decodestring

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

func main(){
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)
	// [72 101 108 108 111]
	fmt.Println(src)
	// 48656c6c6f -> 48(4*16+8=72) 65(6*16+5=101) 6c 6c 6f
	fmt.Println(encodedStr)

	test, _ := hex.DecodeString(encodedStr)
	fmt.Println(string(test))
	fmt.Println(bytes.Compare(test, src)) // 0
}
