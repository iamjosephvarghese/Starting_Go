package main


import "fmt"


func main(){


  var number int = 7

  var floating float64 = 1.723233

  value := 1
  //once declared as a number cant give string value later on in the program

  fmt.Println(number,floating,value)
  //automatically gives spacing


  const pi float64 = 3.141593
  //declaring a constant value pi
  fmt.Printf("The value of constant pi is %f\n",pi);


  var(
    one = 1
    two = 2
  )

  fmt.Printf("one is: %d\n",one);
  fmt.Printf("two is: %d\n",two);


  var name string = "Joseph"
  fmt.Println(len(name))
  //finding string length

  //printing integer value
  fmt.Printf("%d\n",700);

  //printing hexadecimal value
  fmt.Printf("%x\n",15);

  //printing binary value of a number
  fmt.Printf("%b\n",100);

  //printing exponential value
  fmt.Printf("%e\n",pi);

  //printing the ASCII charter
  fmt.Printf("%c\n",65);


}
