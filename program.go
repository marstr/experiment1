package main

import (
	earlier "github.com/marstr/experiment2/print"
	later "github.com/marstr/experiment3/print"
	fromSubtree "github.com/marstr/rolledup/experiment2/print"
)

func main() {
	foo := earlier.StandardOutput{}
	bar := later.StandardOutput{}
	baz := fromSubtree.StandardOutput{}

	foo.Print("Early Import Call")
	bar.Print("Later Import Call", 5)
	baz.Print("Subtree Call")
}
