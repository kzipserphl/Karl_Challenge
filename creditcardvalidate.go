package main

import (
    "fmt"
    "strings"
)

var testValues = []string {
	"6",
	"4123456789123456",
	"5123-4567-8912-3456",
	"61234-567-8912-3456",
	"5133-3367-8912-3456",
	"5123 - 3567 - 8912 - 3456",
	"44244x4424442444",
	"0525362587961578",
}
var repeaters = []string {
	"0000","1111","2222","3333","4444","5555","6666","7777","8888","9999",
}

func hasAtLeast(testValue string) bool {
	if len(testValue) < 16 {
		return false
	}
	return true
}

func hasCorrectCharacters(testValue string) bool {
	var correct = true
	for i := 0; i < len(testValue); i++ { 
		if testValue[i] < '0' || testValue[i] > '9' {
			if testValue[i] != '-' {
				correct = false
			}
		}
	}
	return correct
}

func startsCorrect(testValue string) bool {
	if testValue[0] == '4' || testValue[0] == '5' || testValue[0] == '6' {
		return true
	}
	return false
}

func splitsCorrect(testValue string) bool {
	arr := strings.Split(testValue,"-")
	var correct = true
	for i := 0; i < len(arr); i++ { 
		if len(arr[i]) != 4 {
			correct = false
		}
	}
	return correct
}

func passesRepeatersTest(testValue string) bool {
	var correct = true
	for i := 0; i < len(testValue)-3; i++ { 
		var thisValue = testValue[i:i+4]
		// fmt.Println(thisValue)
		for j := 0; j < len(repeaters); j++ { 
			if thisValue == repeaters[j] {
				correct = false
			}
		}
	}
	return correct
}


func main() {

	for i := 0; i < len(testValues); i++ { 
		var thisValue = testValues[i]

		if !startsCorrect(thisValue) {
			fmt.Println(thisValue, ": Invalid (must start with 4,5, or 6)")
			continue
		}		
		if !hasAtLeast(thisValue) {
			fmt.Println(thisValue, ": Invalid (at least 16 digits required)")
			continue
		}
		if !hasCorrectCharacters(thisValue) { 
			fmt.Println(thisValue, ": Invalid (0-9 or '-' are valid characters)")
			continue
		}
		var checkValue = thisValue
		if len(thisValue) > 16 {
			if !splitsCorrect(thisValue) {
				fmt.Println(thisValue, ": Invalid (uneven count in the separated groups)")
				continue
			} else {
				arr := strings.Split(thisValue,"-")
				checkValue = strings.Join(arr, "")
			}
		}	
		
		if !passesRepeatersTest(checkValue) { 
			fmt.Println(thisValue, ": Invalid (repeating numbers gtr. than 3")
			continue
		}
		
		fmt.Println(thisValue, ": Valid")
	}
}