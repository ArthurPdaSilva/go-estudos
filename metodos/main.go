package main

import "fmt"

type user struct {
	name string
	age  uint8
}

//Todos os users agora tem o método salvar agora
func (u user) save() {
	fmt.Printf("Salvando o %s no banco de dados\n", u.name)
}

func (u user) maiorDeIdade() bool {
	return u.age >= 18
}

//Alterar uma props de um struct, se usa ponteiro
func (u *user) fazerAniversario() {
	u.age++
}

func main() {
	u := user{"Eduardo", 20}
	fmt.Println(u)
	u.save()
	fmt.Println(u.maiorDeIdade())
	u.fazerAniversario()
	fmt.Println(u.age)

}
