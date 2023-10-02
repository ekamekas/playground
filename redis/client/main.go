package main

import (
    "fmt"
    "net"
    "redis/protocol"
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

    conn.Close()
}

func process(conn net.Conn) {
    serialized, _ := protocol.Serialize([]string{ "PING" })

    _, error := conn.Write(serialized)

    if(nil != error) {
        fmt.Printf("[CONN] failed to write data. Cause: %s\n", error.Error())
        return
    }

    buffer := make([]byte, protocol.MESSAGE_MAX)

    _, error = conn.Read(buffer)

    if(nil != error) {
        fmt.Printf("[CONN] failed to read data. Cause: %s\n", error.Error())
        return
    }

    deserialized, _ := protocol.Deserialize(buffer)

    fmt.Printf("[CONN] received data: %s\n", deserialized)
}
