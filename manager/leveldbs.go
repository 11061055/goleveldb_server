package manager

import (
    "time"
    "sync"
    "strings"
    "github.com/syndtr/goleveldb/leveldb"
)

type LevelDBManager struct {
    sync.Mutex
    levelDBs   map[string]*LevelDB
}

func (m *LevelDBManager) Construct() {

    m.Lock()
    defer m.Unlock()

    m.levelDBs = make(map[string]*LevelDB)
}

func (m *LevelDBManager) RefreshAsync() {

    ticker := time.NewTicker(5 * time.Minute)

    go func() {

      for {

        select {

          case _ = <-ticker.C:
            go m.Refresh()

           //default:
           //  done <- true

            }
        }
    }()
}

func (m *LevelDBManager) Refresh() {

    m.Lock()
    defer m.Unlock()

    for _, db := range m.levelDBs {
      db.refresh()
    }
}

func (m *LevelDBManager) Open(table string) (*LevelDB, error) {

    m.Lock()
    defer m.Unlock()

    if db, ok := m.levelDBs[table]; !ok {
      db = new (LevelDB)
      db.setTable(table)
      m.levelDBs[table] = db
    }

    db := m.levelDBs[table]

    err := db.softOpen()

    if err != nil {
      return nil, err
    }

    return db, nil
}

func (m *LevelDBManager) Close(ldb *LevelDB) (err error) {

    m.Lock()
    defer m.Unlock()

    if ldb == nil {
      return
    }

    table := ldb.getTable()

    if _, ok := m.levelDBs[table]; !ok {
      return
    }

    err = ldb.softClose()

    return
}

type LevelDB struct {

    db         *leveldb.DB
    table      string
    refCount   int64 
    accessTime int64 

}

func (ldb *LevelDB) Get(key []byte) (value []byte, err error) {
    return ldb.db.Get(key, nil)
}

func (ldb *LevelDB) Put(key, val []byte) (err error) {
    return ldb.db.Put(key, val, nil)
}

func (ldb *LevelDB) Del(key []byte) (err error) {
    return ldb.db.Delete(key, nil)
}

func (ldb *LevelDB) getDB() *leveldb.DB {
   return ldb.db
}

func (ldb *LevelDB) setDB(db *leveldb.DB) {
   ldb.db = db
}

func (ldb *LevelDB) getTable() string {
   return ldb.table
}

func (ldb *LevelDB) setTable(table string) {
   if strings.Contains(table, "..") {
     panic("dangerous operation")
   }
   ldb.table = table
}

func (ldb *LevelDB) getRefCount() int64 {
   return ldb.refCount
}

func (ldb *LevelDB) setRefCount(count int64) {
   ldb.refCount = count
}

func (ldb *LevelDB) getAccessTime() int64 {
   return ldb.accessTime
}

func (ldb *LevelDB) setAccessTime(actime int64) {
   ldb.accessTime = actime
}

func (ldb *LevelDB) closeAble() bool {
   return (ldb.getDB() != nil) && (ldb.getRefCount() <= 0) && (time.Now().Unix() - 600 > ldb.getAccessTime())
}

func (ldb *LevelDB) close() (err error) {
   err = ldb.getDB().Close()
   //TODO

   ldb.setDB(nil)

   return
}

func (ldb *LevelDB) open() (err error) {
  db, err := leveldb.OpenFile("/data/logs/leveldb/db/" + ldb.getTable(), nil)
  if err != nil {
    return
    //TODO
  }

  ldb.setDB(db)
  return
}

func (ldb *LevelDB) hardOpen() (err error) {

  if db := ldb.getDB(); db == nil {
    if err = ldb.open(); err != nil {
      return
    }
    return
  }
  return
}

func (ldb *LevelDB) softClose() (err error) {

  ldb.setRefCount(ldb.getRefCount() - 1)

  if ldb.closeAble() {

    err = ldb.close() //TODO

    ldb.setAccessTime(0)
    ldb.setRefCount(0)
  }
  return
}

func (ldb *LevelDB) softOpen() (err error) {

  if db := ldb.getDB(); db == nil {

    if err = ldb.open(); err != nil {
      return
    }

    ldb.setAccessTime(time.Now().Unix())
    ldb.setRefCount(1)

    return
  }

  ldb.setAccessTime(time.Now().Unix())
  ldb.setRefCount(ldb.getRefCount() + 1)

  return
}

func (ldb *LevelDB) refresh() (err error) {

  if ldb.closeAble() {

    err = ldb.close() //TODO

    ldb.setAccessTime(0)
    ldb.setRefCount(0)
  }
  return
}