## Build Your Own - Redis (Simplified)

> Redis is an open source, in-memory, key-value data store most commonly used as a primary database, cache, message broker, and queue.

[https://developer.redis.com/explore/what-is-redis/](more about redis)

Implementasi redis yang disederhanakan (dengan sangat), menggunakan bahasa Go.
Tujuannya, untuk mempelajari konsep-konsep pada pemrograman jaringan dan struktur data. Maka dari itu, sebisa mungkin tidak ada implementasi pihak ketiga dalam bentuk library atau framework pada program ini.
Arsitektur dan design pada program ini mengacu dari buku "Build Your Own Redis with C/C++" dari build-your-own.org.

## Server

Redis server untuk menjalankan program dengan spesifikasi redis.
Menggunakan koneksi TCP dengan model client/server untuk mentransmisikan data.
Server akan menerima data, menampilkannya pada terminal, kemudian mengirimkan kembali data string `PONG`.

Program akan mengikat koneksi pada alamat `localhost` dan port `1234`.

```bash
# komando menjalankan program

go run server/main.go
```

## Client

Tools client untuk berkominikasi dengan server redis.
Client akan mengirimkan string `PING`, menerima data, kemudian menampilkannya pada terminal.

Program akan mencoba untuk mengkoneksikan server pada alamat `localhost` dan port `1234`.

```bash
# komando menjalankan program

go run client/main.go
```
