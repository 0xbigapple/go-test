package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64
    var operator string

    // Get first number from user
    fmt.Print("请输入第一个数字: ")
    fmt.Scanln(&num1)

    // Get second number from user
    fmt.Print("请输入第二个数字: ")
    fmt.Scanln(&num2)

    // Get operator from user
    fmt.Print("请输入运算符(+ 或 -): ")
    fmt.Scanln(&operator)

    // Calculate result based on operator
    switch operator {
    case "+":
        fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, num1+num2)
    case "-":
        fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, num1-num2)
    default:
        fmt.Println("不支持的运算符，只支持 + 和 -")
    }
} 