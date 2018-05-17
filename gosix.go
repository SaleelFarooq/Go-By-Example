package main
import "fmt"

func main(){
fmt.Println("Conditionals -- If-Else")
a:=12
if(a>0){
        fmt.Println("a is a +ve number")
        }else{
    fmt.Println("a is a -ve number")
    }

fmt.Println("If can come without else")
if a%2==0{
          fmt.Println("a is even")
          }

fmt.Println("declaration on starting branch is applicable to all branches")
if num:=11;num<0{
  fmt.Println("num is a negative number")
  }else{
    fmt.Println("num is a positive number")
  }
fmt.Println("Illustration of Else if ladder")
if k:=13;k==0{fmt.Println("k is Zero")
  }else if(k%2==1){
    fmt.Println("k is an odd number")
    }else{
      fmt.Println("k is an even number")
    }
}
