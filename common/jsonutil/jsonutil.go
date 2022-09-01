package jsonutil

import (
    "encoding/json"
    "fmt"
)

func PrintJsonStringWithIndent(v interface{}) {
    b, err := json.MarshalIndent(v, "", "  ")
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(b))
}

