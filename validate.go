package main

import ("fmt"
    "strings"
)

var flagPrintDebug = false  // set to true, save, and run the file to see messages

func printDebug(message string) {
    if flagPrintDebug {
        fmt.Println(message)
    }
}

func validateConsecutive(s string) bool {
    s = strings.Replace(s, "-", "", -1) // last validation; replace dashes
    printDebug (" - " + s + " (consecutive digits check)")
    //
    var lastChar rune
    var lastCharCount = 0
    for _, c := range s {
        if c == lastChar {
            lastCharCount++
            if lastCharCount >= 4 {
                printDebug (" - found at least 4 consecutive values")
                return false    
            }
        } else {
            lastChar = c
            lastCharCount = 1
        }
    }

    return true // no consecutive values
}

func validateCC(p_value string) bool {
    
    if strings.Contains(p_value, " ") {
        printDebug (" - found spaces")
        return false
    }
    
    if strings.Contains(p_value, "-") {
        arrs := strings.Split(p_value, "-")
        for i := 0; i < len(arrs); i++ {
            if len(arrs[i]) != 4 {
                printDebug (" - found group of digits not a 4 count")
                return false
            }
        }
    }

    numCount := 0
    for i := 0; i < len(p_value); i++ {
        var thisValue string = p_value[i:(i+1)];
        if thisValue >= "1" && thisValue <= "9" {
            numCount++;
        } else {
            if thisValue != "-" {
                printDebug (" - found invalid character (not 1-9 or -)")
                return false
            }
        }
        
    }
    if numCount != 16 {
        printDebug (" - found number of digits not 16")
        return false    
    }

    if !validateConsecutive(p_value) {
        return false // found consecutive values 
    }

    return true
}

func main() {

    arr := []string{
        "6",
        "4123x456789123456",
        "5123-4567-8912-3456",
        "61234-567-8912-3456",
        "4123356789123456",
        "5133-3367-8912-3456",
        "5123 - 3567 - 8912 - 3456",
    }
    for index, element := range arr {
        if index == 0 {
            continue
        }

        printDebug (element)
        var checked = ""
        if !validateCC(element) {
            checked = "invalid"
        } else {
            checked = "Valid"
        }

        fmt.Println(checked)
        printDebug (" ")
    }

}
