package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getFileContents() string {
	data, err := os.ReadFile("./input")

	if err != nil {
		//apparently, this will kill the function and, therefore, the process
		panic(err)
	}

	// type conversion. data should be a byte array, but my return
	return string(data)
}

func getCalibrationValuesByLine(line string) int {
	first := ""
	second := ""

	i := 0
	for i < len(line) {
		_, err := strconv.Atoi(string(line[i]))
		if err == nil {
			first = string(line[i])
			break
		}

		i++
	}

	j := len(line) - 1
	for j >= 0 {
		_, err := strconv.Atoi(string(line[j]))
		if err == nil {
			second = string(line[j])
			break
		}

		j--
	}

	calibrationValue, _ := strconv.Atoi(first + second)
	return calibrationValue
}

func main() {
	contents := getFileContents()

	lines := strings.Split(contents, "\n")

	// removing final line break
	lines = lines[:len(lines)-1]

	calibrationValue := 0

	for _, line := range lines {
		calibrationValue += getCalibrationValuesByLine(line)
	}

	fmt.Println(calibrationValue)
}
