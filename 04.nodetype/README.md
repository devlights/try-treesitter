# tree-sitterで出てくるノードタイプ(Kind)の調べ方

tree-sitterで処理を書いていると自ずとノードのタイプ毎に処理を分岐することになります。

このノードタイプは、[Node.Kind](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Node.Kind)の値から取得出来ます。

例えば、C言語のパーサーを利用した場合は以下のようなノードタイプが現れます。

- comment
- function_declaration
- identifier

これらのノードタイプは、言語ごとに異なっており、それぞれの言語用のリポジトリにある ```nodetypes.json``` ファイルに定義されています。

なので、このファイルを見てノートタイプ毎の分岐を書いて処理するようにします。

各言語用にリポジトリは以下のようになっており、リポジトリ名が ```tree-sitter-言語``` となっています。

（逆にバインディングライブラリの方は ```言語-tree-sitter``` というネーミングになっています。）

一例として

- [C](https://github.com/tree-sitter/tree-sitter-c/blob/master/src/node-types.json)
- [Go](https://github.com/tree-sitter/tree-sitter-go/blob/master/src/node-types.json)
- [Python](https://github.com/tree-sitter/tree-sitter-python/blob/master/src/node-types.json)
- [Java](https://github.com/tree-sitter/tree-sitter-java/blob/master/src/node-types.json)
- [C#](https://github.com/tree-sitter/tree-sitter-c-sharp/blob/master/src/node-types.json)

のようになります。フルリストは [Parses](https://tree-sitter.github.io/tree-sitter/#parsers) にあります。
