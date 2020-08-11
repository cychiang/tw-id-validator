package validator

import (
	"fmt"
	"regexp"
)

var unifyArray = []int{1, 9, 8, 7, 6, 5, 4, 3, 2, 1}
var mapAlphabet = map[string][]int{
	"A": {1, 0},
	"B": {1, 1},
	"C": {1, 2},
	"D": {1, 3},
	"E": {1, 4},
	"F": {1, 5},
	"G": {1, 6},
	"H": {1, 7},
	"J": {1, 8},
	"K": {1, 9},
	"L": {2, 0},
	"M": {2, 1},
	"N": {2, 2},
	"P": {2, 3},
	"Q": {2, 4},
	"R": {2, 5},
	"S": {2, 6},
	"T": {2, 7},
	"U": {2, 8},
	"V": {2, 9},
	"X": {3, 0},
	"Y": {3, 1},
	"W": {3, 2},
	"Z": {3, 3},
	"I": {3, 4},
	"O": {3, 5},
}

func arrayExpand(id string) ([]int, int) {
	idArray := mapAlphabet[string(id[0])]
	for _, val := range id[1:len(id)-1] {
		idArray = append(idArray, int(val) - 48)
	}
	checkNumber := int(id[len(id)-1]) - 48
	return idArray, checkNumber
}

func valid(idArray []int, unifyArray []int, checkNumber int) bool {
	var sum = 0
	resultArray := [10]int{0}
	for idx:=0; idx < 10; idx++{
		resultArray[idx] = (idArray[idx] * unifyArray[idx]) % 10
	}
	for _, val := range resultArray {
		sum += val
	}
	if 10 - sum == checkNumber {
		return true
	} else {
		return false
	}
}

func IdForeign(id string) {
	match, _ := regexp.Match(`^[A-Z][0-9]{9}`, []byte(id))
	if match {
		idArray, checkNumber := arrayExpand(id)
		fmt.Println(valid(idArray, unifyArray, checkNumber))
	}
}
