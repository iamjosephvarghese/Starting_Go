package main


import "fmt"


func main(){


  var number int = 7

  var floating float64 = 1.723233

  value := 1
  //once declared as a number cant give string value later on in the program

  fmt.Println(number,floating,value)
  //automatically gives spacing


  const pi float64 = 3.14
  //declaring a constant value pi
  fmt.Println("The value of constant pi is",pi);


  var(
    one = 1
    two = 2
  )

  fmt.Println("one is:",one);
  fmt.Println("two is:",two);


  var name string = "Joseph"
  fmt.Println(len(name))
  //finding string length



}
