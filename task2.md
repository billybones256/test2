# Пирамида.
Напишите программу, которая выводит пирамиду из символов "@" заданной высоты.
На вход программе подается число ```n``` - высота пирамиды.
На выходе - сама пирамида.

## Где тесты?
Тесты смотрите в файле main_test.go


## Как решать?
Сначала решите все обычным спосбом (сделайте простую программу ```main.go``` с функцией ```Pyramid(n int) string ```) . После чего протестируйте ее, использую файл ```main_test.go```. В случае, если всё ок - переходите к следующей стадии.


## Следуюящая стадия.
Соберите grpc сервер и укоплектуйте его в контейнер с выброшенным наружу портом ```8081```. Описание ```.proto``` файла смотрите рядом с заданием. grpc сервер обрабатывает клиентскую строку таким же образом, как действует ```Pyramid(...)...```. 
Соберите локальный клиент (не нужно помещать в контейнер).
Изучаемое число принимается клиентом как аргумент командной строки. Ответ - пирамида.


Структура проекта - аналогична ```georpc```.
Необходимо наличие: 
* makefile с необходимыми (по вашем мнению) мейкапами.
* Dockerfile с торчащим 8081 портом
* grpc-Клиента, принимающего порт и аргумент через командную строку
* grpc-Сервера со всей логикой
