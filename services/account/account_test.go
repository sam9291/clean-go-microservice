package account

import "testing"

func TestAccount_NewService_WithoutConfiguration(t *testing.T) {

	_, err := NewService()

	if err != ErrInvalidServiceConfiguration {
		t.Fail()
	}
}

func TestAccount_GetUsers_ShouldReturnEmptyArray(t *testing.T) {

	service := setupSqlService(t)

	users, err := service.GetUsers()

	if err != nil {
		t.Error(err)
	}

	if len(users) != 0 {
		t.Fail()
	}
}

func TestAccount_CreateUser_ShouldCreateTheUser(t *testing.T) {

	service := setupMemoryService(t)

	err := service.CreateUsers("A", "B", "C")

	if err != nil {
		t.Error(err)
	}

	users, err := service.GetUsers()

	if err != nil {
		t.Error(err)
	}

	if len(users) != 1 {
		t.Fail()
	}
}

func setupMemoryService(t *testing.T) Service {
	service, err := NewService(
		WithSqlServerRepository("Data Source=.;Initial Catalog=goddd;Integrated Security=True;"),
	)

	if err != nil {
		t.Error(err)
	}
	return service
}

func setupSqlService(t *testing.T) Service {
	service, err := NewService(
		WithSqlServerRepository("Data Source=.;Initial Catalog=goddd;Integrated Security=True;"),
	)

	if err != nil {
		t.Error(err)
	}
	return service
}
