package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getLinesInFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	file.Close()

	return fileLines, nil
}

func containsSequence(line string, sequences []string, threshold int) bool {
	count := 0
	for _, sequence := range sequences {
		count += strings.Count(line, sequence)
	}

	return count >= threshold
}

func containsPairOfLetters(line string) bool {
	count := 0
	for i := 0; i < len(line)-3; i++ {
		sequence := line[i : i+4]
		if sequence[0] == sequence[2] {
			if sequence[0:2] == sequence[2:] {
				count += 1
				continue
			}
			continue
		}

		if strings.Count(line, sequence[0:2]) < 2 {
			continue
		}

		count += 1
	}

	return count >= 1
}

func containsRepeatOfLetter(line string) bool {
	count := 0
	for i := 0; i < len(line)-2; i++ {
		sequence := line[i : i+3]
		if sequence[0] != sequence[2] {
			continue
		}
		count += 1
	}

	return count >= 1
}

func partOneCount(lines []string) int {
	vowels := []string{"a", "e", "i", "o", "u"}
	forbidden := []string{"ab", "cd", "pq", "xy"}
	// lmao, its late lets just hardcode them
	doubleLetters := []string{
		"aa", "bb", "cc", "dd", "ee", "ff", "gg",
		"hh", "ii", "jj", "kk", "ll", "mm", "nn",
		"oo", "pp", "qq", "rr", "ss", "tt", "uu",
		"vv", "ww", "xx", "yy", "zz",
	}

	totalCount := 0
	for _, line := range lines {
		lineHasThreeVowels := containsSequence(line, vowels, 3)
		if !lineHasThreeVowels {
			continue
		}
		lineHasForbiddenSequence := containsSequence(line, forbidden, 1)

		if lineHasForbiddenSequence {
			continue
		}

		lineHasDoubleLetter := containsSequence(line, doubleLetters, 1)
		if !lineHasDoubleLetter {
			continue
		}

		totalCount += 1
	}
	return totalCount
}

func partTwoCount(lines []string) int {
	totalCount := 0
	for _, line := range lines {
		containsPair := containsPairOfLetters(line)
		hasRepeat := containsRepeatOfLetter(line)
		if containsPair && hasRepeat {
			totalCount += 1
		}
	}

	return totalCount
}

func main() {
	fileContents, err := getLinesInFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	firstCount := partOneCount(fileContents)
	secondCount := partTwoCount(fileContents)

	fmt.Printf("%d\n", firstCount)
	fmt.Printf("%d\n", secondCount)
}
