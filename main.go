package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// データベースに接続する
func connectDB() (*sql.DB, error) {
	// .envファイルを読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から取得
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	database := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	// MySQLに接続するためのDSN(Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, database)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// 接続確認
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("データベースに接続しました。")
	return db, nil
}

// トランザクションを使ってデータを挿入する
func insertUserWithTransaction(db *sql.DB, name, email string, age int) error {
	// トランザクションの開始
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// 挿入クエリ
	query := "INSERT INTO users (name, email, age) VALUES (?, ?, ?)"

	// 挿入処理
	_, err = tx.Exec(query, name, email, age)
	if err != nil {
		// エラーが発生した場合、ロールバック
		tx.Rollback()
		return err
	}

	// トランザクションのコミット
	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("トランザクションでユーザーを追加しました。", name)
	return nil
}

// ユーザーを取得する
func getUsers(db *sql.DB) ([]string, error) {
	// クエリ
	query := "SELECT name FROM users"

	// クエリ実行
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 結果を処理
	var users []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		users = append(users, name)
	}

	return users, nil
}

func main() {
	// データベースに接続
	db, err := connectDB()
	if err != nil {
		log.Fatal("データベース接続に失敗しました：", err)
	}
	defer db.Close()

	// トランザクションでユーザーを追加
	err = insertUserWithTransaction(db, "Taro Yamada", "taro@example.com", 30)
	if err != nil {
		log.Fatal("ユーザー追加に失敗しました：", err)
	}

	// ユーザー一覧を取得
	users, err := getUsers(db)
	if err != nil {
		log.Fatal("ユーザー取得に失敗しました：", err)
	}

	fmt.Println("ユーザー一覧", users)
}
