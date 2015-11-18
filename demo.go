package main


import "fmt"
import _ "math"
import "math/big"

import "github.com/ricolau/evaluator"




func main(){
a := "((1 -3 +3*4)+5 > 100) && (2<3)"
//a = "4"
a = " (2<3 || 2<1) && 3 <4"




result, err := evaluator.Eval(a)


ret, err := evaluator.Evaluate(a)
fmt.Println(ret,2222222222222222222)

fmt.Println(result,777, evaluator.BigratToBool(result), err)


}
