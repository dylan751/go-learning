## Framework

- [GIN](https://github.com/gin-gonic/gin) là một web framework cho Golang. Gin có thể giúp chúng ta xây dựng nhanh các web/api service Golang với cú pháp rất gọn (giống với Express bên NodeJS). Một sự lựa chọn khác cũng khá thú vị là [Echo](https://echo.labstack.com/).
- [GORM](https://gorm.io/) là một thư viện ORM (Object-relational Mapping) dành cho Golang. Thư viện này giúp các developer Golang đỡ phải thực hiện các câu lệnh SQL thuần tuý. Đương nhiên sự tiện lợi sẽ đánh đổi bằng hiệu năng. Nếu các bạn yêu thích SQL và muốn service mình chạy nhanh hơn nữa thì cân nhắc dùng [sqlx](https://github.com/jmoiron/sqlx).

## Database

```sql
USE `go_gin_server`

DROP TABLE IF EXISTS `todo_items`
CREATE TABLE `todo_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(150) CHARACTER SET utf8 NOT NULL,
  `status` enum('Doing','Finished') DEFAULT 'Doing',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
```

## How to run the project

- Run command: `go run main.go`
- Open `http://localhost:8080` - API Endpoint of Gin Framework
