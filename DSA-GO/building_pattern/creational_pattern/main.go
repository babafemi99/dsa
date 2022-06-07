package main

import "fmt"

type person struct {
	name, age             string //personal profile
	City, street, houseNo string // address profile
	job, salary, title    string // employment profile
}

type personBuilder struct {
	person *person
}

func NewPersonBuilder() *personBuilder {
	return &personBuilder{person: &person{}}
}
func (p *personBuilder) Build() *person {
	return p.person
}

//// employment building

type employmentBuilder struct {
	personBuilder
}

// address building
type addressBuilder struct {
	personBuilder
}

func (p *personBuilder) Work() *employmentBuilder {
	return &employmentBuilder{*p}
}

func (p *personBuilder) Lives() *addressBuilder {
	return &addressBuilder{*p}
}

// setting Employment fields

func (e *employmentBuilder) Title(title string) *employmentBuilder {
	e.person.title = title
	return e
}

func (e *employmentBuilder) Salary(salary string) *employmentBuilder {
	e.person.salary = salary
	return e
}
func (e *employmentBuilder) Job(job string) *employmentBuilder {
	e.person.job = job
	return e
}

// setting Address Fields

func (e *addressBuilder) City(city string) *addressBuilder {
	e.person.City = city
	return e
}

func (e *addressBuilder) Street(street string) *addressBuilder {
	e.person.street = street
	return e
}
func (e *addressBuilder) HouseNo(number string) *addressBuilder {
	e.person.houseNo = number
	return e
}

func main() {
	femi := NewPersonBuilder()
	femi.Lives().
		City("Yaba").Street("Eletu-Odibo").HouseNo("25").
		Work().
		Job("Software Engineer").Salary("120,000").Title("SWE 1")
	p := femi.Build()
	fmt.Printf("%v", p)
}
