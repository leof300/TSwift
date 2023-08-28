package main

import (
	"TSwiftCompiler/ast"
	"TSwiftCompiler/interpreter"
	TSVisitor "TSwiftCompiler/visitor"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"log"
)

func main() {
	fs, err := antlr.NewFileStream("./grammar/input2.txt")
	if err != nil {
		log.Fatal(err)
	}
	lexer := TSVisitor.NewTSParser_rulesLexer(fs)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := TSVisitor.NewTSParser_rulesParser(tokens)
	tree := parser.Start_()

	visitor := interpreter.NewTSwiftVisitor()

	root := visitor.Visit(tree).(ast.TSExpressioner)
	//root := visitor.Visit(tree)
	a := "ds"
	fmt.Print("Ejemplo" + a)
	fmt.Print(root)
}

/*
type Creature struct {
	Name string

	Real bool
}

func Dump(c *Creature) {

	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)

}

func (c Creature) Dump() {

	fmt.Printf("Name: '%s', Real: %t\n", c.Name, c.Real)

}

type FlyingCreature struct {
	Creature

	WingSpan int
}

func (fc FlyingCreature) Dump() {

	fmt.Printf("Name: '%s', Real: %t, WingSpan: %d\n",

		fc.Name,

		fc.Real,

		fc.WingSpan)

}

type Unicorn struct {
	Creature
}

type Dragon struct {
	FlyingCreature
}

type Pterodactyl struct {
	FlyingCreature
}

func NewPterodactyl(wingSpan int) *Pterodactyl {

	pet := &Pterodactyl{

		FlyingCreature{

			Creature{

				"Pterodactyl",

				true,
			},

			wingSpan,
		},
	}

	return pet

}

type Dumper interface {
	Dump()
}

type Door struct {
	Thickness int

	Color string
}

func (d Door) Dump() {

	fmt.Printf("Door => Thickness: %d, Color: %s", d.Thickness, d.Color)

}

func main() {

	creature := &Creature{

		"some creature",

		false,
	}

	uni := Unicorn{

		Creature{

			"Unicorn",

			false,
		},
	}

	pet1 := &Pterodactyl{

		FlyingCreature{

			Creature{

				"Pterodactyl",

				true,
			},

			5,
		},
	}

	pet2 := NewPterodactyl(8)

	door := &Door{3, "red"}

	Dump(creature)

	creature.Dump()

	uni.Dump()

	pet1.Dump()

	pet2.Dump()

	creatures := []Creature{

		*creature,

		uni.Creature,

		pet1.Creature,

		pet2.Creature}

	fmt.Println("Dump() through Creature embedded type")

	for _, creature := range creatures {

		creature.Dump()

	}

	dumpers := []Dumper{creature, uni, pet1, pet2, door}

	fmt.Println("Dump() through Dumper interface")

	for _, dumper := range dumpers {

		dumper.Dump()

	}

}
*/
