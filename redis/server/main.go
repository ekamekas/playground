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
    listener, error := net.Listen(SERVER_TYPE, SERVER_HOST + ":" + SERVER_PORT)

    if(nil != error) {
        fmt.Println("failed to create a server")
        panic(error)
    }

    fmt.Println("server is created, and waiting for client connection")

    defer listener.Close()
    for true {
        connection, error := listener.Accept()

        if(nil != error) {
            fmt.Println("failed to accept incoming connection")
            continue
        }

        process(connection)
    }
}

func process(conn net.Conn) {
    fmt.Printf("[CONN] incoming connection from %s\n", conn.RemoteAddr())

    defer conn.Close()
    defer fmt.Printf("[CONN] closing connection from %s\n", conn.RemoteAddr())
    
    for {
        buffer := make([]byte, protocol.MESSAGE_MAX)

        _, error := conn.Read(buffer)

        if(nil != error && "EOF" == error.Error()) {
            break
        } else if(nil != error) {
            fmt.Printf("[CONN] failed to read data. Cause: %s\n", error.Error())
            break
        }

        deserialized, _ := protocol.Deserialize(buffer)

        fmt.Printf("[CONN] received data: %s\n", deserialized)

        serialized, _ := protocol.Serialize([]string{ "PONG" })

        conn.Write(serialized)
    }
}
