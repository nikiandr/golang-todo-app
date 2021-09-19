package service

import "github.com/nikiandr/golang-todo-app/pkg/repository"

type Authorization interface {
}

type List interface {
}

type Item interface {
}

type Service struct {
	Authorization
	List
	Item
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
