package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func day1_star1() {
	var input_file = "./day1_input.txt"

	var lcol = []int{}
	var rcol = []int{}
	var totalDiff = 0

	f, _ := os.Open(input_file)
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		nums := strings.Split(line, "   ")

		lnum, _ := strconv.Atoi(nums[0])
		rnum, _ := strconv.Atoi(strings.TrimSuffix(nums[1], "\r\n"))

		lcol = append(lcol, lnum)
		rcol = append(rcol, rnum)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("error reading file line by line")
		}
	}

	slices.Sort(lcol)
	slices.Sort(rcol)

	fmt.Println(len(lcol))

	for i := 0; i < len(lcol); i++ {
		var diff = lcol[i] - rcol[i]
		if diff < 0 {
			diff *= -1
		}
		//fmt.Printf("lcol: %d, rcol: %d, diff: %d\n", lcol[i], rcol[i], diff)
		totalDiff += diff
	}

	fmt.Println("total diff", totalDiff)
}
