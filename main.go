package main

import (
    "errors"
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"

    "github.com/11061055/goleveldb_server/manager"
)

var levelDBManager manager.LevelDBManager

func main()  {

    levelDBManager.Construct()

    http.HandleFunc("/data",DataHandler)

    http.ListenAndServe("0.0.0.0:8880",nil)
}

func DataHandler(w http.ResponseWriter,r *http.Request)  {

    path := ""
    hour := ""
    act  := ""
    key  := ""
    body := []byte{}

    body, err := ioutil.ReadAll(r.Body)
    if  err != nil {
      return
    }

    purl, err := url.Parse(r.URL.String())
    if  err != nil {
      return
    }
    path = purl.Path

    query, err := url.ParseQuery(purl.RawQuery)
    if  err != nil {
      return
    }

    if val, ok := query["act"]; ok {
      act = val[0]
    }

    if val, ok := query["hour"]; ok {
      hour = val[0]
    }

    if val, ok := query["key"]; ok {
      key = val[0]
    }

    if len(path) == 0 || len(act) == 0 || len(hour) == 0 || len(body) == 0 {
      return
    }

    fmt.Println(path)
    fmt.Println(act)
    fmt.Println(hour)
    fmt.Println(string(body[:]))

    db, err := levelDBManager.Open(hour)
    if err != nil {
      w.Write([]byte(err.Error()))
      return
    }
    defer levelDBManager.Close(db)

    var ret []byte = []byte("success")

    switch act {

      case "put":
        err      = db.Put([]byte(key), body)

      case "get":
        ret, err = db.Get([]byte(key))

      case "del":
        err      = db.Del([]byte(key))

      default:
        err      = errors.New("invalid act")

     }

     if err != nil {
       w.Write([]byte(err.Error()))
     } else {
       w.Write(ret)
     }
}