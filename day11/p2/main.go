package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	itemWorryLevels                                   []int
	testDivisor                                       int
	throwToMonkeyDivisible, throwToMonkeyNotDivisible int
	monkeyBusiness                                    int
	UpdateWorryLevel                                  func(worryLevel, productRemainder int) int
}

func (monkey *Monkey) SelectNextMonkeyIndex(itemIndex int) (monkeyIndex int) {
	if monkey.itemWorryLevels[itemIndex]%monkey.testDivisor == 0 {
		return monkey.throwToMonkeyDivisible
	}
	return monkey.throwToMonkeyNotDivisible
}

func (monkey *Monkey) ThrowItem(itemIndex int, to *Monkey) {
	to.itemWorryLevels = append(to.itemWorryLevels, monkey.itemWorryLevels[itemIndex])
	monkey.itemWorryLevels[itemIndex] = -1
}

var ops = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"*": func(a, b int) int { return a * b },
}

func StartingItemsParser(line string) ([]int, error) {
	items := []int{}
	worryLevelsRegexp := regexp.MustCompile(`\d+`)
	itemWorryLevelsMatches := worryLevelsRegexp.FindAllString(line, -1)
	for _, worryLevelStr := range itemWorryLevelsMatches {
		worryLevel, err := strconv.Atoi(worryLevelStr)
		if err != nil {
			fmt.Println("Unable to parse worry level:", err)
			return nil, err
		}
		items = append(items, worryLevel)
	}
	return items, nil
}

func MonkeyInputParser(lines []string) (Monkey, error) {
	monkey := Monkey{}
	for _, line := range lines {
		switch {
		case strings.Contains(line, "Starting items:"):
			itemWorryLevels, err := StartingItemsParser(line)
			if err != nil {
				return monkey, err
			}
			monkey.itemWorryLevels = itemWorryLevels
		case strings.Contains(line, "Operation:"):
			operationRegexp := regexp.MustCompile(`= old (?P<Operator>[+*]) (?P<Operand>.+)`)
			operationMatches := operationRegexp.FindStringSubmatch(line)
			if operationMatches == nil {
				return monkey, fmt.Errorf("unable to parse operation: %s", line)
			}
			operator := operationRegexp.SubexpIndex("Operator")
			operand := operationRegexp.SubexpIndex("Operand")
			if operationMatches[operand] == "old" {
				monkey.UpdateWorryLevel = func(worryLevel, productRemainder int) int {
					return ops[operationMatches[operator]](worryLevel, worryLevel) % productRemainder
				}
			} else {
				operandWorryLevel, err := strconv.Atoi(operationMatches[operand])
				if err != nil {
					return monkey, err
				}
				monkey.UpdateWorryLevel = func(worryLevel, productRemainder int) int {
					return ops[operationMatches[operator]](worryLevel, operandWorryLevel) % productRemainder
				}
			}
		case strings.Contains(line, "Test:"):
			divisibleByRegexp := regexp.MustCompile(`divisible by (?P<Divisor>\d+)`)
			divisibleMatches := divisibleByRegexp.FindStringSubmatch(line)
			if divisibleMatches == nil {
				return monkey, fmt.Errorf("unable to parse test: %s", line)
			}
			divisorStr := divisibleByRegexp.SubexpIndex("Divisor")
			divisor, err := strconv.Atoi(divisibleMatches[divisorStr])
			if err != nil {
				return monkey, err
			}
			monkey.testDivisor = divisor
		case strings.Contains(line, "throw to monkey"):
			var throwToMonkeyDivisible, throwToMonkeyNotDivisible int
			throwToMonkeyRegexp := regexp.MustCompile(`If (?P<Truthy>true|false): throw to monkey (?P<MonkeyIndex>\d+)`)
			throwToMonkeyMatches := throwToMonkeyRegexp.FindStringSubmatch(line)
			if throwToMonkeyMatches == nil {
				return monkey, fmt.Errorf("unable to parse throw to monkey: %s", line)
			}
			truthyStr := throwToMonkeyRegexp.SubexpIndex("Truthy")
			monkeyIndexStr := throwToMonkeyRegexp.SubexpIndex("MonkeyIndex")
			if throwToMonkeyMatches[truthyStr] == "true" {
				var err error
				throwToMonkeyDivisible, err = strconv.Atoi(throwToMonkeyMatches[monkeyIndexStr])
				if err != nil {
					return monkey, err
				}
				monkey.throwToMonkeyDivisible = throwToMonkeyDivisible
			}
			if throwToMonkeyMatches[truthyStr] == "false" {
				var err error
				throwToMonkeyNotDivisible, err = strconv.Atoi(throwToMonkeyMatches[monkeyIndexStr])
				if err != nil {
					return monkey, err
				}
				monkey.throwToMonkeyNotDivisible = throwToMonkeyNotDivisible
			}
		}
	}
	return monkey, nil
}

func main() {
	var monkeys []Monkey
	var monkeyData []string
	productRemainder := 1

	f, err := os.Open("./input")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			monkeyData = append(monkeyData, line)
		} else {
			monkey, err := MonkeyInputParser(monkeyData)
			if err != nil {
				log.Fatal(err)
			}
			monkeys = append(monkeys, monkey)
			monkeyData = []string{}
			productRemainder *= monkey.testDivisor
		}
	}
	if len(monkeyData) != 0 {
		monkey, err := MonkeyInputParser(monkeyData)
		if err != nil {
			log.Fatal(err)
		}
		monkeys = append(monkeys, monkey)
		productRemainder *= monkey.testDivisor
	}
	for rounds := 0; rounds < 10000; rounds++ {
		for monkeyIndex, monkey := range monkeys {
			for i, item := range monkey.itemWorryLevels {
				monkey.itemWorryLevels[i] = monkey.UpdateWorryLevel(item, productRemainder)
				nextMonkeyIndex := monkey.SelectNextMonkeyIndex(i)
				monkey.ThrowItem(i, &monkeys[nextMonkeyIndex])
				monkey.monkeyBusiness += 1
			}
			monkeys[monkeyIndex].itemWorryLevels = []int{}
			monkeys[monkeyIndex].monkeyBusiness = monkey.monkeyBusiness
		}
	}
	// Sort monkeys by their monkey business
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].monkeyBusiness > monkeys[j].monkeyBusiness
	})
	fmt.Println("The monkey business is:", monkeys[0].monkeyBusiness*monkeys[1].monkeyBusiness)
	f.Close()
}
