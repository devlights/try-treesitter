package main

import (
	"log"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	//
	// 最初の手順は
	//   1. パーサーを生成
	//   2. 言語を設定
	// となる。
	//
	// パーサーはCloseする必要がある。
	//
	var (
		p   *tree_sitter.Parser   = tree_sitter.NewParser()
		l   *tree_sitter.Language = tree_sitter.NewLanguage(tree_sitter_c.Language())
		err error
	)
	defer p.Close()

	if err = p.SetLanguage(l); err != nil {
		return err
	}

	return nil
}
