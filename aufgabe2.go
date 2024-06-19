package main

type association struct {
	members int
	staff   int
}

type company struct {
	employees int
}

// with method overloading
func (a association) persons() int {
	return a.members + a.staff
}

func (c company) persons() int {
	return c.employees
}

type organization interface {
	persons() int
}

func sumPersons(x, y organization) int {
	return x.persons() + y.persons()
}

// Run-time method lookup
func persons_association(a association) int {
	return a.members + a.staff
}

func persons_company(c company) int {
	return c.employees
}

func persons_lookup(x interface{}) int {
	var y int

	switch v := x.(type) {
	case association:
		y = persons_association(v)
	case company:
		y = persons_company(v)
	}
	return y
}

func sumPersons_Lookup(x, y interface{}) int {
	return persons_lookup(x) + persons_lookup(y)
}

// Dictionary translation
type organization_Value struct {
	val     interface{}
	persons func(interface{}) int
}

func sumPersons_Dict(x, y organization_Value) int {
	return x.persons(x.val) + y.persons(y.val)
}

// wrapper functions
func persons_association_wrapper(v interface{}) int {
	return persons_association(v.(association))
}
func persons_company_wrapper(v interface{}) int {
	return persons_company(v.(company))
}
