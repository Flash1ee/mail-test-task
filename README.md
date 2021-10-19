# mail-test-task
cli -- клиент, для некоторого абстрактного удаленного tcp/ip сервера cube для проверки токенов, по бнф описанию протокола.

# Постановка задачи

Реализовать утилиту cli -- клиент, для некоторого абстрактного удаленного tcp/ip сервера cube для проверки токенов, по бнф описанию протокола.

Для соискателей на вакансию программиста Си:  
Реализация клиента должна быть на языке Си под компилятор gcc.  
Для реализации клиента разрешается использовать только libc.  

Для соискателей на вакансию программиста Go:  
Реализация клиента должна быть на языке Go, v 1.12.  
Для реализации клиента разрешается использовать только стандартную библиотеку языка Go, внешние библиотеки использовать запрещено.  

Реализация должна иметь систему сборки и функциональные тесты (в качестве ответов для тестирования использовать mock заглушки, в соответствии с описанием протокола).  
Все, что не оговорено в задаче, остается на усмотрение программиста.  
  
Параметры cube cli клиента: host port token scope  
  
Срок: 1 неделя  
Большая просьба задание никуда не выкладывать.  
  
Пример запроса и ответа проверки токена `abracadabra` и scope `test`
```bash
$ cube_cli cube.testserver.mail.ru 4995 abracadabra test
client_id: test_client_id
client_type: 2002
expires_in: 3600
user_id: 101010
username: testuser@mail.ru
```
  
Пример запроса и ответа неуспешной проверки токена abracadabra и scope xxx  
```
$ cube cube.testserver.mail.ru 4995 abracadabra xxx
error: CUBE_OAUTH2_ERR_BAD_SCOPE
message: bad scope
```

CUBE IPROTO protocol  
```
<packet> ::= <request> | <response>
<request> ::= <header><svc_request_body>
<response> ::= <header><response_body>
<header> ::= <svc_id><body_length><request_id>
<svc_id> ::= <int32> - идентификатор CUBE сервиса
<body_length> ::= <int32> - длина тела запроса
<request_id> ::= <int32> - идентификатор запроса, возвращается в ответе
<return_code> ::= <int32> - код ответа, описаны ниже
<string> ::= <str_len><str>
<str_len> ::= <int32> - длина строки, больше 0
<str> ::= <int8>+ - строка
<int32> ::= целочисленное число со знаком в бинарном виде, размер 4 байта, порядок байт little-endian
<int8> ::= целочисленное число со знаком в бинарном виде, размер 1 байт
```
CUBE OAUTH2 предоставляет возможность проверить oauth2 token и scope.
```
<svc_id> ::= 0x00000002
<svc_msg> ::= <int32> - номер сообщения для проверки access token и scope, равен 0x00000001
```
Запрос на проверку access_token и scope: 
```
<svc_request_body> ::= <svc_msg><token><scope>
<token> ::= <string> - проверяемый токен
<scope> ::= <string> - проверяемый scope

Ответ:
<svc_ok_response_body> ::= <return_code><client_id><client_type><username><expires_in><user_id> - return_code == 0x00000000 

<client_id> ::= <string> - идентификатор клиента
<client_type> ::= <int32> - тип клиента
<username> ::= <string> - имя пользователя
<expires_in> ::= <int32> - оставшееся время жизни токена в секундах
<user_id> ::= <int64> - идентификатор пользователя
<svc_error_response_body> ::= <return_code><error_string> - return_code != 0x00000000 
<error_string> ::= <string> - сообщение об ошибке
```
Коды ошибок (<return_code>):
```
0x00000000 -- CUBE_OAUTH2_ERR_OK (<без error_string> -- код успешного ответа)
0x00000001 -- CUBE_OAUTH2_ERR_TOKEN_NOT_FOUND (token not found -- токен не найден)
0x00000002 -- CUBE_OAUTH2_ERR_DB_ERROR  (db error -- ошибка базы данных)
0x00000003 -- CUBE_OAUTH2_ERR_UNKNOWN_MSG (unknown svc message type -- неизвестный svc_msg)
0x00000004 -- CUBE_OAUTH2_ERR_BAD_PACKET (bad packet -- неверный формат запроса)
0x00000005 -- CUBE_OAUTH2_ERR_BAD_CLIENT (bad client -- clientid и/или clientsecret не прошли проверку)
0x00000006 -- CUBE_OAUTH2_ERR_BAD_SCOPE (bad scope -- неправильный scope)
```

