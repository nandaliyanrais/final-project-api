package service

import (
	"mygram-api/models/domain"
	"mygram-api/comments/repository"
)

type CommentService interface {
	Create(comment *domain.Comment) (err error)
	GetAll() (comments []domain.Comment, err error)
	GetOne(id uint) (comment domain.Comment, err error)
	Update(comment domain.Comment) (updatedComment domain.Comment, err error)
	Delete(id uint) (err error)
}

type CommentServiceRepository struct {
	CommentRepository repository.CommentRepository
}

func NewCommentService(commentRepository repository.CommentRepository) CommentService {
	return &CommentServiceRepository{CommentRepository: commentRepository}
}

func (commentService *CommentServiceRepository) Create(comment *domain.Comment) (err error) {

	if err = commentService.CommentRepository.Create(comment); err != nil {
		return
	}
	return
}

func (commentService *CommentServiceRepository) GetAll() (comments []domain.Comment, err error) {

	if comments, err = commentService.CommentRepository.GetAll(); err != nil {
		return
	}

	return
}

func (commentService *CommentServiceRepository) GetOne(id uint) (comment domain.Comment, err error) {
	if comment, err = commentService.CommentRepository.GetOne(id); err != nil {
		return
	}

	return
}

func (commentService *CommentServiceRepository) Update(comment domain.Comment) (updatedComment domain.Comment, err error) {

	if updatedComment, err = commentService.CommentRepository.Update(comment); err != nil {
		return
	}

	return
}

func (commentService *CommentServiceRepository) Delete(id uint) (err error) {

	if err = commentService.CommentRepository.Delete(id); err != nil {
		return
	}

	return
}