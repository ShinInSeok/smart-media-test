package test2480

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Test2480() {
    br := bufio.NewReader(os.Stdin)
    str, _ := br.ReadString(('\n'))

    diceNum := strings.TrimSpace(str)

    blank := strings.Fields(diceNum)

	for i:=0; i < 3; i++ {
		countDice := strings.Count(diceNum, blank[i])
		if countDice == 3 {
			num, _ := strconv.Atoi(blank[i])
			result := 10000 + num * 1000
			fmt.Println(result)
			return
		} else if countDice == 2{
			num, _ := strconv.Atoi(blank[i])
			result := 1000 + num * 100
			fmt.Println(result)
			return
		} else {
			for j:=0; j < 3; j++{
				num1, _ := strconv.Atoi(blank[i])
				num2, _ := strconv.Atoi(blank[j])
				num3, _ := strconv.Atoi(blank[2])
				num4, _ := strconv.Atoi(blank[1])

				if num3 == num4{
					continue
				} else {

					if num1 < num2{
						num1, num2 = num2, num1
						if num1 < num3{
							continue
						} else {
							result := num1 * 100
							fmt.Println(result)
							return
						}
					}
				}
			}
		}
	}
}