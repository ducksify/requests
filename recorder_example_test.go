package requests_test

import (
    "context"
    "fmt"
    "testing/fstest"

    "github.com/ducksify/requests"
)

func ExampleReplayFS() {
    fsys := fstest.MapFS{
        "fsys.example - MKIYDwjs.res.txt": &fstest.MapFile{
            Data: []byte(`HTTP/1.1 200 OK
Content-Type: text/plain; charset=UTF-8
Date: Mon, 24 May 2021 18:48:50 GMT

An example response.`),
        },
    }
    var s string
    const expected = `An example response.`
    if err := requests.
        URL("http://fsys.example").
        Transport(requests.ReplayFS(fsys)).
        ToString(&s).
        Fetch(context.Background()); err != nil {
        panic(err)
    }
    fmt.Println(s == expected)
    // Output:
    // true
}
