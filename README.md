# vezdecode-password-generator
Password generator (information security 10)

Генератор паролей с возможностью проверки существующих.

Написан на языка Go , требуется Go не ниже 1.14

Для запуска необходимо перейти в корневую папку проекта и выполнить 
`go run . [-args] [GENERATE|CHECK]`

Для настройки используются следующие аргументы:
```
-l - требуемая длина пароля 
-d - требуются ли в пароле цифры
-s - требуются ли в пароле специальные символы
-u - требуются ли в пароле символы верхнего регистра

-input - только для действия CHECK, позволяет указать путь на файл в котором должны находится пароли для проверки.
    Каждый пароль на новой строчке.
-report - только для действия CHECK и указанного файла с паролями, выводит отчет о соответствующих и несоответствующих паролях.
```

В папке files лежит пример файла с паролями.

###Примеры запуска:

Генерация пароля
```
go run . -l=10 -d -s GENERATE
```
Вывод
```
r421lt@ly@
```
---
Проверка пароля
```
go run . -l=3 -d -s -u CHECK abc2!!_
```
Вывод
```
Check failed: password has no uppercase
```
---
Проверка файла с выводом отчета
```
>go run . -l=3 -d -s -u -file=files/passwords.txt -report CHECK
``` 
Вывод:
```
`qwerty` Check failed: password has no uppercase
`qwerty123` Check failed: password has no uppercase
`qwe` Check failed: password has no uppercase
`qq` Check failed: length of password is 2, required: 3
`qwerqwr12` Check failed: password has no uppercase
`qw!!#evd30` Check failed: password has no uppercase
`FGFDGF#2gfd2!` meets the requirements
`FDFSEE234,d_!!lfff` meets the requirements
`asdQWE123!!` meets the requirements
`VDSV234WOoe_` meets the requirements
REPORT:
GOOD PASSWORDS (4):
FGFDGF#2gfd2!
FDFSEE234,d_!!lfff
asdQWE123!!
VDSV234WOoe_
BAD PASSWORDS (6):
qwerty
qwerty123
qwe
qq
qwerqwr12
qw!!#evd30
Total checks: 10
```