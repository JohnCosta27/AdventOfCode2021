package day13

import (
	"fmt"
	"strings"
)

type coord struct {
	x int
	y int
} 

func Run(dat []byte) (int, int) {

	stringData := string(dat)
	splitData := strings.Split(stringData, "\n\n")

	fmt.Println(splitData)

	return 0, 0

}