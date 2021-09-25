package usecase

import "wizeline/repository"

func GetAllUsers() (string ,error) {
	return repository.GetUsersFromCSV("data/users.csv")
}
