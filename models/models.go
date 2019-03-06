package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	//User Struct
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty "`
	}

	//Task Struct
	Task struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   string        `json:"createdon,omitempty"`
		Due         string        `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tag         []string      `json:"tags,omitempty"`
	}

	//TaskNote Struct
	TaskNote struct {
		ID          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskID      bson.ObjectId `json:"taskid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
)
