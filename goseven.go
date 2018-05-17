package main
import "fmt"
import "time"

func main(){
fmt.Println("illustration of Switch")
c:='e'
switch c{case 'a','e','i','o','u' :
                            fmt.Println("Vowel")
        default :
                      fmt.Println("Consonent")
        }
w:=time.Now().Weekday()
fmt.Println("Today is ",w)
switch w{
case time.Saturday,time.Sunday:
          fmt.Println("it's a weekend")
case time.Monday:
          fmt.Println("It's a monday")
default :
          fmt.Println("A weekday")
        }

h:=time.Now().Hour()
fmt.Printf("Time is now around %d Ó clock" ,h)
fmt.Println("Switch without variable with it")
t:=time.Now()
fmt.Printf("Time ==== ",t)
fmt.Println("")
switch {
case h<12:
      fmt.Println("Forenoon")
default :
      fmt.Println("After noon")
        }
}
