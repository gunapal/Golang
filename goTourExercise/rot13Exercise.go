package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 *rot13Reader) Read(p []byte)(n int, err error) {
    
    n, err = r13.r.Read(p)
    
    if err != nil {
        return
    }
    
    for i := 0; i < n; i++ {
        rotate13(&p[i])
    }
    
    return
}

func rotate13(inChar *byte){
    
    c := *inChar
    
    if c >= 'a' && c <= 'm' || c >= 'A' && c <= 'M' {
        *inChar = c + 13
    }else {
        *inChar = c- 13
    }
}

func main() {
    s := strings.NewReader(
        "Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}