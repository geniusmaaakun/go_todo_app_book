//go:build tools

//go run　で最新のバージョンのプログラムが実行されることを防ぐ
//tools.goには該当ツールをimportしたtools.goを定義することでgo.mod　によるバージョン管理ができる
//このファイルはビルドタグを指定しない場合は無視される

package main

//下記のパッケージを利用してモックをgo generateで生成する
import _ "github.com/matryer/moq"
