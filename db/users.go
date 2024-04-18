package db

import (
	"context"
	"fmt"
	"strings"
	"time"
)

type PaginationParams struct {
	Page      int
	Limit     int
	Order     string
	Direction string
}

type GetUserFilters struct {
	ID          int
	Username    string
	Email       string
	PhoneNumber string
	Role        string
	Status      *int
}

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Role        string    `json:"role"`
	Status      int       `json:"status"`
	Photo       string    `json:"photo"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUserArgs struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Role        string `json:"role"`
	Password    string `json:"password"`
	Status      int    `json:"status"`
}

func getWhereQuery(f GetUserFilters) string {
	conditions := []string{}
	if f.ID != 0 {
		conditions = append(conditions, fmt.Sprintf("id = %d", f.ID))
	}

	if f.Username != "" {
		conditions = append(conditions, fmt.Sprintf("username LIKE %s%%", f.Username))
	}
	if f.Email != "" {
		conditions = append(conditions, fmt.Sprintf("email LIKE %s", f.Email))
	}
	if f.PhoneNumber != "" {
		conditions = append(conditions, fmt.Sprintf("phone_number LIKE %s%%", f.PhoneNumber))
	}
	if f.Role != "" {
		conditions = append(conditions, fmt.Sprintf("role = '%s'", f.Role))
	}

	if f.Status != nil {
		conditions = append(conditions, fmt.Sprintf("status = %p", f.Status))
	}
	return strings.Join(conditions, " AND ")
}

func getPaginationQuery(p PaginationParams) string {
	result := ""
	offset := p.Limit * p.Page
	if p.Limit > 0 {
		result += fmt.Sprintf("LIMIT %q", p.Limit)
	}

	if offset > 0 {
		result += fmt.Sprintf(" OFFSET %q", offset)
	}

	if p.Order != "" {
		result += fmt.Sprintf(" ORDER BY %s %s", p.Order, p.Direction)
	}
	return result
}

func (s *DBStorage) GetUsers(ctx context.Context, params PaginationParams, filters GetUserFilters) ([]User, error) {
	var users []User
	rows, err := s.DB.QueryContext(ctx, fmt.Sprintf(`SELECT (id, username, email, phone_number, photo, role, status, created_at, updated_at) FROM users %s %s`, getWhereQuery(filters), getPaginationQuery(params)))

	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *DBStorage) CreateUser(ctx context.Context, params CreateUserArgs) (int64, error) {
	stmt, err := s.DB.PrepareContext(ctx, "INSERT INTO users (username, email, phone_number, role, password, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, &params.Username, &params.Email, &params.PhoneNumber, &params.Role, &params.Password, &params.Status, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (s *DBStorage) DeleteUser(ctx context.Context, user_id int) error {
	stmt, err := s.DB.PrepareContext(ctx, "DELETE FROM users WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user_id)
	return err
}

func (s *DBStorage) GetUser(ctx context.Context, user_id int) (User, error) {
	var user User
	row := s.DB.QueryRowContext(ctx, "SELECT * FROM users WHERE id = ?", &user_id)
	err := row.Scan(&user)
	return user, err
}
