package repository

import (
	"w8s/models"

	"gorm.io/gorm"
)

type PersonRepo struct {
	Conn *gorm.DB
}

func (r *PersonRepo) GetPersons(persons *[]models.Person) error {
	return r.Conn.Find(persons).Error
}

func (r *PersonRepo) GetPerson(id uint64, person *models.Person) error {
	return r.Conn.Where("id = ?", id).First(&person).Error
}

func (r *PersonRepo) CreatePerson(dataPerson *models.Person) error {
	return r.Conn.Create(dataPerson).Error
}

func (r *PersonRepo) UpdatePerson(id uint64, dataPerson *models.Person) error {
	return r.Conn.Model(&models.Person{}).Where("id = ?", id).Updates(dataPerson).Error
}

func (r *PersonRepo) DeletePerson(id uint64) error {
	return r.Conn.Where("id = ?", id).Delete(&models.Person{}).Error
}
