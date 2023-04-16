package repository

import (
	"gorm.io/gorm"

	"mygram-api/models/domain"
)

type CommentRepository interface {
	Create(comment *domain.Comment) (err error)
	GetAll() (comments []domain.Comment, err error)
	GetOne(id uint) (comment domain.Comment, err error)
	Update(comment domain.Comment) (updatedComment domain.Comment, err error)
	Delete(id uint) (err error)
}

type CommentRepositoryDB struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &CommentRepositoryDB{DB: db}
}

func (commentRepository *CommentRepositoryDB) Create(comment *domain.Comment) (err error) {

	if err = commentRepository.DB.Create(&comment).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryDB) GetAll() (comments []domain.Comment, err error) {

	if err = commentRepository.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "email", "username")
	}).Preload("Photo", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "user_id", "title", "photo_url", "caption")
	}).Find(&comments).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryDB) GetOne(id uint) (comment domain.Comment, err error) {
    if err = commentRepository.DB.Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "email", "username")
    }).Preload("Photo", func(db *gorm.DB) *gorm.DB {
        return db.Select("id", "user_id", "title", "photo_url", "caption")
    }).First(&comment, id).Error; err != nil {
        return
    }
    return
}

func (commentRepository *CommentRepositoryDB) Update(comment domain.Comment) (updatedComment domain.Comment, err error) {

	if err = commentRepository.DB.First(&updatedComment, comment.ID).Error; err != nil {
		return
	}

	if err = commentRepository.DB.Model(&updatedComment).Updates(comment).Error; err != nil {
		return
	}

	return
}

func (commentRepository *CommentRepositoryDB) Delete(id uint) (err error) {

	if err = commentRepository.DB.First(&domain.Comment{}, id).Error; err != nil {
		return
	}

	if err = commentRepository.DB.Delete(&domain.Comment{}, id).Error; err != nil {
		return
	}

	return
}