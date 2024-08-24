package funcopts

// import (
// 	_ "github.com/ecordell/optgen/helpers"
// )

//go:generate optgen -output gen.go . Redis
type Redis struct {
	Addr     string `debugmap:"visible"`
	Password string `debugmap:"sensitive"`
	UserName string `debugmap:"visible"`
}
