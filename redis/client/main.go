package main

import (
    "fmt"
    "net"
)

const (
    SERVER_HOST = "localhost"
    SERVER_PORT = "1234"
    SERVER_TYPE = "tcp"
)

func main() {
    conn, error := net.Dial(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)

    if(nil != error) {
        fmt.Println("failed to accept incoming connection\n")
        return
    }

    process(conn)
    process(conn)
    process(conn)

    conn.Close()
}

func process(conn net.Conn) {
    _, error := conn.Write([]byte("PING"))

    if(nil != error) {
        fmt.Printf("[CONN] failed to write data. Cause: %s\n", error.Error())
        return
    }

    buffer := make([]byte, 1024)

    size, error := conn.Read(buffer)

    if(nil != error) {
        fmt.Printf("[CONN] failed to read data. Cause: %s\n", error.Error())
        return
    }

    data := string(buffer[:size])

    fmt.Printf("[CONN] received data: %s\n", data)
}
