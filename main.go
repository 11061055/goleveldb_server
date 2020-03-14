package main

import (

    "errors"
    "io/ioutil"
    "net/http"
    "net/url"

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

    table := ""
    act   := ""
    key   := ""

    parsedUrl, err := url.Parse(r.URL.String())
    if  err != nil {
      return
    }

    query, err := url.ParseQuery(parsedUrl.RawQuery)
    if  err != nil {
      return
    }

    if val, ok := query["table"]; ok {
      table = val[0]
    }

    if val, ok := query["act"]; ok {
      act = val[0]
    }

    if val, ok := query["key"]; ok {
      key = val[0]
    }

    if (len(table) == 0)  ||
       (len(act) == 0)    ||
       (len(key) == 0)    {

      return
    }

    db, err := levelDBManager.Open(table)
    if err != nil {
      w.Write([]byte("leveldb open table error " + err.Error()))
      return
    }
    defer levelDBManager.Close(db)

    var ret []byte = []byte("success")
    switch act {

      case "put":

        if  body, err := ioutil.ReadAll(r.Body); err == nil {
          err    = db.Put([]byte(key), body)
        }

      case "get":
        ret, err = db.Get([]byte(key))

      case "del":
        err      = db.Del([]byte(key))

      default:
        err      = errors.New("invalid act")

     }

     if err != nil {
       w.Write([]byte("leveldb error " + err.Error()))
     } else {
       w.Write(ret)
     }
}
