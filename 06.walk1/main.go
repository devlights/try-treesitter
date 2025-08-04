package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

const (
	C_CODE = `#include <stdio.h>
#include <stdlib.h>

int add(int x, int y) {
	// ADD
	return x+y;
}

int main(void) {
	int z = add(1, 2);
	if (z < 10) {
	    printf("%d\n", z);
	}

	for (int i = 0; i < z; i++) {
		switch (i) {
		case 0:
			printf("first\n");
			return;
		default:
			printf("other\n");
			return;
		}
	}

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
		tree = p.Parse([]byte(C_CODE), nil)
	)
	if tree == nil {
		return errors.New("構文解析に失敗: Parse()")
	}

	var (
		root     = tree.RootNode()
		callback = func(node *tree_sitter.Node, depth int) {
			var (
				k = node.Kind()
			)
			fmt.Printf("%s%-25s\n", strings.Repeat(" ", depth*2), k)
		}
	)
	walk(root, 0, callback)

	return nil
}

// walk は、TreeCursorを利用せずに指定したノード配下をトラバースする関数です。
func walk(node *tree_sitter.Node, depth int, callback func(*tree_sitter.Node, int)) {
	callback(node, depth)

	for i := uint(0); i < node.NamedChildCount(); i++ {
		walk(node.NamedChild(i), depth+1, callback)
	}
}
