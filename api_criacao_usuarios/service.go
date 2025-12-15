package apicriacaousuarios

import "sync"

type UserService struct {
	mu     sync.Mutex
	users  map[int]User
	nextID int
}

func NovoService() *UserService {
	return &UserService{
		users:  make(map[int]User),
		nextID: 1,
	}
}

func (service *UserService) CreateUser(name, email string) User {
	service.mu.Lock()
	defer service.mu.Unlock()

	user := User{
		ID:    service.nextID,
		Name:  name,
		Email: email,
	}

	service.users[service.nextID] = user
	service.nextID++
	return user

}

func (service *UserService) GetUser(id int) (User, bool) {
	service.mu.Lock()
	defer service.mu.Unlock()

	user, ok := service.users[id]
	return user, ok
}

func (service *UserService) ListUsers() []User {
	service.mu.Lock()
	service.mu.Unlock()

	var result []User
	for _, user := range service.users {
		result = append(result, user)
	}
	return result
}

func (service *UserService) UpdateUser(id int, name, email string) (User, bool) {

	service.mu.Lock()
	defer service.mu.Unlock()

	user, ok := service.users[id]
	if !ok {
		return User{}, false
	}

	user.Name = name
	user.Email = email
	service.users[id] = user

	return user, true

}

func (service *UserService) DeletaUser(id int) bool {
	service.mu.Lock()
	service.mu.Unlock()

	_, ok := service.users[id]

	if !ok {
		return false
	}

	delete(service.users, id)
	return true

}
