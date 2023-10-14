# Агрегатор новостей (часть итоговой аттестации)

[Агрегатор из прошлого дз](https://github.com/DoktorGhost/36a2) требовал некоторой доработки:  
1. Добавить поиск по названию;
2. Добавить постраничную навигацию к запросу списка новостей;
3. Добавить журналирование запросов.

Для тестирования сервиса отдельно, можно воспользоваться dockerfile, который есть в проекте:
``` go
docker build -t my-postgres .  
docker run -d --name my-postgres-container -p 5432:5432 my-postgres
go run server.go
```

*Если же запускаем в связке с сервисами [Комменатриев](https://github.com/DoktorGhost/comments) , [Верификацией](https://github.com/DoktorGhost/verification) и [APIGateway](https://github.com/DoktorGhost/api_gateway) , то нужно воспользоваться docker-compose из сервиса комменатриев.*
