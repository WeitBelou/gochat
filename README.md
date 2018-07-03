# Go Chat
Тестовое задание

# Запуск
```bash
./bin/run.sh # Будет слушать на 127.0.0.1:8080
```

# Как пользоваться
1) Запрос на `/api/v1/auth/register` с телом
    ```json
    {
      "login": "sdasd",
      "password": "xxxxxx"
    }
    ```
    вернёт ответ
    ```json
    {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA2NDQxMTcsImlhdCI6MTUzMDYzNjkxNywic3ViIjoic29zIiwibmlja25hbWUiOiJzb3NzaXRvIn0.ikVEOPSO49b2uyX4bJNJDkPhacnaGeLcBy7hrDBnLio"
    }
    ```

2) Этот токен используется в апи концах, требующих авторизации: добавляется в HTTP заголовок
Authorization с префиксом `"Bearer "`, например: `"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MzA2NDQxMTcsImlhdCI6MTUzMDYzNjkxNywic3ViIjoic29zIiwibmlja25hbWUiOiJzb3NzaXRvIn0.ikVEOPSO49b2uyX4bJNJDkPhacnaGeLcBy7hrDBnLio"`.

3) Логин POST `/api/v1/auth/login`

4) Смена ника POST `/api/v1/profile/edit`

5) Список сообщений GET `/api/v1/messages`

6) Новое сообщение POST `/api/v1/messages`

## Как подписаться на новые сообщения с помощью websocket:
1) Запрос на `/api/v1/auth/ws` вернёт одноразовый токен.

2) Далее обычный запрос на установление ws соединения `/api/v1/messages/ws`

3) Теперь при публикации нового сообщения они будут прилетать сюда.