package main

import (
	"fmt"
	"github.com/cychiang/tw-id-validator"
)

func main() {
	fmt.Println(validator.IdForeign("A800000014"))
	fmt.Println(validator.IdForeign("A900000016"))
	fmt.Println(validator.IdForeign("A870000015"))
	fmt.Println(validator.IdForeign("A970000017"))
	fmt.Println(validator.IdForeign("A880000018"))
	fmt.Println(validator.IdForeign("A980000010"))
	fmt.Println(validator.IdForeign("A890000011"))
	fmt.Println(validator.IdForeign("A990000013"))
}
