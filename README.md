# try-treesitter

This is my TUTORIAL project for tree-sitter.

## tree-sitterを使ったソースファイルの解析

[tree-sitter](https://github.com/tree-sitter/tree-sitter)は、インクリメンタル（増分）な構文解析が可能なパーサジェネレータ・ライブラリです。
[tree-sitter-doc](https://tree-sitter.github.io/tree-sitter/)に公式ドキュメントがあります。

いろいろな言語の構文にデフォルトで対応しており、独自の構文解析を構築することも出来ます。

また、様々なプログラム言語に対してのバインディングも存在します。（Go,Python,C#,Javaのバインディングは公式対応）

## tree-sitterの用意

例題としてC言語のソースファイルを解析する場合は以下のようにバインディングライブラリを選択します。

```go
$ go get github.com/tree-sitter/go-tree-sitter@latest
$ go get github.com/tree-sitter/tree-sitter-c/bindings/go@latest
```

Goのバインディングは、これで完了です。

## tree-sitter-cliの用意

```sh
$ cargo install tree-sitter-cli
```

### CLIのためにC言語バインディングを配置

```sh
$ cd
$ tree-sitter init-config
$ mkdir -p ~/tree-sitter-parsers
$ cd ~/tree-sitter-parsers
$ git clone https://github.com/tree-sitter/tree-sitter-c.git
```

```~/.config/tree-sitter/config.json```のparser-directoriesを以下のように設定する。

```json
  "parser-directories": [
    "/home/dev/tree-sitter-parsers"
  ],
```

これでCLI側で構文木が出力できるようになる。デバッグ時などに便利。

```sh
$ tree-sitter parse app.c
```

