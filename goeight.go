package main
import "fmt"

func main(){
fmt.Println("--ARRAYS--")
fmt.Println("Non initialised => Zero values for all")
var arr1 [5]int
var arr2 [6]string

fmt.Println(arr1)
fmt.Println(arr2)

fmt.Println("Initialisation")
arr1[3]=21
arr1[1]=15
arr2[2]="Reus"
arr2[4]="Bale"
fmt.Println("After initialisation ---> ")
fmt.Println(arr1)
fmt.Println(arr2)

fmt.Println("Accessing")
fmt.Printf("arr1[3] = %d \n",arr1[3])
fmt.Printf("arr2[4] = %s \n",arr2[4])

fmt.Println("declaration + initialisation -->")
/*var arr3 [6]int
arr3={11,13,56,9,43}
fmt.Println(arr3)
*/
arr4 :=[3]string{"India","Japan","Canada"}
fmt.Println(arr4)
fmt.Printf("Length of array4 is %d\n",len(arr4))
fmt.Println("Multi dimensional arrays exist")
var arr5 [2][3]int
for i:=0;i<2;i++{
            for j:=0;j<3;j++{
                            arr5[i][j]=(i+j)*(i+j)
                              }
                }
fmt.Println("The 2D array thus declared is ")
for i:=0;i<2;i++{
            for j:=0;j<3;j++{
                            fmt.Printf("%d ",arr5[i][j])
                              }
              fmt.Printf("\n")
                }
}
