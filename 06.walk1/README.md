# ノードのトラバース（１）

tree_sitterには [TreeCursor](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#TreeCursor) というオブジェクトが存在します。

このオブジェクトは [Tree.Walk](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Tree.Walk) や [Node.Walk](https://pkg.go.dev/github.com/tree-sitter/go-tree-sitter#Node.Walk) から得られます。

このオブジェクトを利用すると配下のノードを再帰的にトラバースすることが可能なのですが、若干利用方法が難しいです。

本サンプルでは、普通にループとコールバックを用いたノードのトラバースを行っています。
