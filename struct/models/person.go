package models

type person struct {
	id int
	name string
}

func NewPerson(id int, name string) person {
	return person{
		id: id,
		name: name,
	}
}

func (p *person)SetNickName(name string) {
	p.name
}