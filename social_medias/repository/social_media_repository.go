package repository

import (
	"gorm.io/gorm"

	"mygram-api/models/domain"
)

type SocialMediaRepository interface {
	Create(socialMedia *domain.SocialMedia) (err error)
	GetAll() (socialMedias []domain.SocialMedia, err error)
	GetOne(id uint) (socialMedia domain.SocialMedia, err error)
	Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error)
	Delete(id uint) (err error)
}

type SocialMediaRepositoryDB struct {
	DB *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &SocialMediaRepositoryDB{DB: db}
}

func (socialMediaRepository *SocialMediaRepositoryDB) Create(socialMedia *domain.SocialMedia) (err error) {

	if err = socialMediaRepository.DB.Create(&socialMedia).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryDB) GetAll() (socialMedias []domain.SocialMedia, err error) {

	if err = socialMediaRepository.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Find(&socialMedias).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryDB) GetOne(id uint) (socialMedia domain.SocialMedia, err error) {

	if err = socialMediaRepository.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username")
	}).Find(&socialMedia, id).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryDB) Update(socialMedia domain.SocialMedia) (updatedSocialMedia domain.SocialMedia, err error) {

	if err = socialMediaRepository.DB.First(&updatedSocialMedia, socialMedia.ID).Error; err != nil {
		return
	}

	if err = socialMediaRepository.DB.Model(&updatedSocialMedia).Updates(socialMedia).Error; err != nil {
		return
	}

	return
}

func (socialMediaRepository *SocialMediaRepositoryDB) Delete(id uint) (err error) {

	if err = socialMediaRepository.DB.Delete(&domain.SocialMedia{}, id).Error; err != nil {
		return
	}

	return
}