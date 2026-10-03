package main

import (
	"bufio"
	"fmt"
	"linear-stats/calculators"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := []int{}

	if len(os.Args) != 2 { // Error handling for user's arguments
		fmt.Println("Unacceptable user input")
		return
	}
	args := os.Args[1]
	file, err := os.Open(args) // open file from arguments for reading
	if err != nil {
		fmt.Println("Error in reading file:", args)
		return
	}
	defer file.Close() // close file when done reading

	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // read file line by line
		line := strings.TrimSpace(scanner.Text())
		if line == "" { // skip an empty line
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Error in string to number conversion")
			return
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error in reading line from user's input")
		return
	}

	a, b, err := calculators.LinearRegressionLine(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Linear Regression Line: y = %fx + %f\n", a, b)

	r, err := calculators.PearsonCorrelationCoefficient(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Pearson Correlation Coefficient: %.10f\n", r) // print 10 decimals with %.10f
}
