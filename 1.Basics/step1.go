package main

func main() { // STRING, NUMBER, BOOLEAN

	//STRING

	//o'zg - nomi -  VALUE  - qiymati
	var message string = "Hello Student"
	println(message)

	var messagel string     //DECLARE --> Tanitdin
	messagel = "Hello Word" // ASSIGNMED --> Tenglashtirdim
	println(messagel)

	var boxl, box2 string = "first", "second"
	print(boxl, box2)

	// NOTE

	var box = "THis is variable" // TIPI berilmagan ammo bu ham ishlaydi !
	print(box)                   // tipi berilmasa o'zi avtomat topib sekin ishlaydi
	// agar tipi berilsa kod tezlik jihatidan yaxshi ishlaydi
	//_____________________________________________________________________________________

	// NUMBER
	// INT
	var a, b, c int  //BUTUN SON --Ex: -7,--5,-11,-23,--94,-768,-811
	println(a, b, c) // --> 0 0 0

	var n1, n2, n3 int = 11, 22, 33
	println(n1, n2, n3)

	var num1, num2, num3 int = 11, 22, 33
	println(num1 + num2 + num3)
	//
	//_____________________________________________________________________________________

	// FLOAT

	var x, y, z float32 //Qoliq son Ex: 1.5, 11.6, 123.5,4567.11, 12.5
	println(x, y, z)    //-->+0.000000e+000 +0.000000+000 +0.000000+000

	var f1, f2, f3 float64 = 11.1, 22.2, 33.3
	println(f1, f2, f3)

	// Boolean

	var erkakM1 bool = true //false
	println(erkakM1)

}






















