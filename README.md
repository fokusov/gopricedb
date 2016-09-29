# Описание
gopricedb загружает курсы валют с сервиса [fixer.io](http://fixer.io/) в формате файла pricedb для ledger-cli

# Запуск
Установите [Go](https://golang.org/) , скачайте репозиторий и запустите командой

```
go run main.go
```

после этого в папке появится файл prices.dat (если он не был создан ранее), который можно подключить к ledger так:

```
ledger -f journal.dat bal -X USD --price-db prices.dat
```

также можете создать бинарный файл командой

```
go build
```

или скачать готовый бинарник gopricedb.exe если нет возможности установить Go.