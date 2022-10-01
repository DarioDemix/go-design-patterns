package decorator

import "fmt"

type Decorator struct{}

func (d *Decorator) Demo() {
	fmt.Println("Decorator demo:")

	dragon := newDragon()
	dragon.SetAge(10)
	dragon.Fly()
	dragon.Crawl()
}
