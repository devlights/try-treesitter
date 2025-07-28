package main

import (
	"bufio"
	"bytes"
	"errors"
	"log"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

const (
	C_CODE = `#include <stdio.h>
int main(void) {
	int x
	int y = 10;
	if {
	}
	int  = 20;

	printf("hello %d\n", z)
	return 0;
}`
)

var (
	srcCode = []byte(C_CODE)
)

func main() {
	log.SetFlags(0)

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	// パーサーの生成と言語の設定
	var (
		p   = tree_sitter.NewParser()
		l   = tree_sitter.NewLanguage(tree_sitter_c.Language())
		err error
	)
	defer p.Close()
	if err = p.SetLanguage(l); err != nil {
		return err
	}

	// 解析を実施しツリー取得
	var (
		tree = p.Parse(srcCode, nil)
	)
	if tree == nil {
		return errors.New("解析に失敗: Parse()")
	}
	defer tree.Close()

	// エラーが存在するか検査
	var (
		root = tree.RootNode()
	)
	if root.HasError() {
		walk(root, func(n *tree_sitter.Node) {
			if n.IsError() || n.IsMissing() {
				var (
					start = n.StartPosition()
					end   = n.EndPosition()
				)
				log.Printf("エラー： %d行%d列 - %d行%d列", start.Row+1, start.Column+1, end.Row+1, end.Column+1)

				// コード行表示
				var (
					scanner = bufio.NewScanner(bytes.NewReader(srcCode))

					line    string
					numLine int
				)
				for ; scanner.Scan(); numLine++ {
					if numLine == int(start.Row) {
						line = scanner.Text()
						log.Printf("  問題のコード: %s\n", line)
						break
					}
				}
			}
		})
	}

	return nil
}

func walk(node *tree_sitter.Node, callback func(*tree_sitter.Node)) {
	callback(node)
	for i := uint(0); i < node.ChildCount(); i++ {
		walk(node.Child(i), callback)
	}
}
