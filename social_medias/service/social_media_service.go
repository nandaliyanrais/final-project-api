package service

import (
	"mygram-api/models/domain"
	"mygram-api/social_medias/repository"
)

type SocialMediaService interface {
	Create(socialMedia *domain.SocialMedia) (err error)
	GetAll() (socialMedias []domain.SocialMedia, err error)
	GetOne(id uint) (socialMedia domain.SocialMedia, err error)
	Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error)
	Delete(id uint) (err error)
}

type SocialMediaServiceRepository struct {
	SocialMediaRepository repository.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepository repository.SocialMediaRepository) SocialMediaService {
	return &SocialMediaServiceRepository{SocialMediaRepository: socialMediaRepository}
}

func (socialMediaService *SocialMediaServiceRepository) Create(socialMedia *domain.SocialMedia) (err error) {

	if err = socialMediaService.SocialMediaRepository.Create(socialMedia); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceRepository) GetAll() (socialMedias []domain.SocialMedia, err error) {

	if socialMedias, err = socialMediaService.SocialMediaRepository.GetAll(); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceRepository) GetOne(id uint) (socialMedia domain.SocialMedia, err error) {

	if socialMedia, err = socialMediaService.SocialMediaRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceRepository) Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error) {

	if updatedSocialMedia, err = socialMediaService.SocialMediaRepository.Update(socialMedia); err != nil {
		return
	}

	return
}

func (socialMediaService *SocialMediaServiceRepository) Delete(id uint) (err error) {

	if err = socialMediaService.SocialMediaRepository.Delete(id); err != nil {
		return
	}

	return
}