package main
import "fmt"

func main() {
    num := -5

    if num > 0 {
        fmt.Println("Positive number")
    } else if num < 0 {
        fmt.Println("Negative number")
    } else {
        fmt.Println("Zero")
    }
}