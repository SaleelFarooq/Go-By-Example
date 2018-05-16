package main
import "fmt"

func main(){
var preamble string
preamble="..Variables.."
fmt.Println(preamble)
a:=1
b:=2.35
c:="abcdef"
fmt.Println("a= %d b=%f32 c=%s64",a,b,c)
fmt.Println("abc"+c)
fmt.Println(1+2+3+a+2)
}
