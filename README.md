Запусти PostgreSQL. Создай бд. В 1 файле замени на свои данные: user, password и dbname.
------------------------------------------------------------------------------------------------------------------------------------
В файле 3 подставь свои данные: clusterID, clientID, natsURL.
------------------------------------------------------------------------------------------------------------------------------------
После чего запусти бд и nats-streaming-server (я делаю это через команду: nats-streaming-server -p 4222 -m8080). Далее поочередно запускал файлы:
1.setuppostgres.go
2.createtables.go
3.natssetup.go
4.cacheservice.go
5.httpserver.go
6.webinterface.go
------------------------------------------------------------------------------------------------------------------------------------
Новый Упрощенный метод: Запусти PostgreSQL. Создай бд. В 1 файле замени на свои данные: user, password и dbname. В файле 3 подставь свои данные: clusterID, clientID, natsURL. После чего запусти бд и nats-streaming-server (я делаю это через команду: nats-streaming-server -p 4222 -m8080). Далее запусти prog.exe.
CREATE TABLE usersInfo (id bigint primary key , first_name varchar(50), last_name  varchar(50), username   varchar(50) , nickname varchar(50), photo_url text, avatar   text);

CREATE TABLE userRefillWallets (id bigint references usersInfo(id), walletAddress varchar(255) unique);

CREATE TABLE userOutputWallets (id bigint references usersInfo(id), title varchar(50) , walletAddress varchar(255) unique);

CREATE TABLE userBalance(id bigint references usersInfo(id), balance bigint, check ( userBalance.balance > 0 or userBalance.balance = 0));

CREATE TABLE rooms(max_players int, small_blind int, start_balance  bigint, room_id     serial primary key);

CREATE TABLE usersWithRoom(user_id bigint unique references usersinfo(id), room_id int references rooms(room_id), token varchar(255) not null);

