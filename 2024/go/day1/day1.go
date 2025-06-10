package day1

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1_star1() {
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

	fmt.Println("star1", totalDiff)

	day1_star2(lcol, rcol)
}

func day1_star2(lcol []int, rcol []int) {
	var similarityList = []int{}
	var simScore = 0

	for i := 0; i < len(lcol); i++ {
		lnum := lcol[i]
		count := 0
		for j := 0; j < len(rcol); j++ {
			if lnum == rcol[j] {
				count++
			}
		}
		score := lnum * count
		similarityList = append(similarityList, score)
	}

	for i := 0; i < len(similarityList); i++ {
		simScore += similarityList[i]
	}

	fmt.Println("star 2", simScore)
}
