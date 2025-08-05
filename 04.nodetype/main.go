package main

import (
	"log"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

const (
	C_CODE = `#include <stdio.h>
int main(void) {
	int x = 10;
	printf("%d\n", x);
	return 0;
}`
)

var (
	code = []byte(C_CODE)
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	p := tree_sitter.NewParser()
	defer p.Close()

	_ = p.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))

	t := p.Parse(code, nil)
	r := t.RootNode()

	//
	// C言語のパーサーを使っているため
	//   https://github.com/tree-sitter/tree-sitter-c/blob/master/src/node-types.json
	// に定義されている要素が Node.Kind() の値として登場する
	//
	log.Printf("RootNode=%s", r.Kind())
	for i := uint(0); i < r.NamedChildCount(); i++ {
		log.Printf("\tChildNode=%s", r.NamedChild(i).Kind())
	}

	return nil
}
