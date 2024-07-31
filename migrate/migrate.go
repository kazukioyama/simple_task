package main //このファイル(のmain関数)を、プログラムのエントリーポイント(一番最初に実行される部分)にさせるので（main関数を使うので）

import (
	"fmt"
	"go-rest-api-udemy/db"
	"go-rest-api-udemy/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")       // 遅延実行（main関数の終了時に実行させる）
	defer db.CloseDB(dbConn)                         // 遅延実行
	dbConn.AutoMigrate(&model.User{}, &model.Task{}) //引数にDBに反映させたいモデル構造を入れる（この場合は、model構造体からフィールド値を入れずインスタンス化させたもの）
}
