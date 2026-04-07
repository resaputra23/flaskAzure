package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    fmt.Println("=================================")
    fmt.Println("     Simple Calculator App")
    fmt.Println("=================================")
    fmt.Println("Operations: +, -, *, /")
    fmt.Println("Type 'exit' to quit")
    fmt.Println("=================================")
    
    for {
        fmt.Print("\nEnter calculation (e.g., 10 + 5): ")
        
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        
        if input == "exit" {
            fmt.Println("Goodbye!")
            break
        }
        
        // Parse input
        parts := strings.Fields(input)
        if len(parts) != 3 {
            fmt.Println("Invalid format! Please use: number operator number")
            continue
        }
        
        // Parse numbers
        num1, err1 := strconv.ParseFloat(parts[0], 64)
        operator := parts[1]
        num2, err2 := strconv.ParseFloat(parts[2], 64)
        
        if err1 != nil || err2 != nil {
            fmt.Println("Invalid numbers! Please enter valid numbers")
            continue
        }
        
        // Calculate result
        var result float64
        var validOp bool = true
        
        switch operator {
        case "+":
            result = num1 + num2
        case "-":
            result = num1 - num2
        case "*":
            result = num1 * num2
        case "/":
            if num2 == 0 {
                fmt.Println("Error: Division by zero!")
                continue
            }
            result = num1 / num2
        default:
            fmt.Println("Invalid operator! Use +, -, *, or /")
            validOp = false
        }
        
        if validOp {
            fmt.Printf("%.2f %s %.2f = %.2f\n", num1, operator, num2, result)
        }
    }
}
