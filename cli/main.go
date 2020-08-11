package main

import (
	"fmt"
	"github.com/cychiang/tw-id-validator"
)

func main() {
	fmt.Println(validator.IdForeign("A800000014"))
	fmt.Println(validator.IdForeign("A900000014"))
}
