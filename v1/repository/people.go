package repository

import "post-person/v1/model"

type peopleRepository struct{}

func (r *peopleRepository) Save(m model.Person) error {
	return nil
}

func NewPeopleRepository() *peopleRepository {
	return &peopleRepository{}
}
