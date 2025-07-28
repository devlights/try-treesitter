package main

import (
	"errors"
	"log"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

const (
	C_CODE = `#include <stdio.h>
#include <stdlib.h>

int main(void) {
	char *msg = "helloworld";
	printf("%s\n", msg);

	return EXIT_SUCCESS;
}
	`
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	var (
		p   = tree_sitter.NewParser()
		l   = tree_sitter.NewLanguage(tree_sitter_c.Language())
		err error
	)
	defer p.Close()
	if err = p.SetLanguage(l); err != nil {
		return err
	}

	var (
		tree *tree_sitter.Tree = p.Parse([]byte(C_CODE), nil)
		root *tree_sitter.Node
	)
	if tree == nil {
		return errors.New("構文解析に失敗: Parse()")
	}

	root = tree.RootNode()
	log.Println(root.ToSexp()) // S式で構文木を出力
	log.Printf("\n\nKind: %s, Start: %v, End: %v", root.Kind(), root.StartPosition(), root.EndPosition())

	return nil
}
