# snippetbox
Describe

Чтобы отключить просмотр списка файлов из дирректории:
Веб-сервер всегда ищет сначала файл index.html и если он пустой, то ничего не отображается
Создать его можно командой:
```go
$ find ./ui/static -type d -exec touch {}/index.html \;
```
или с помощью http.FileSystem