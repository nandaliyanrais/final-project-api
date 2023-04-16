package service

import (
	"mygram-api/models/domain"
	"mygram-api/photos/repository"
)

type PhotoService interface {
	Create(photo *domain.Photo) (err error)
	GetAll() (photos []domain.Photo, err error)
	GetOne(id uint) (photo domain.Photo, err error)
	Update(photo domain.Photo) (updatedPhoto domain.Photo, err error)
	Delete(id uint) (err error)
}

type PhotoServiceRepository struct {
	PhotoRepository repository.PhotoRepository
}

func NewPhotoService(photoRepository repository.PhotoRepository) PhotoService {
	return &PhotoServiceRepository{PhotoRepository: photoRepository}
}

func (photoService *PhotoServiceRepository) Create(photo *domain.Photo) (err error) {

	if err = photoService.PhotoRepository.Create(photo); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceRepository) GetAll() (photos []domain.Photo, err error) {

	if photos, err = photoService.PhotoRepository.GetAll(); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceRepository) GetOne(id uint) (photos domain.Photo, err error) {

	if photos, err = photoService.PhotoRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceRepository) Update(photo domain.Photo) (updatedPhoto domain.Photo, err error) {

	if updatedPhoto, err = photoService.PhotoRepository.Update(photo); err != nil {
		return
	}

	return
}

func (photoService *PhotoServiceRepository) Delete(id uint) (err error) {

	if err = photoService.PhotoRepository.Delete(id); err != nil {
		return
	}

	return
}