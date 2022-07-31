package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {
	///外部キャンセル操作を受け取ったらサーバーを終了できるようにする
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
	}
}

//mainをrun関数に分離
func run(ctx context.Context) error {
	//関数のListenAndServeではなメソッドを使うことでグレースフルシャットダウンできる
	//タイムアウト設定なども可能な為、こちらの方が定番
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
		}),
	}
	eg, ctx := errgroup.WithContext(ctx)
	// 別ゴルーチンでHTTPサーバーを起動する
	//通常のゴルーチンを使うとDoneチャネルで通知を受け取る必要がある為、errorgroupを使う
	//戻り値にエラーが含まれるゴルーチンの並行処理が簡単にできる
	//sync.WaitGroupでは別ゴルーチン上で実行する関数からエラーを受け取ることができない
	eg.Go(func() error {
		// http.ErrServerClosed は
		// http.Server.Shutdown() が正常に終了したことを示すので異常ではない。
		if err := s.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			log.Printf("failed to close: %+v", err)
			return err
		}
		return nil
	})

	//グレースフルシャットダウン
	//今回は外部から終了通知がこない。テストで外部からキャンセルしてみる
	// チャネルからの通知（終了通知）を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown: %+v", err)
	}
	// Goメソッドで起動した別ゴルーチンの終了を待つ。
	return eg.Wait()
}
