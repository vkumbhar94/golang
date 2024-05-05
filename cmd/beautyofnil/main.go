package main

import "fmt"

type SayHello struct{}

func (s *SayHello) Hello() {
	fmt.Println("Hello")
}
func main() {
	fmt.Println("Start")
	var s *SayHello
	fmt.Println(s)
	// why does this not panic? It should panic because s is nil
	s.Hello()
	fmt.Println("End")

	usecase1SoftDependency()
	usecase2LazyInitialisation()
}

type Config struct {
	settings map[string]any
}

func (c *Config) Get() map[string]any {
	if c == nil {
		// this will get called again and again as the following assignment doesn't
		// reflect to the outside object of the caller.
		// .
		// However, a direct approach like the following won't work
		// as expected due to Go's receiver semantics:goCopy code
		fmt.Println("initialising config")
		// this is repetitive code, but only gain is that it doesn't panic
		c = &Config{settings: map[string]any{
			"key": "value",
		}}
	}
	return c.settings
}

func usecase2LazyInitialisation() {
	var c *Config
	fmt.Println(c.Get())
	fmt.Println(c.Get())
}

type A struct {
	Name string
}

type B struct {
	al []A
}

func (b *B) GetAList() []A {
	// if you handle nil check within the method, it will not panic
	// it doesn't enforce the objects to initialise this when embedding
	if b == nil {
		return nil
	}
	return b.al
}

type C struct {
	b *B
}

func usecase1SoftDependency() {
	var c = C{}
	fmt.Println(c.b.GetAList())
}
