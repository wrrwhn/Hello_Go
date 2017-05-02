package utils

var gsVal = "G"

func HelloGlobgsVallScope() {
	nn()
	mm()
	nn()
}

func nn() {
	print(gsVal)
}

func mm() {
	gsVal = "O"
	print(gsVal)
}
