package main

import (
    "bytes"
    "encoding/gob"
    "fmt"
)


func SerializeVM(p VM) []byte {
    var buf bytes.Buffer
    encoder := gob.NewEncoder(&buf)
    err := encoder.Encode(p)
    if err != nil {
        fmt.Println("Error encoding struct:", err)
        return nil
    }
    byteArray := buf.Bytes()
    return byteArray
}

func DeSerializeVM(p []byte) VM {
    var decoded VM
    decoder := gob.NewDecoder(bytes.NewReader(p))
    _ = decoder.Decode(&decoded)
    return decoded;
}
