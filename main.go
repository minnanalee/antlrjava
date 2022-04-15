package main

import (
	"fmt"
	"github.com/minnanalee/antlrjava/antlr"
	"github.com/minnanalee/antlrjava/examples"
	"github.com/minnanalee/antlrjava/parser"
	"log"
)

func main() {
	fs, err := antlr.NewFileStream("java_module_test.java")
	if err != nil {
		log.Fatal(err)
	}
	lex := parser.NewJavaLexer(fs)
	tokens := antlr.NewCommonTokenStream(lex, antlr.TokenDefaultChannel)
	p := parser.NewJavaParser(tokens)
	visitor := examples.NewVisitor()
	tree := p.CompilationUnit()
	fmt.Println(tree.ToStringTree(nil, p))
	fmt.Println(visitor.Visit(tree))

}
