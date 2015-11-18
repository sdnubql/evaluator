package main


import "fmt"
import _ "math"
import "math/big"

import "github.com/ricolau/evaluator"




func main(){
a := "((1 -3 +3*4)+5 > 100) && (2<3)"
//a = "4"
a = " (2<3 || 2<1) && 3 >4"

/*
-1 vs 0, -1
0 vs 0, 0
1 vs 0, 1

*/

rtrue := big.NewRat(0,1)


result, err := evaler.Eval(a)

fmt.Println(result,rtrue,777, evaler.RatToBool(result), err)


}
