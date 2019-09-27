package admin

import (
	config "bobby-website/server/configurations"
	"bytes"
	"encoding/gob"
	"fmt"

	"go.etcd.io/bbolt"
)

//create user

var (
	path   = config.PATH + "/admin.db"
	bucket = "admin"
	key    = "user"
)

//first load user if we get empty user then create user write it in DB

//ADMIN is the user object
var ADMIN *admin

func init() {
	if ADMIN = load(); ADMIN == nil {
		fmt.Println("creating admin")
		ADMIN = create()
	}
	fmt.Println("admin created:::::", ADMIN)
}

func load() *admin {
	var adminUser = new(admin)
	if db, _ := bbolt.Open(path, 0660, nil); db != nil {
		db.View(func(tx *bbolt.Tx) error {
			adminBucket := tx.Bucket([]byte(bucket))
			if user := adminBucket.Get([]byte(key)); user != nil {
				dec := gob.NewDecoder(bytes.NewReader(user))
				err := dec.Decode(&adminUser)
				return err
			}
			return nil
		})
	}
	return adminUser
}

func create() *admin {
	var temp = new(admin)
	if db, _ := bbolt.Open(path, 0660, nil); db != nil {
		db.Update(func(tx *bbolt.Tx) error {
			adminBucket, _ := tx.CreateBucketIfNotExists([]byte(bucket))
			var buff = new(bytes.Buffer)
			encoder := gob.NewEncoder(buff)
			encoder.Encode(temp)
			adminBucket.Put([]byte(key), buff.Bytes())
			return nil
		})
	}
	return temp
}
