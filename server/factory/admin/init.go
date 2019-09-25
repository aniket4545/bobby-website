package admin

import (
	config "bobby-website/server/configurations"
	"bytes"
	"encoding/gob"
	"os"

	"go.etcd.io/bbolt"
)

//create user

var (
	path   = config.PATH + "/admin.db"
	bucket = "root"
	key    = "user"
)

//first load user if we get empty user then create user write it in DB

//ADMIN is the user object
var ADMIN *Admin

func init() {
	load()
	if ADMIN == nil {
		os.MkdirAll(config.PATH, 0777)
		ADMIN = create()
	}
}

func load() {
	var temp = new(Admin)
	if db, _ := bbolt.Open(path, 0660, nil); db != nil {
		defer db.Close()
		db.View(func(tx *bbolt.Tx) error {
			adminBucket := tx.Bucket([]byte(bucket))
			if user := adminBucket.Get([]byte(key)); user != nil {
				dec := gob.NewDecoder(bytes.NewReader(user))
				dec.Decode(temp)
				ADMIN = temp
			}
			return nil
		})
	}
}

func create() *Admin {
	var temp = new(Admin)
	if db, _ := bbolt.Open(path, 0660, nil); db != nil {
		defer db.Close()
		db.Update(func(tx *bbolt.Tx) error {
			adminBucket, _ := tx.CreateBucketIfNotExists([]byte(bucket))
			var buff = new(bytes.Buffer)
			encoder := gob.NewEncoder(buff)
			temp = &Admin{
				Name:     "admin",
				Email:    "admin@admin.com",
				Password: "admin@filmwork.com",
			}
			encoder.Encode(temp)
			adminBucket.Put([]byte(key), buff.Bytes())
			return nil
		})
	}
	return temp
}
