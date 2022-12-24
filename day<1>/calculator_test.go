package main

import (
	"math"
	"testing"
)
func TestSuccessfulAddition(t *testing.T) {
	got :=add(10,5);
	wanted:=float64(15);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
	}

}
func TestSuccessfulSubtraction(t *testing.T) {
	got :=subtract(10,5);
	wanted:=float64(5);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
	}

}
func TestSuccessfulMultiplication(t *testing.T) {
	got :=multiply(10,5);
	wanted:=float64(50);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
	}

}
func TestSuccessfulDivision(t *testing.T){
    got , _:=divide(10,5);
	wanted:=float64(2);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
	}
}
func TestFailureDivisionByZero(t *testing.T){
	_,result:=divide(10,0);
	if(result!=divisionByZero){
		t.Errorf("Denominator is not equal to zero and division is possible")
	}

}
func TestSuccessfulSin(t *testing.T){
	got :=sin(90);
	wanted:=float64( 0.893996663600558);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
}
}
func TestSuccessfulCos(t *testing.T){
	got :=cos(90);
	wanted:=float64(-0.4480736161291701);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
}
}
func TestSuccessfulTan(t *testing.T){
	got :=tan(90);
	wanted:=float64 (-1.995200412208242);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
}
}
func TestSqrtWhenNumberGreaterThanZero(t *testing.T){
	got :=squareRoot(25);
	wanted:=float64 (5);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
}
}
func TestSqrtWhenNumberLessThanZero(t *testing.T){
	got :=squareRoot(-1);
	if(!math.IsNaN(got)){
	t.Error("Number is greater than zero");
}
}
func TestSqrtWhenNumberEqualToZero(t *testing.T){
	got :=squareRoot(0);
	wanted:=float64 (0);
	if(got!=wanted){
	t.Errorf("got %g but wanted %g",got,wanted);
}
}
func TestSqrtWhenNUmberIsNaN(t *testing.T){
	got :=squareRoot(math.NaN());
	if(!math.IsNaN(got)){
	t.Error("Number is greater than zero");
}
}