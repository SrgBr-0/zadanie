Запусти PostgreSQL. Создай бд. В 1 файле замени поставь свои данные user, password и dbname.
------------------------------------------------------------------------------------------------------------------------------------------------------------ 
В файле 3 подставь свои данные clusterID, clientID, natsURL.
------------------------------------------------------------------------------------------------------------------------------------------------------------ 
После чего для простоты запусти бд и nats-streaming-server (я делаю это через команду: nats-streaming-server -p 4222 -m8080). Далее поочередно запускал файлы:
1.setuppostgres.go
2.createtables.go
3.natssetup.go
4.cacheservice.go
5.httpserver.go
6.webinterface.go
