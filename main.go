package main

import (
    "errors"
    "net/http"
    "net/url"
    "io/ioutil"

    "github.com/11061055/goleveldb_server/manager"
)

var levelDBManager manager.LevelDBManager

func main()  {

    levelDBManager.Construct()
    levelDBManager.RefreshAsync()

    http.HandleFunc("/data",DataHandler)

    http.ListenAndServe("0.0.0.0:8880",nil)
}

func DataHandler(w http.ResponseWriter,r *http.Request)  {

    path := ""
    file := ""
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

    if val, ok := query["file"]; ok {
      file = val[0]
    }

    if val, ok := query["key"]; ok {
      key = val[0]
    }

    if len(path) == 0 || len(act) == 0 || len(file) == 0 || len(body) == 0 {
      return
    }

    db, err := levelDBManager.Open(file)
    if err != nil {
      w.Write([]byte("leveldb server inner error " + err.Error()))
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
       w.Write([]byte("leveldb server inner error " + err.Error()))
     } else {
       w.Write(ret)
     }
}
