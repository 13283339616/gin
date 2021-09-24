module service

go 1.17

require (
	github.com/13283339616/orm v0.0.0
	github.com/13283339616/vo v0.0.0
	github.com/13283339616/util v0.0.0
)

require (
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	gorm.io/driver/mysql v1.1.2 // indirect
	gorm.io/gorm v1.21.15 // indirect
)

replace (
	github.com/13283339616/orm => ../orm
	github.com/13283339616/vo => ../vo
	github.com/13283339616/util => ../util
)
