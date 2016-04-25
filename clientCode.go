package main                                                

import (
    "fmt"

    "github.com/kless/osutil/user/crypt/sha512_crypt"
)

func main() {
    c := sha512_crypt.New()
    hash, err := c.Generate([]byte("rasmuslerdorf"), []byte("$6$usesomesillystringforsalt"))
    if err != nil {
        panic(err)
    }

    fmt.Println(hash)
}