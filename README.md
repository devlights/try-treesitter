# try-treesitter

This is my TUTORIAL project for tree-sitter.

## tree-sitterとは

[tree-sitter-doc](https://tree-sitter.github.io/tree-sitter/)にライブラリの紹介が書いてあります。

> Tree-sitter is a parser generator tool and an incremental parsing library. It can build a concrete syntax tree for a source file and efficiently update the syntax tree as the source file is edited. Tree-sitter aims to be:

> (Tree-sitterは構文解析ツールであり、インクリメンタル構文解析ライブラリである。ソースファイルの具体的な構文ツリーを構築し、ソースファイルの編集に応じて構文ツリーを効率的に更新することができます。Tree-sitterは次のようなことを目指しています：)

> General enough to parse any programming language (あらゆるプログラミング言語を解析できる汎用性)

> Fast enough to parse on every keystroke in a text editor (テキストエディタでのすべてのキーストロークを解析するのに十分な速度)

> Robust enough to provide useful results even in the presence of syntax errors (構文エラーがあっても有用な結果を提供できるほど頑健である。)

> Dependency-free so that the runtime library (which is written in pure C11) can be embedded in any application (依存性がないため、ランタイムライブラリ（純粋なC11で書かれている）をあらゆるアプリケーションに組み込むことができる。)

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

