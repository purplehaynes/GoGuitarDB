package repo

import (
	"GuitarDB/entities"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Repo struct {
	Filename string
}

func NewRepository(fn string) Repo {
	return Repo {
		Filename: fn,
	}
}

func (r Repo) CreateNewGuitar(axe entities.Guitar) error {

	gt := entities.GTStruct{}

	output, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(output, &gt)
	if err != nil {
		return err
	}

	for _, v := range gt.Guitars { 
		if v.Model == axe.Model {
		return errors.New("guitar already exists in database")			
		}
	}

	gt.Guitars = append(gt.Guitars, axe)

	input, err := json.MarshalIndent(gt, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, input, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) ReadAll() (*entities.GTStruct, error) {
	gt := entities.GTStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return nil, errors.New("data field empty")
	}

	err = json.Unmarshal(file, &gt)

	return &gt, nil
}

func (r Repo) GetGuitarId(id string) (*entities.Guitar, error) {

	gt := entities.GTStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(file, &gt)
	if err != nil {
		return nil, err
	}

	for _, v := range gt.Guitars {
		if v.Id == id {
			return &v, nil
		}
	}

	return nil, errors.New("guitar not found")
}

func (r Repo) DeleteGuitarId(id string) error {
	gt := entities.GTStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &gt)
	if err != nil {
		return err
	}

	for i, v := range gt.Guitars {
		if id == v.Id {
			gt.Guitars = append(gt.Guitars[:i], gt.Guitars[i + 1:]...)
		}
	}

	output, err := json.MarshalIndent(gt, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, output, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (r Repo) UpdateGuitarInfo(id string, axe entities.Guitar) error {
	gt := entities.GTStruct{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &gt)
	if err != nil {
		return err
	}

	// ranges through our existing guitarDB. In Postman, if the Id in the DB matches the info 
	// being sent back, anything new in the fields (i) of the movieDB will be updated
	for i, v := range gt.Guitars {
		if v.Id == id {
			gt.Guitars[i] = axe
		}
	}

	output, err := json.MarshalIndent(gt, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, output, 0644)
	if err != nil {
		return err
	}
	return nil
}