package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input_file = "day1_input-1.csv"

	var lcol = []int{}
	var rcol = []int{}

	f, err := os.Open(input_file)
	if err != nil {
		fmt.Println("error opening file")
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("error reading file line by line")
		}

		nums := strings.Split(line, " ")

		lnum, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("error convert left number to integer")
		}
		lcol = append(lcol, lnum)

		rnum, err := strconv.Atoi(nums[0])
		if err != nil {
			fmt.Println("error convert left number to integer")
		}
		rcol = append(rcol, rnum)
	}

}
