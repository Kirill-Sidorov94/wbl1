package main

type Human struct {
	Name string
}

type Action struct {
	Human
}

func (h *Human) GetName() string {
	return h.Name
}