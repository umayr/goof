package util

import (
	"os"
	"gopkg.in/mgo.v2"
)

const (
	uri = "mongodb://localhost/phdb-local"
	name = "phdb-local"
)
type Db struct {
	Uri     string
	Session *mgo.Session
	Name    string
}

func NewDb() (*Db) {
	var u string = func(uri string) (string) {
		if os.Getenv("MONGO_URI") != "" {
			return string(os.Getenv("MONGO_URI"))
		} else {
			return uri
		}
	}(uri)

	var s *mgo.Session = func() (*mgo.Session) {
		s, err := mgo.Dial(u)
		if err != nil {
			panic(err)
		}
		return s
	}()

	var n string = func(name string) (string) {
		if os.Getenv("DB_NAME") != "" {
			return string(os.Getenv("DB_NAME"))
		} else {
			return name
		}
	}(name)

	return &Db{
		Uri: u,
		Session: s,
		Name: n,
	}
}

func (d *Db) Collection(name string) (collection *mgo.Collection) {
	return d.Session.DB(d.Name).C(name)
}

func (d *Db) Posts() (collection *mgo.Collection) {
	return d.Collection("posts")
}

func (d *Db) Close() {
	d.Session.Close()
}

