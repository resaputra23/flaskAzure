package main

import (
    "errors"
    "fmt"
    "strings"
)

// Calculator struct
type Calculator struct {
    history []string
}

// Add two numbers
func (c *Calculator) Add(a, b float64) float64 {
    result := a + b
    c.addToHistory(fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
    return result
}

// Subtract two numbers
func (c *Calculator) Subtract(a, b float64) float64 {
    result := a - b
    c.addToHistory(fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
    return result
}

// Multiply two numbers
func (c *Calculator) Multiply(a, b float64) float64 {
    result := a * b
    c.addToHistory(fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
    return result
}

// Divide two numbers
func (c *Calculator) Divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero is not allowed")
    }
    result := a / b
    c.addToHistory(fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
    return result, nil
}

// addToHistory stores calculation in history
func (c *Calculator) addToHistory(entry string) {
    c.history = append(c.history, entry)
}

// GetHistory returns calculation history
func (c *Calculator) GetHistory() []string {
    return c.history
}

// ClearHistory clears calculation history
func (c *Calculator) ClearHistory() {
    c.history = []string{}
}

// Fibonacci generates Fibonacci sequence
func Fibonacci(n int) ([]int, error) {
    if n < 0 {
        return nil, errors.New("n cannot be negative")
    }
    if n == 0 {
        return []int{}, nil
    }
    if n == 1 {
        return []int{0}, nil
    }
    
    sequence := make([]int, n)
    sequence[0] = 0
    sequence[1] = 1
    
    for i := 2; i < n; i++ {
        sequence[i] = sequence[i-1] + sequence[i-2]
    }
    return sequence, nil
}

// IsPalindrome checks if a string is palindrome
func IsPalindrome(s string) bool {
    s = strings.ToLower(strings.ReplaceAll(s, " ", ""))
    for i := 0; i < len(s)/2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

// Factorial calculates factorial of a number
func Factorial(n int) (int, error) {
    if n < 0 {
        return 0, errors.New("factorial is not defined for negative numbers")
    }
    if n == 0 || n == 1 {
        return 1, nil
    }
    
    result := 1
    for i := 2; i <= n; i++ {
        result *= i
    }
    return result, nil
}

// main function
func main() {
    fmt.Println("=== Simple Go Application ===")
    fmt.Println("Version: 1.0.0")
    fmt.Println()
    
    // Calculator demo
    calc := &Calculator{}
    
    fmt.Println("--- Calculator Demo ---")
    fmt.Printf("10 + 5 = %.2f\n", calc.Add(10, 5))
    fmt.Printf("10 - 5 = %.2f\n", calc.Subtract(10, 5))
    fmt.Printf("10 * 5 = %.2f\n", calc.Multiply(10, 5))
    
    if result, err := calc.Divide(10, 5); err == nil {
        fmt.Printf("10 / 5 = %.2f\n", result)
    }
    
    if result, err := calc.Divide(10, 0); err != nil {
        fmt.Printf("Error: %v\n", err)
    }
    
    fmt.Println("\n--- Calculation History ---")
    for i, entry := range calc.GetHistory() {
        fmt.Printf("%d. %s\n", i+1, entry)
    }
    
    // Fibonacci demo
    fmt.Println("\n--- Fibonacci Sequence ---")
    fib, _ := Fibonacci(10)
    fmt.Printf("First 10 Fibonacci numbers: %v\n", fib)
    
    // Palindrome demo
    fmt.Println("\n--- Palindrome Checker ---")
    testStrings := []string{"racecar", "hello", "A man a plan a canal panama", "golang"}
    for _, s := range testStrings {
        if IsPalindrome(s) {
            fmt.Printf("✓ '%s' is a palindrome\n", s)
        } else {
            fmt.Printf("✗ '%s' is not a palindrome\n", s)
        }
    }
    
    // Factorial demo
    fmt.Println("\n--- Factorial Calculator ---")
    numbers := []int{0, 1, 5, 7, -3}
    for _, n := range numbers {
        if result, err := Factorial(n); err == nil {
            fmt.Printf("%d! = %d\n", n, result)
        } else {
            fmt.Printf("Error: %v\n", err)
        }
    }
    
    fmt.Println("\n✅ Application completed successfully!")
}
