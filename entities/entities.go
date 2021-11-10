package entities

import (
	"github.com/google/uuid"
)

type Guitar struct {
	Id 				string
	Brand 			string
	Model 			string
	Strings 		int
	Body 			string
	Neck 			string
	Finish 			string
	Scale			float32
}

type GTStruct struct {
	Guitars []Guitar
}

func (mv *Guitar) SetId() {
	mv.Id = uuid.New().String()
}