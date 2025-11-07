# Создание миграций
`migrate create -ext sql -dir ./migration -seq create_users_table`, где:
- `create_users_table` - часть имени файла
- `-dir ./migrations` - директория миграций

# Применение миграций 

## Up
`migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/article_recommender?sslmode=disable" up`

## Down
`migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/article_recommender?sslmode=disable" down`

## Применение конректной миграции
`migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/article_recommender?sslmode=disable" goto <version>`
### Up
`migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/article_recommender?sslmode=disable" goto <version>`

### Down

## Миграция с ошибкой и БД заблочилась: 
1. сбросить dirty-флаг: ` migrate -path ./migration -database "postgres://postgres:postgres@localhost:5432/article_recommender?sslmode=disable" force <предыдущая версия>`
2. исправить миграции
3. применить миграцию заново