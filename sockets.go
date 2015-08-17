/* HTTP GET with plain sockets */

package main

import (
    "fmt"
    "net"
    "bufio"
)

func main(){
    conn, _ := net.Dial("tcp", "crypto.cat:80")
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    reader := bufio.NewReader(conn)
    for {
        repl, err := reader.ReadString('\n')
        fmt.Print(repl)
        if err != nil {
            fmt.Println(err)
            conn.Close()
            break
        }
    }
}
