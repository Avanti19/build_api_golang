package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
    "strings"
)

const version = "1.1"
const progname = "AV@StringReplacer"

const default_ip = "0.0.0.0"
const default_port = "9090"

func init() {
    fmt.Printf("Initalizing program %s (ver. %s)\n", progname, version)
    http.HandleFunc("/", home_page)
    http.HandleFunc("/replace", find_and_replace)
}

func replace(str, keyword, replacement string) (int, string) {
    if(len(str) == 0) {
        // fmt.Printf("Input string is empty!\n")
        return 1, ""
    }

    fmt.Printf("Replace in string: \"%s\" , Keyword: \"%s\" with String: \"%s\"\n", str, keyword, replacement)
    c := strings.Count(str, keyword)

    if(c == 0) {
        return 1, ""
    }

    // fmt.Printf("String '''%s''' contains %d instances of string '''%s'''.\n---\nReplacing string...\n---\n", str, c, strFindVal)

    return 0, strings.Replace(str, keyword, replacement, -1)
}

func home_page(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "You are at home!\n\n-%s (ver. %s)", progname, version)
}

type Replacer struct {
    InputText string
    KeywordString string
    ReplacementString string
}

func find_and_replace(w http.ResponseWriter, r *http.Request) {
    // fmt.Fprintf(w, "You are at home!\n\n-%s (ver. %s)", progname, version)
    var m Replacer

    fmt.Printf("Request=\n%s\n\n", r.Body)
    err := json.NewDecoder(r.Body).Decode(&m)

    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fmt.Printf("Parsed message:\n%+v", m)

    var e int
    var s string

    e, s = replace(m.InputText, m.KeywordString, m.ReplacementString)

    if e==0 {
        fmt.Printf("Replaced string=\n%s\n", s)
    } else {
        s=m.InputText
    }

    fmt.Fprintf(w, "{\"ReplacedText\":\"%s\",\"status\":%d}", s, e)
    fmt.Printf("Reply=\n{\"ReplacedText\":\"%s\",\"status\":%d}\n", s, e)

    return
}

func main() {
    var listener_ip string = default_ip
    var listener_port string = default_port
    switch argc := len(os.Args); argc {
        case 2:
            listener_ip = os.Args[1]
        case 3:
            listener_ip = os.Args[1]
            listener_port = os.Args[2]
    }

    fmt.Printf("Starting String Replacement Service on http://%s:%s/replace ...", listener_ip, listener_port)
    log.Fatal(http.ListenAndServe(listener_ip+":"+listener_port, nil))
}
