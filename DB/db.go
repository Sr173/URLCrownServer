package DB

import (
	"encoding/binary"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"time"
)

var db *leveldb.DB

var data = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func InitDB() (err error) {
	db_opt := &opt.Options{
		ErrorIfMissing: true,
	}

	db, err = leveldb.OpenFile("./task_db", db_opt)
	var temp = make([]byte, 8)
	if err != nil {
		db, err = leveldb.OpenFile("./task_db", nil)
		if err != nil {
			return
		}
		for i0 := 0; i0 < len(data); i0++ {
			for i1 := 0; i1 < len(data); i1++ {
				for i2 := 0; i2 < len(data); i2++ {
					for i3 := 0; i3 < len(data); i3++ {
						str := "5" + string(data[i0]) + string(data[i1]) + string(data[i2]) + string(data[i3])
						binary.BigEndian.PutUint64(temp, 0)
						db.Put([]byte(str), temp, nil)
					}
				}
			}
		}
	}

	return
}

func GetTask() (res string, err error) {
	iter := db.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		key := iter.Key()
		if int64(binary.BigEndian.Uint64(iter.Value())+60*10) < time.Now().Unix() {
			var temp = make([]byte, 8)
			binary.BigEndian.PutUint64(temp, uint64(time.Now().Unix()))
			db.Put(key, temp, nil)
			return string(key), nil
		}
	}
	err = iter.Error()
	if err == nil {
		err = errors.New("not found data")
	}

	return "", err
}
