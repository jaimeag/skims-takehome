package main

import (
	modules "github.com/jaimeag/skims-takehome/modules"
)

func main() {
	topModule := modules.NewTopModule()

	topModule.GetTopFivePokemonInfo()
}
