package main

import (
	hanlder "github.com/sebarray/backendproyect/handler"
	"github.com/sebarray/tooldevgo/config"
)

func main() {
	config.Loadenv()
	hanlder.Manager()
}
