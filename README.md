# Сервис обработки заказов
Демонстрационный сервис с простейшим интерфейсом для обработки заказов. Данные о заказах хранятся в базе данных PostgresSQL. В сервисе реализована подписка на топик Kafka,  данные полученные из брокера сохраняются в БД и кешируются(in memory).
Сервис предоставляет возможность получить данные о заказе по его UID.

## Требования
* Версия Go 1.22.2
* Docker-Compose
* Make
* WRK (Для нагрузочного тестирования)

## Конфигурация
В директории `service/configs` находится файл `conf.yaml` с конфигурациями для запуска сервиса и его инфраструктуры
```
database:
  dmbs: postgres
  postgres:
    host: localhost
    port: 5436
    user: myuser
    password: mypassword
    dbname: mydb
    sslmode: disable
server:
  host: localhost
  port: 8080
broker:
  name: kafka
  kafka:
    host: localhost
    port: 9092
    group: myGroup
    reset: earliest
    autoCommit: true
```

## Тестирование 
1) Модульное тестирование
```
make test
```
2) Нагрузочное тестирование
```
wrk -t12 -c400 -d10s http://localhost:8080/orders/b563febdb2b84b6test
``` 

## Инструкция по запуску 

1) Запуск инфраструктуры(Postgres, Kafka)
```
make up
```
2) Запуск сервиса
```
make orderService
```

## Использование

1) Получение html страницы с интерфейсом для получения заказа по UID
```
GET /page HTTP/1.1
Host: 127.0.0.1:8080
Accept: text/html
```

2) Получение заказа по UID
```
GET /orders/{id} HTTP/1.1
Host: 127.0.0.1:8080
Accept: application/json
```
## Отправка сообщения в kafka

Для отправки сообщения в kafka можно использовать скрипт
```
docker-compose exec kafka kafka-console-producer --bootstrap-server kafka:29092 --topic ord
```
