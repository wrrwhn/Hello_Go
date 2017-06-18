package utils

import (
	"fmt"
	"time"
)

// 全名	utils.Person
type Person struct {
	name   string
	height int
	weight int
}

type Address struct {
	addr string
}

type Card struct {
	name  string
	addr  *Address
	img   string
	birth time.Time
}

func HelloType() {

	// Person
	// person.yao
	var yao *Person = new(Person)
	yao.name = "yqj"
	yao.height = 180
	yao.weight = 75
	fmt.Println(yao)
	// person.lu
	var lu Person = Person{name: "lmm", height: 160, weight: 45}
	fmt.Println(lu)
	// person.chen
	var chen Person = Person{"chen", 175, 65}
	fmt.Println(chen)

	// Card
	addr := Address{"china"}
	card := Card{name: "yao", addr: &addr, birth: time.Now()}
	fmt.Printf("{name= %v, img= %v, addr= %v, birth= %v}", card.name, card.img, (*card.addr), card.birth)
}
