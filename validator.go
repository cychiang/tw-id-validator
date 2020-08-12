package validator

var unifyArray = []int{8, 7, 6, 5, 4, 3, 2, 1}
var lookUpTable = map[string]int{
	"A": 1,
	"B": 10,
	"C": 9,
	"D": 8,
	"E": 7,
	"F": 6,
	"G": 5,
	"H": 4,
	"J": 3,
	"K": 2,
	"L": 2,
	"M": 11,
	"N": 10,
	"P": 9,
	"Q": 8,
	"R": 7,
	"S": 6,
	"T": 5,
	"U": 4,
	"V": 3,
	"X": 3,
	"Y": 12,
	"W": 11,
	"Z": 10,
	"I": 9,
	"O": 8,
}

func idForeignValid(id string) bool {
	var sum = lookUpTable[string(id[0])]
	checkNumber := int(id[len(id)-1]) - 48
	for idx, val := range id[1 : len(id)-1] {
		sum += ((int(val) - 48) * unifyArray[idx]) % 10
	}
	if (sum%10+checkNumber)%10 == 0 {
		return true
	} else {
		return false
	}
}

func IdForeign(id string) bool {
	return idForeignValid(id)
}
