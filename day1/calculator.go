package main
import "fmt"

func main() {
    a := 10
    b := 5
    op := "+"

    switch op {
    case "+":
        fmt.Println("Result:", a+b)
    case "-":
        fmt.Println("Result:", a-b)
    case "*":
        fmt.Println("Result:", a*b)
    case "/":
        fmt.Println("Result:", a/b)
    default:
        fmt.Println("Invalid operator")
    }
}