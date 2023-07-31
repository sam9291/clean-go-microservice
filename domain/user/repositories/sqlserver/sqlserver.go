package sqlserver

import (
	"database/sql"
	"stagingmanager/domain/user"

	"github.com/google/uuid"
)

type sqlServerRepository struct {
	connectionString string
}

type userRead struct {
	id        string
	email     string
	firstName string
	lastName  string
}

func (u *userRead) ToAggregate() (user.User, error) {
	uuid, err := uuid.Parse(u.id)
	if err != nil {
		return user.User{}, err
	}
	return user.NewUserWithId(uuid, u.email, u.firstName, u.lastName)
}

// CreateUser implements user.Repository.
func (r sqlServerRepository) CreateUser(user user.User) (uuid.UUID, error) {
	// insertQuery := `
	// 	insert into users ()
	// `
	// _, err := r.db.Exec(insertQuery,
	// 	user.ID,
	// 	user.FirstName,
	// 	user.LastName,
	// 	user.Email,
	// )
	// if err != nil {
	// 	return uuid.UUID{}, err
	// }
	return uuid.New(), nil
}

// GetUsers implements user.UserRepository.
func (r sqlServerRepository) GetUsers() ([]user.User, error) {
	selectQuery := `
		select * from dbo.users
	`

	db, err := sql.Open("mssql", r.connectionString)

	if err != nil {
		return nil, err
	}
	defer db.Close()
	rows, err := db.Query(selectQuery)

	if err != nil {
		return nil, err
	}

	var users []user.User

	for rows.Next() {
		var user userRead

		if err := rows.Scan(&user.id, &user.email, &user.firstName, &user.lastName); err != nil {
			return users, err
		}
		aggregate, err := user.ToAggregate()

		if err != nil {
			return users, err
		}

		users = append(users, aggregate)
	}

	return users, nil
}

func NewSqlServerRepository(connectionString string) user.Repository {
	return &sqlServerRepository{
		connectionString: connectionString,
	}
}
