package main
import "fmt"

func main() {
    grade := "A"

    switch grade {
    case "A":
        fmt.Println("Excellent")
    case "B":
        fmt.Println("Good")
    case "C":
        fmt.Println("Average")
    case "D":
        fmt.Println("Below Average")
    case "F":
        fmt.Println("Fail")
    default:
        fmt.Println("Invalid grade")
    }
}