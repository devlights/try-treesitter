# パーサーの生成と言語の設定

以下、[Using Parsers](https://tree-sitter.github.io/tree-sitter/using-parsers/index.html)から引用。

> Tree-sitter's parsing functionality is implemented through its C API, with all functions documented in the tree_sitter/api.h header file, but if you're working in another language, you can use one of the following bindings found here, each providing idiomatic access to Tree-sitter's functionality. Of these bindings, the official ones have their own API docs hosted online at the following pages:

> (Tree-sitterのパース機能はC APIで実装されており、すべての関数はtree_sitter/api.hヘッダファイルに記述されていますが、他の言語で作業している場合は、Tree-sitterの機能への慣用的なアクセスを提供する以下のバインディングのいずれかを使用することができます。これらのバインディングのうち、公式のものは以下のページに独自のAPIドキュメントがあります：)

公式バインディングには以下が存在する。(2025-07 現在)

- [C#](https://github.com/tree-sitter/csharp-tree-sitter)
- [Go](https://github.com/tree-sitter/go-tree-sitter)
- Haskell
- Java (JDK 22+)
- JavaScript (Node.js)
- JavaScript (Wasm)
- Kotlin
- [Python](https://github.com/tree-sitter/py-tree-sitter)
- Rust
- Swift
- Zig

どのバインディングを利用していても、使い方は大差ないため一つ覚えれば他の言語バインディングでも通用する。

手順としては

1.parserの生成(NewParser)
2.言語の設定(SetLanguage)

となる。

```go
p := tree_sitter.NewParser()
defer p.Close()

p.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))
```
