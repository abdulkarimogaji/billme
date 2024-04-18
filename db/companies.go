package db

import (
	"context"
	"fmt"
	"time"
)

type GetCompaniesFilters struct {
	ID          int
	ProductID   int
	Name        string
	Email       string
	PhoneNumber string
	Website     string
	AgentID     int
	Status      *int
}

type Company struct {
	ID          int       `json:"id"`
	ProductID   string    `json:"product_id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	Website     string    `json:"website"`
	Logo        string    `json:"logo"`
	AgentID     int       `json:"agent_id"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CompanyMember struct {
	ID int `json:"id"`

	CompanyName        string `json:"company_name"`
	CompanyEmail       string `json:"company_email"`
	CompanyPhoneNumber string `json:"company_phone_number"`
	CompanyWebsite     string `json:"company_website"`
	CompanyLogo        string `json:"company_logo"`
	CompanyID          int    `json:"company_id"`

	Username        string `json:"username"`
	UserEmail       string `json:"user_email"`
	UserPhoneNumber string `json:"user_phone_number"`
	UserPhoto       string `json:"user_photo"`
	UserID          int    `json:"user_id"`

	Role      string    `json:"role"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCompanyMembersFilters struct {
	ID int

	CompanyID          int
	CompanyName        string
	CompanyEmail       string
	CompanyPhoneNumber string

	UserID          int
	Username        string
	UserEmail       string
	UserPhoneNumber string

	Role   string
	Status *int
}

type CreateCompanyArgs struct {
	ProductID   int    `json:"product_id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Website     string `json:"website"`
	Logo        string `json:"logo"`
	AgentID     int    `json:"agent_id"`
	Status      int    `json:"status"`
}

func (s *DBStorage) GetCompanies(ctx context.Context, params PaginationParams, filters GetCompaniesFilters) ([]Company, error) {
	var companies []Company
	rows, err := s.DB.QueryContext(ctx, fmt.Sprintf(`SELECT (id, product_id, name, email, phone_number, website, logo, agent_id, status, created_at, updated_at) FROM companies %s %s`, "", getPaginationQuery(params)))

	if err != nil {
		return companies, err
	}

	for rows.Next() {
		var c Company
		err = rows.Scan(&c)
		if err != nil {
			return companies, err
		}
		companies = append(companies, c)
	}

	return companies, nil
}

func (s *DBStorage) CreateCompany(ctx context.Context, params CreateCompanyArgs) (int64, error) {
	stmt, err := s.DB.PrepareContext(ctx, "INSERT INTO companies (product_id, name, email, phone_number, website, logo, agent_id, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(ctx, &params.ProductID, &params.Name, &params.Email, &params.PhoneNumber, &params.Website, &params.Logo, &params.AgentID, &params.Status, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (s *DBStorage) DeleteCompany(ctx context.Context, user_id int) error {
	stmt, err := s.DB.PrepareContext(ctx, "DELETE FROM companies WHERE id = ?")
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, user_id)
	return err
}

func (s *DBStorage) GetCompanyMembers(ctx context.Context, params PaginationParams, filters GetCompanyMembersFilters) ([]CompanyMember, error) {
	var members []CompanyMember

	rows, err := s.DB.QueryContext(ctx, fmt.Sprintf(`
	SELECT 
		cm.id,
		cm.role,
		cm.status,
		cm.created_at,
		cm.updated_at,
		cm.company_id,
		cm.user_id,

		c.name AS company_name,
		c.email AS company_email,
		c.phone_number AS company_phone_number,
		c.website AS company_website,
		c.logo AS company_logo,

		u.username,
		u.email AS user_email,
		u.phone_number AS user_phone_number,
		u.photo AS user_photo

	FROM 
		company_member cm
	LEFT JOIN 
		companies c ON c.id = cm.company_id
	LEFT JOIN 
		users u ON u.id = cm.user_id
%s %s`, "", getPaginationQuery(params)))

	if err != nil {
		return members, err
	}

	for rows.Next() {
		var m CompanyMember
		err = rows.Scan(&m)
		if err != nil {
			return members, err
		}
		members = append(members, m)
	}

	return members, nil
}
