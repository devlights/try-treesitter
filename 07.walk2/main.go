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
	defer tree.Close()

	var (
		root     = tree.RootNode()
		cursor   = root.Walk()
		callback = func(node *tree_sitter.Node, depth uint32) {
			var (
				k = node.Kind()
				d = int(depth) * 2
			)
			fmt.Printf("%s%-25s\n", strings.Repeat(" ", d), k)
		}
	)
	defer cursor.Close() // *TreeCursorは閉じる必要がある

	walk(cursor, callback)

	return nil
}

// walk は、TreeCursorを利用して指定したノード配下をトラバースする関数です。
func walk(cursor *tree_sitter.TreeCursor, callback func(*tree_sitter.Node, uint32)) {
	var (
		node  = cursor.Node()  // 現在のノード
		depth = cursor.Depth() // 現在の深さ
	)
	// *TreeCursorでは、Namedだけではなく全てのノードが回るのでNamedだけ対象にするには絞り込む必要がある
	if node.IsNamed() {
		callback(node, depth)
	}

	// 子ノードが存在する場合は移動
	if cursor.GotoFirstChild() {
		for {
			walk(cursor, callback)

			// 兄弟ノードが存在すれば移動
			if !cursor.GotoNextSibling() {
				break
			}
		}

		// 親ノードに戻る
		cursor.GotoParent()
	}
}
