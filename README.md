# avito-chats-app
## Запуск
Запуск осуществуляется с помощью ```docker-compose up``` из корня приложения
## Описание проекта
Основные моменты:
- Индентификаторы пользователей и чатов - числовые значения
- В качестве HTTP роутера используется ```https://github.com/gorilla/mux```
- В качестве хранилища используется база MySQL (v. 8.0.21)
  + База данных содержит 4 таблицы:
    * **users** - Пользователи
    * **chats** - Чаты 
    * **messages** - Сообщения
    * **user_chat_relation** - Талбица для реализации отношения `многие ко многим` чатов с пользователями
  + Структуру таблиц можно посмотреть в файле инициализации ```_sql/db.sql```
  
## Структура проекта
Пытался структурировать проект в соответствии с ```https://github.com/golang-standards/project-layout```
```
AvitoChats
│   README.md
│   Dockerfile
|   docker-compose.yml
|
└───bin
│   | chatsapp
|
└───_sql
│   | db.sql
│
└───cmd
│   └───chatsapp
│       │   main.go
│   
└───pkg
|   |
│   └───database
│   |   │   repo.go
│   |   │   repoChat.go
│   |   │   repoMessage.go
│   |   │   repoUser.go
|   |
│   └───handlers
│   |   │   handler.go
│   |   │   chats.go
│   |   │   messages.go
│   |   │   users.go
|   |
│   └───models
│       │   chat.go
│       │   message.go
│       │   user.go
|
└───script
    │   wait-for-it.sh
```

- ```bin/chatsapp``` - бинарник для записка сервера в контейнере
- ```_sql/db.sql``` - файл инициализации базы данных с созданием нужных таблиц
- ```cmd/chatsapp/main.go``` - файл для запуска приложения
- ```pkg/database``` - реализация работы с БД по паттерну "Репозиторий"
- ```pkg/handlers``` - HTTP обработчики для запросов
- ```pkg/models``` - описания объектов
- ```script/wait-for-it.sh``` - скрипт для ожидания доступности TCP хоста с портом ```https://github.com/vishnubob/wait-for-it```
  > Используется во время развертывания
