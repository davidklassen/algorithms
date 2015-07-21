package main

import "os"
import "fmt"
import "io/ioutil"

func main() {
    bytes, err := ioutil.ReadAll(os.Stdin)

    if err != nil {
        panic(err)
    }

    fmt.Println(string(bytes))
}
