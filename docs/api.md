## json response
- 200

    ```
        {
            "body" : {
                "field1" : "value1",
                "field2" : "value2",
            } 
        }
    ```
    - ниже в секции out указываются названия поля 
- 400

    ```
        {
            "message" : "message"
        } 
    ```
    - ниже в секции out указываются возможные сообщения об ошибках
- 500

    ```
        нет тела
    ```


## auth
- **POST /user/signup**
    - in (json body)
        - email
        - password
        - avatar_url
        - birthday
    - out
        - 200
        - 400, fail messages (json body) 
            - failed to parse json string from the body
            - failed to parse birthday from string to time
            - user already exists
            - email or password not validated
        - 500, no messages 
- **POST /user/signin**
    - notes
        - не нужны куки
        - не нужен csrf токен
    - in (json body)
        - email
        - password
    - out
        - 200
        - 400, fail messages (json body) 
            - failed to parse json string from the body
            - user not found
        - 500, no messages 
- **POST /user/logout**
    - notes
        - нужны куки
        - нужен csrf токен
    - out
        - 200
        - 400, fail messages (json body) 
            - failed to parse json string from the body
            - user not found
        - 500, no messages 
## user
- **POST /user/avatar/upload**
    - notes
        - нужны куки
        - нужен csrf токен
        - принимаем multipart/form-data
    - in (multipart/form-data)
        - ```<input name="avatar">```
    - out
        - 200
        - 400, fail messages (json body) 
            - can't update avatar without session
            - the size exceeded the maximum size equal to 10 mb
            - failed to parse avatar file from the body
            - avatar file can't be read
            - avatar does not have type: image/jpeg
        - 500, no messages 
- **GET /user/csrf**
    - notes 
        - нужны куки
        - ручка для выдачи нового csrf токена на основе сессии
    - out
        - 200
            - csrf
        - 400, fail messages (json body) 
        - 500, no messages 

- **GET /user/profile**
    - notes
        - нужны куки
        - возвращает все информацию о пользователе по сессии
    - out
        - 200
            - email
            - birthday
            - avatar_url
        - 400, fail messages (json body) 
            - can't get info without session
        - 500, no messages 
        
## film selections
