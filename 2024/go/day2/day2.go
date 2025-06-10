package day2

import (
	"bufio"
	"fmt"
	"os"
)

func Day2() {
	day2_star1()
}

func day2_star1() {
	// which lines from the file are "safe"
	// based on only increasing/decreasing numbers
	// on each line from right-to-left by 1-3

	var safe_reports = 0

	//var input_file = "./input_day2.txt"
	var input_file = "./day1_input.txt"
	f, err := os.Open(input_file)
	if err != nil {
		fmt.Println("error opening file", err.Error())
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		line, err := reader.ReadString('\n')
		//nums := strings.Split(line, " ")
		fmt.Println("error reading line ", line)

		// for i := 0; i < len(nums); i++ {
		// 	lnum, _ := strconv.Atoi(nums[i])
		// 	if i+1 >= len(nums) {
		// 		safe_reports++
		// 		break
		// 	}
		// 	rnum, _ := strconv.Atoi(strings.TrimSuffix(nums[i+1], "\r\n"))
		// 	if lnum-rnum > 3 || lnum-rnum < -3 {
		// 		break
		// 	}
		// }

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println(err.Error())
		}
	}

	fmt.Println(safe_reports)
}
