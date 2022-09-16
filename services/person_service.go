package services

import (
	"w8s/models"
	"w8s/repository"
)

type PersonSvc struct {
	PersonRepo repository.PersonRepo
}

func (s *PersonSvc) GetPersons(persons *[]models.Person) error {
	return s.PersonRepo.GetPersons(persons)
}

func (s *PersonSvc) GetPerson(id uint64, person *models.Person) error {
	return s.PersonRepo.GetPerson(id, person)
}

func (s *PersonSvc) CreatePerson(person *models.Person) error {
	return s.PersonRepo.CreatePerson(person)
}

func (s *PersonSvc) UpdatePerson(id uint64, person *models.Person) error {
	return s.PersonRepo.UpdatePerson(id, person)
}

func (s *PersonSvc) DeletePerson(id uint64) error {
	return s.PersonRepo.DeletePerson(id)
}
