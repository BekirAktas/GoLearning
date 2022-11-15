package models

type Person struct {
	id int;
	name string;
	age int;
}

type Identify interface {
	Id() string
}

func (p *Person) Id() int {
	return p.id;
}

func NewPerson(name string, age int ) Person {
	return Person{
		name: name,
		age: age,
	}
}

func (p *Person) GetName() string{
	return p.name;
}

func (p *Person) SetName(name string) {
	p.name = name;
}