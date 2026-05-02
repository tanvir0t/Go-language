package main
import "fmt"

func main() {
    marks := 75

    if marks >= 80 {
        fmt.Println("Grade: A+")
    } else if marks >= 70 {
        fmt.Println("Grade: A")
    } else if marks >= 60 {
        fmt.Println("Grade: B")
    } else if marks >= 50 {
        fmt.Println("Grade: C")
    } else {
        fmt.Println("Fail")
    }
}