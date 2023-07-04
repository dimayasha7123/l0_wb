# L0_WB

Для запуска:
```shell
docker compose up --build
```

Для уничтожения:
```shell
docker compose down -v
```

Для записи в канал (пишет файлик model/example.json):
```shell
make pub
```

Для записи может потребоваться:
```shell
go get github.com/nats-io/stan.go/@v0.10.4
```

Есть [демка](https://www.youtube.com/watch?v=Nwmzp6UI9Rk) с комментариями по коду. 

PS: забыл показать на видео, но валидация работает, честно-честно)