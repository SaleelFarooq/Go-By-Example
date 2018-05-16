package main
import "fmt"

func main(){
fmt.Println("LOOP means for loop")
for i:=0;i<=10;i++{
                  fmt.Printf("%d ",i)
                  }
fmt.Println("")
fmt.Println("While loop is set up by using for loop with single condition")
j:=1
for j<=15{fmt.Printf("%d ",j)
          j++
          }

fmt.Println("Use of break and continue")
k:=1
for k<=20{if(k%2==1){k++
                    continue
                    }
        fmt.Printf("%d ",k)
        k++
          }
fmt.Println("")
k=1;
for true{
          if(k>20){
                  break
                  }
          if(k%2==0){
                    fmt.Printf("%d ",k)
                    }
          k++
          }
fmt.Println("")
for i:=1;i<21;i++{if(i%2==1){continue
                            }
                  fmt.Printf("%d ",i)
                  }
}
