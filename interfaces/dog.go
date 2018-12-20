package main

type Dog struct {
	name string
}

func (d *Dog) Name() string {
	return d.name
}

func (d *Dog) SetName(value string) {
	d.name = value
}

func (d *Dog) Bay() string {
	return d.name + " say: gua! gua!"
}
