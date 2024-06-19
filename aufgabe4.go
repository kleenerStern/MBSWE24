package main

import "strconv"

type node[T any] struct {
	val  T
	next *node[T]
}

type Show interface {
	show() string
}

type boolean struct {
	val bool
}

type integer struct {
	val int
}

// Variant with method overloading and interfaces
func (this boolean) show() string {
	if this.val {
		return "true"
	} else {
		return "false"
	}
}

func (this integer) show() string {
	return strconv.Itoa(this.val)
}

func showNode[T Show](n *node[T]) string {
	var s string

	for n != nil {
		s = s + n.val.show() + " -> "
		n = n.next
	}

	s = s + " nil"

	return s
}

// Variant with method lookup and without interfaces / method overloading
func show_integer(i integer) string {
	return strconv.Itoa(i.val)
}

func show_boolean(b boolean) string {
	if b.val {
		return "true"
	} else {
		return "false"
	}
}

func show_lookup(i interface{}) string {
	var y string

	switch v := i.(type) {
	case integer:
		y = show_integer(v)
	case boolean:
		y = show_boolean(v)
	}
	return y
}

func showNodeLookup[T interface{}](n *node[T]) string {
	var s string

	for n != nil {
		s = s + show_lookup(n.val) + " -> "
		n = n.next
	}

	s = s + " nil"

	return s
}

// Variant with Dictionary translation and without interfaces / method overloading
type show_value struct {
	val  interface{}
	show func(interface{}) string
}

// wrapper functions
func show_boolean_wrapper(i interface{}) string {
	return show_boolean(i.(boolean))
}

func show_integer_wrapper(i interface{}) string {
	return show_integer(i.(integer))
}

func showNodeDict(n *node[show_value]) string {
	var s string

	for n != nil {
		s = s + n.val.show(n.val.val) + " -> "
		n = n.next
	}

	s = s + " nil"

	return s
}
