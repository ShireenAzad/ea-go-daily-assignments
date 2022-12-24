package main

import (
	"errors"
	"math"
	"fmt"
)
var divisionByZero=errors.New("division By Zero")

func main(){
	var number int;
	loop :
	for(true){
		fmt.Println("Enter the number for operation \n 1.Addition \t 2.Subtraction \t 3.MultiPlication \t 4.Division \t 5.Sin \t 6.Cos \t 7.Tan \t 8.Square Root \t 9.exit" );
			fmt.Scanln(&number);
		 switch number {
		 case 1:
			var number1,number2 float64;
			fmt.Println("Enter number 1");
			fmt.Scanln(&number1);
			fmt.Println("Enter number 2");
			fmt.Scanln(&number2);
			fmt.Println(add(number1,number2));
			break;
		 case 2:
			var number1,number2 float64;
			fmt.Println("Enter number 1");
			fmt.Scanln(&number1);
			fmt.Println("Enter number 2");
			fmt.Scanln(&number2);
		fmt.Println(subtract(number1,number2));
		break;
		 case 3:
			var number1,number2 float64;
			fmt.Println("Enter number 1");
			fmt.Scanln(&number1);
			fmt.Println("Enter number 2");
			fmt.Scanln(&number2);
		 fmt.Println(multiply(number1,number2));
		 break;
		 case 4:
			var number1,number2 float64;
			fmt.Println("Enter number 1");
			fmt.Scanln(&number1);
			fmt.Println("Enter number 2");
			fmt.Scanln(&number2);
		  fmt.Println(divide(number1,number2));
		  break;
		  case 5:
			var number float64;
			fmt.Println("Enter number");
			fmt.Scanln(&number);
		  fmt.Println(sin(number));
		  break;
		  case 6:
			var number float64;
			fmt.Println("Enter number");
			fmt.Scanln(&number);
		  fmt.Println(cos(number));
		  break;
		  case 7:
			var number float64;
			fmt.Println("Enter number");
			fmt.Scanln(&number);
		  fmt.Println(tan(number));
		  break;
		 case 8:
			var number float64;
			fmt.Println("Enter number");
			fmt.Scanln(&number);
		  fmt.Println(squareRoot(number));
		  break;
		  case 9:
		  break loop;
		 }

	}
}
func add(x,y float64)(float64){
	return x+y;
}
func subtract(x,y float64)(float64){
	return x-y;
}
func multiply(x,y float64)(float64){
	return x*y;
}
func divide(x,y float64 )(float64,error){
if (y == 0){
return 0,divisionByZero;
}
return x/y,nil;
}
func sin(x float64)(float64){
	 return math.Sin(x);
}
func cos(x float64)(float64){
	 return math.Cos(x);
}
func tan(x float64)(float64){
	 return math.Tan(x);
}
func squareRoot(x float64)(float64){
	if(x<0 || math.IsNaN(x)){
		return math.NaN();
	}
	if(x==0){
		return 0;
	}
	return math.Sqrt(x);
}