# 解析処理

解析を実施するには [Parse](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Parser.Parse)メソッド を呼び出します。[Tree-sitter](https://github.com/tree-sitter/tree-sitter) は、インクリメンタルな構文解析に対応しているため
第二引数に旧ノードを指定することが出来るようになっています。初回の場合、または、一回きりの解析の場合は ```nil``` を渡します。戻り値は ```*tree_sitter.Tree``` です。解析に失敗した場合は ```nil``` が返ります。

[parser](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Parser)と同様に[tree](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Tree)もCloseが必要になります。

```go
p := tree_sitter.NewParser()
defer p.Close()

p.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language())

srcCode, _ := os.ReadFile("/path/to/src/file")
tree := p.Parse(srcCode, nil)
defer tree.Close()
```
