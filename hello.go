package main

// fmt - format package
import (
	"fmt"
)

var pl = fmt.Println  // pl is alias for fmt.Println


// if statements
func main() {
  // conditional ops : > < >= <= == !=
  // logical ops : && || !  

  iAge := 60


  if (iAge >= 1) && (iAge <= 18) {
    pl("important birthday")
  } else if(iAge == 21) || (iAge == 50) {
    pl("important birthday")
  } else if iAge >= 65 {
    pl("important birthday")
  } else {
    pl("not important birthday")
  }
  pl("!true =", !true)
}


// if statement
// func main() {
//   cV1 := "3.14"
//   if cV2, err := strconv.ParseFloat(cV1, 64); err == nil {
//     pl(cV2, reflect.TypeOf(cV2))
//   }
// }



// string to float
// func main() {
//   cV1 := "3.14"
//   cV2, err := strconv.ParseFloat(cV1, 64) 
//   pl(cV2, err, reflect.TypeOf(cV2))
// }


// int to string
// func main() {
//   cV5 := 500000
//   cV6 := strconv.Itoa(cV5) // Itoa - int to Ascii
//   pl(cV6, reflect.TypeOf(cV6))
// }


// func main() {
//   // casting
//   cV3 := "500000"
//   cV4, err := strconv.Atoi(cV3) // Atoi - Ascii to int
//   pl(cV4, err, reflect.TypeOf(cV4)) 
// }


// casting function
// func main() {
//   // casting
//   cV1 := 1.5
//   cV2 := int(cV1)
//   pl(cV2) 

//   // 1
// }

// data types

/*
int, float64, bool, string, rune

what is rune?
rune is alias for int32
*/

// type of implemenation
// func main() {
//   pl(reflect.TypeOf(1)) 
//   pl(reflect.TypeOf(2.15))
//   pl(reflect.TypeOf(true)) 
//   pl(reflect.TypeOf("true"))
//   pl(reflect.TypeOf("🐒")) 

//   // int
//   // float64
//   // bool
//   // string
//   // string
// }
