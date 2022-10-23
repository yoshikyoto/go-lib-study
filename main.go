package main

import "fmt"

func main() {
	repo := NewRepositoryImpl()
	service := NewService(repo)
	service.Show()
}

// service
type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Show() {
	num := s.repo.Get()
	fmt.Println(num)
}

// domain repository
type Repository interface {
	Get() int
}

type RepositoryImpl struct {
}

func NewRepositoryImpl() RepositoryImpl {
	return RepositoryImpl{}
}

func (r RepositoryImpl) Get() int {
	return 1
}
