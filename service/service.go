package service

import (
	"GuitarDB/entities"
	"errors"
	"github.com/google/uuid"
)

type Repo interface {
	CreateNewGuitar(axe entities.Guitar) error
	ReadAll() (*entities.GTStruct, error)
	GetGuitarId(id string) (*entities.Guitar, error)
	DeleteGuitarId(id string) error
	UpdateGuitarInfo(id string, axe entities.Guitar) error
}

type Service struct {
	Repo Repo
}

func NewService(r Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateNewGuitar(axe entities.Guitar) error {
	axe.Id = uuid.New().String()

	// if axe.Rating >= 0 && axe.Rating <= 10 {
		err := s.Repo.CreateNewGuitar(axe)
		if err != nil {
			return err
		// }
	}
	return nil
}

func (s Service) ReadAll() (*entities.GTStruct, error) {
	view, err := s.Repo.ReadAll()
	if err != nil {
		return nil, errors.New("cannot locate data")
	}
	return view, nil
}

func (s Service) GetByGuitarId(id string) (*entities.Guitar, error) {
	searchRequest, err := s.Repo.GetGuitarId(id)
	if err != nil {
		return nil, err
	}
	return searchRequest, nil
}

func (s Service) DeleteGuitarId(id string) error {
	err := s.Repo.DeleteGuitarId(id)
	if err != nil {
		return errors.New("guitar does not exist")
	}
	return err
}

func (s Service) UpdateGuitarInfo(id string, axe entities.Guitar) error {
	if id != axe.Id {
		return errors.New("id is mismatched")
	}

	err := s.Repo.UpdateGuitarInfo(id, axe)
	if err != nil {
		return err
	}
	return nil
}