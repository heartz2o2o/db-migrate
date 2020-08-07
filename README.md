# db-migrate

mysql migrate

## Getting Started

### Prerequisites

1. Installed [golang](https://golang.org/doc/install)

## Running

### Golang專案引用db-migrate

1. 定義 env 變數, 其中 `Environment.Dir` 為 sql folder
2. 建立 sql folder, 將 sql files 放置該目錄底下, 檔名格式為 `yyyyMMddHHmmss-簡述.sql`

```golane

import (
	command "github.com/heartz2o2o/db-migrate/command"
)

func main() {
	env := &command.Environment{
		Dialect:    "mysql",
		DataSource: "root:123456@tcp(localhost:3306)/bac?parseTime=true",
		Dir:        "./sql"}
	command.SetEnvironment(env)
	migrate := command.UpCommand{}
	migrate.Run([]string{})
}

```

### 使用CLI

#### Build binary file

```bash
$ go build -o db-migrate ./main
```

#### 配置dbconfig.yml

```bash
$ cat dbconfig.yml
development:
    dialect: mysql
    datasource: root:123456@tcp(localhost:3306)/bac?parseTime=true
    dir: sql
```

#### Run

```bash
$ ./db-migrate status
+---------------+---------+
|   MIGRATION   | APPLIED |
+---------------+---------+
| 1-initial.sql | no      |
| 2-record.sql  | no      |
+---------------+---------+

$ ./db-migrate up
Applied 2 migrations

$ ./db-migrate status
+---------------+-------------------------------+
|   MIGRATION   |            APPLIED            |
+---------------+-------------------------------+
| 1-initial.sql | 2020-08-07 12:44:20 +0000 UTC |
| 2-record.sql  | 2020-08-07 12:44:20 +0000 UTC |
+---------------+-------------------------------+

$ ./db-migrate down
Applied 1 migration

$ ./db-migrate status
+---------------+-------------------------------+
|   MIGRATION   |            APPLIED            |
+---------------+-------------------------------+
| 1-initial.sql | 2020-08-07 12:44:20 +0000 UTC |
| 2-record.sql  | no                            |
+---------------+-------------------------------+
```
