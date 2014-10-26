package dsl

import (
	"fmt"
	"log"
	"os"
	//"strings"
	//"strconv"
	//"strings"
)

type MyDsl struct {
	ast *Node
}

func Create() (d *MyDsl) {
	return &MyDsl{ast: nil}
}

func (d *MyDsl) InitAST(root *Node) {
	if d.ast == nil {
		d.ast = root
	}
}

// Walk traverses the AST in depth-first order
func (d *MyDsl) Walk(n *Node) {
	if n == nil {
		return
	}
	d.Walk(n.Left)
	fmt.Println(n.Type)
	d.Walk(n.Right)
}

func (d *MyDsl) Init(filename string) {
	file, err := os.Open("test.txt") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	yyParse(NewLexerWithInit(file, func(y *Lexer) { y.p = d }))

	d.Walk(d.ast)
}
