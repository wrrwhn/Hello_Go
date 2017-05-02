package utils

var fcVal string

//GOG
func HeeloFuncCall() {
	fcVal = "G"
	print(fcVal)
	f1()
}

func f1() {
	fcVal := "O"
	print(fcVal)
	f2()
}

func f2() {
	print(fcVal)
}
