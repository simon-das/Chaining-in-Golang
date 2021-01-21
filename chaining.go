package main

import "fmt"

type workers struct {
}

func (w workers) add() workers {
	fmt.Println("Added")
	return w
}

func (w workers) remove() workers {
	fmt.Println("Removed")
	return w
}

func main() {
	w := workers{}
	w.add().remove().add()
}
