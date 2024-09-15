# 何をしているものか
- Go言語とトランザクションを用いた、MySQLのハンズオン

## 手順
- 必要なパッケージのインストール
```sh
go mod init mysql-hands-on
go get -u github.com/go-sql-driver/mysql
```
- MySQLのセットアップ
```sql
CREATE DATABASE IF NOT EXISTS testdb;

USE testdb;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100),
    age INT
)
```
- Goコードの作成
- Dockerを使ってMySQLサーバーをセットアップ
```sh
docker run --name mysql-container -e MYSQL_ROOT_PASSWORD=my-secret-pw -d -p 3306:3306 mysql:latest
```