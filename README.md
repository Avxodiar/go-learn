# Язык Go

[Сайт проекта](https://go.dev/)

Актуальная версия 1.22 от 02.2024

## Учебники и туториалы

[Тур по Go от создателей](https://go.dev/tour/welcome/1)

[Маленькая книга о Go (от 5 августа 2015)](https://sefus.ru/little-go-book/)

[Введение в программирование на Go (от 2019)](http://golang-book.ru)

[Туториал metanit (от 18 февраля 2023)](https://metanit.com/go/tutorial/1.1.php)

## Книги
По неподтвержденным данным большинство это водянистая вода.
Из относительно свежих можно порекомендовать только одну:
* Донован, Керниган: Язык программирования Go (2020)

https://www.labirint.ru/books/533140/

## Полезные ссылки
[Go в примерах](https://gobyexample.com.ru)

[Список фреймворков и библиотек](https://github.com/avelino/awesome-go)


## Установка
[Руководство от создателей](https://go.dev/doc/install)

Если возникают ошибки подобные "tar: go/somefilename: Cannot open: No such file or directory"
добавьте sudo перед tar в команде установке

`sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go.arcive.tar.xz`

В профайл добавить проще так

`echo 'export PATH=$PATH:/usr/local/go/bin' >> $HOME/.profile`

Не забываем проверить корректность установки

`go version`

## IDE
* [Jetbrains GoLand](https://www.jetbrains.com/go/)
* [MS VSCode](https://code.visualstudio.com/)

  с расширением Go и инструментами (Ctrl + Shift + P, затем выбрать Go: Install/Update Tools):
  * gotests – инструмент для автоматической генерации юнит-тестов (тестовых функций) на основе кода;
  * gomodifytags – инструмент для добавления, изменения и удаления тегов структур в коде;
  * impl – инструмент, который автоматически генерирует заглушки (scaffolding) для методов интерфейсов в коде;
  * goplay – интерактивная среда для запуска и тестирования фрагментов кода прямо в браузере;
  * dlv (Delve Debugger) – отладчик для Go, который позволяет разработчикам отлаживать свои программы;
  * staticcheck – инструмент для статического анализа кода с целью поиска потенциальных ошибок и улучшения качества кода;
  * gopls – сервер языка Go, который предоставляет возможности автодополнения, структурного анализа и других функций поддержки кода.

  [Как начать писать на Go в VSCode]( https://www.youtube.com/watch?v=Ptk_zcOVPg4)
