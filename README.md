## ディレクトリ構成

```
echo_test/
├── db
│   └── sample.db
├── handler
│   ├── auth.go
│   ├── articles.go
│   └── todos.go
├── model
│   ├── db.go
│   ├── aritcle.go
│   ├── todo.go
│   └── user.go
├── public
│   ├── assets
│   │   └── js
│   │       ├── login.js
│   │       ├── signup.js
│   │       ├── artlcies.js
│   │       ├── artlcie_dteail.js
│   │       └── todos.js  
│   ├── index.html
│   ├── login.html
│   ├── signup.html
│   ├── articels.html
│   ├── articel_dteail.html
│   └── todos.html
├── README.md
├── go.mod
├── go.sum
├── main.go
└── router.go
```

## サーバーの実行

```
$ go run main.go router.go
```

## 画面構成

### トップページ

`GET /`

### ユーザー登録

`GET /signup`

### ログイン

`GET /login`

### Todo一覧

`GET /todos`

### Article一覧

`GET /articles`

## 使用技術

- Go言語 (https://golang.org/)
- Echo (https://echo.labstack.com/)
- GORM (http://gorm.io/)
- SQLite (https://www.sqlite.org/index.html)
- JWT (https://jwt.io/)
