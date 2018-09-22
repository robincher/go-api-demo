package dao

import (
	"log"

	model "github.com/robincher/go-api-example/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MoviesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

func (m *MoviesDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *MoviesDAO) FindAll() ([]model.Movie, error) {
	var movies []model.Movie
	err := db.C(COLLECTION).Find(bson.M{}).All(&movies)
	return movies, err
}

func (m *MoviesDAO) FindById(id string) (model.Movie, error) {
	var movie model.Movie
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&movie)
	return movie, err
}

func (m *MoviesDAO) Insert(movie model.Movie) error {
	err := db.C(COLLECTION).Insert(&movie)
	return err
}

func (m *MoviesDAO) Update(movie model.Movie) error {
	err := db.C(COLLECTION).UpdateId(movie.ID, &movie)
	return err
}

func (m *MoviesDAO) Delete(movieId string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(movieId))
	if err != nil {
		return err
	}
	return nil
}
