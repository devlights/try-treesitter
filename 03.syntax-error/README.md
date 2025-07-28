# 構文エラー

[Parse](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Parser.Parse)メソッドの結果から取得できる[Node](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Node)は、自身のエラー状態を持っています。もし、構文エラーが発生している場合はエラーありとなっています。

ルートノードにて [HasError](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Node.HasError)を呼び出して、エラーが存在すると判定できたら、再帰でノードを下って処理します。

該当ノードから、[StartPosition](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Node.StartPosition)と[EndPosition](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Node.EndPosition)を使うと、何行目の何列目なのかが取得できるので、この情報を元のソースコードと合わせると該当箇所がコード上でも特定できます。

ただし、Tree-sitterはコンパイラでは無いので、エラーとして検出されるのはあくまでも「シンタックスエラー」と判定されるもののみとなります。

例として再帰関数は以下のように実装できますね。

```go
func walk(node *tree_sitter.Node, callback func(*tree_sitter.Node)) {
	callback(node)
	for i := uint(0); i < node.ChildCount(); i++ {
		walk(node.Child(i), callback)
	}
}
```

こんな感じで使えます。

```go
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
```

本サンプルを実行すると以下のような出力となります。

```sh
$ task
task: [default] go run .
エラー： 3行7列 - 3行7列
  問題のコード:         int x
エラー： 5行2列 - 5行4列
  問題のコード:         if {
エラー： 7行5列 - 7行5列
  問題のコード:         int  = 20;
エラー： 9行25列 - 9行25列
  問題のコード:         printf("hello %d\n", z)
```
