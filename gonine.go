package main
import "fmt"

func main(){
fmt.Println("---SLICES---")
fmt.Println("Declaration by using make()")
slc1 := make([]int,5)
slc1[2]=12
slc1[4]=19
fmt.Println(slc1)

fmt.Println("Accessing..")
fmt.Println(slc1[2],slc1[3],slc1[4])

fmt.Println("Like in array, len() returns length")
fmt.Println("Length of slice1 is ",len(slc1))

fmt.Println("Slices can be expanded by using append() function")
slc1=append(slc1,22)
slc1=append(slc1,33,11)
fmt.Println("New length of slice1 is",len(slc1))
fmt.Println("Last element is ",slc1[len(slc1)-1])


fmt.Println("Slices are copied among themselves")
slc2 :=make([]int,8)
copy(slc2,slc1)
fmt.Println("New copy of slice is ",slc2)
fmt.Println("Partial copying is also allowed")
slc3 :=make([]int,4)
copy(slc3,slc1)
fmt.Println("New copy of slice is ",slc3)

fmt.Println("Setting range of indices to copy -->")
slc4 := slc1[5:8]
fmt.Println("New slice is ",slc4)

slc5 :=slc1[:5]
fmt.Println("slc5 :=slc1[:5] -->",slc5)
slc6 :=slc1[6:]
fmt.Println("slc6 :=slc1[6:]--> ",slc6)

fmt.Println("2D slices ..")
/*twoD :=make([][]int,4)
for i:=0;i<4;i++{
        innerlen:=i+1
        twoD[i]:=make([]int,innerlen)
        for j:=0;j<innerlen;j++{
                              twoD[i][j]=(i+j+1)
                                }
                }
fmt.Println("The 2D slice created is :: ")
                for i:=0;i<4;i++{
                        innerlen:=i+1
                        twoD[i]:=make([]int,innerlen)
                        for j:=0;j<innerlen;j++{
                                              fmt.Printf("%d ",twoD[i][j]=(i+j+1))
                                                }
                                fmt.Printf("\n")
                                }
*/
twoD := make([][]int, 3)
   for i := 0; i < 3; i++ {
       innerLen := i + 1
       twoD[i] = make([]int, innerLen)
       for j := 0; j < innerLen; j++ {
           twoD[i][j] = i + j
       }
   }
   fmt.Println("2d: ", twoD)
}
