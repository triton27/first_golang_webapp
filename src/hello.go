package main

// 必要なライブラリ
// fmt はフォーマット付きの I/O をサポートするパッケージ
// http は HTTP でやり取りするのに必要なパッケージ
import (
		"fmt"
		"net/http"
)

// ハンドラ関数
// インターフェース ResponseWriter と、構造体 Request へのポインタ
func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello Golang, %s!", request.URL.Path[1:])
}

func main() {
	// / が呼び出されたとき、上で定義したハンドラが起動されるようになる
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
