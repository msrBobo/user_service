package postgres

import (
	pb "USER_SERVICE/genproto/user-service"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

// NewUserRepo ...
func NewUserRepo(db *sqlx.DB) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(user *pb.User) (*pb.User, error) {
	query := `INSERT INTO users (
		id, 
		name, 
		last_name
	) 
	VALUES ($1, $2, $3) 
	RETURNING 
		id, 
		name, 
		last_name
	`
	var respUser pb.User
	err := r.db.QueryRow(
		query,
		user.Id,
		user.Name,
		user.LastName,
	).Scan(
		&respUser.Id,
		&respUser.Name,
		&respUser.LastName,
	)
	if err != nil {
		return nil, err
	}
	return &respUser, nil
}

func (r *userRepo) Update(user *pb.User) (*pb.User, error) {
	query := `
	UPDATE 
		users 
	SET 
		name = $1, 
		last_name = $2 
	WHERE 
		id = $3 
	RETURNING
		id, 
		name, 
		last_name
	`
	var respUser pb.User
	err := r.db.QueryRow(
		query,
		user.Name,
		user.LastName,
		user.Id,
	).Scan(
		&respUser.Id,
		&respUser.Name,
		&respUser.LastName,
	)
	if err != nil {
		return nil, err
	}
	return &respUser, nil
}

func (r *userRepo) Delete(user *pb.UserRequest) (*pb.User, error) {
	query := `DELETE FROM users WHERE id = $1 RETURNING id, name, last_name`
	var respUser pb.User
	err := r.db.QueryRow(query, user.UserId).Scan(&respUser.Id, &respUser.Name, &respUser.LastName)
	if err != nil {
		return nil, err
	}
	return &respUser, nil
}

func (r *userRepo) Get(user *pb.UserRequest) (*pb.User, error) {
	query := `SELECT id, name, last_name FROM users WHERE id = $1`
	var respUser pb.User
	err := r.db.QueryRow(query, user.UserId).Scan(&respUser.Id, &respUser.Name, &respUser.LastName)
	if err != nil {
		return nil, err
	}
	return &respUser, nil
}

func (r *userRepo) GetAll(req *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	offset := req.Limit * (req.Page - 1)
	query := `SELECT id, name, last_name FROM users LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(query, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	var allUsers pb.GetAllUsersResponse
	for rows.Next() {
		var user pb.User
		if err := rows.Scan(&user.Id, &user.Name, &user.LastName); err != nil {
			return nil, err
		}
		allUsers.AllUsers = append(allUsers.AllUsers, &user)
	}
	return &allUsers, nil
}
