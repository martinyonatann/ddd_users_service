package repository

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/martinyonathann/users-service/internal/auth"
	"github.com/martinyonathann/users-service/internal/models"
	"github.com/martinyonathann/users-service/pkg/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

type authRepo struct {
	db *sqlx.DB
}

const (
	createUserQuery = `INSERT INTO users (first_name, last_name, email, password, role, about, avatar, phone_number, address,
	               		city, gender, postcode, birthday, created_at, updated_at, login_date)
						VALUES ($1, $2, $3, $4, COALESCE(NULLIF($5, ''), 'user'), $6, $7, $8, $9, $10, $11, $12, $13, now(), now(), now()) 
						RETURNING *`

	updateUserQuery = `UPDATE users 
						SET first_name = COALESCE(NULLIF($1, ''), first_name),
						    last_name = COALESCE(NULLIF($2, ''), last_name),
						    email = COALESCE(NULLIF($3, ''), email),
						    role = COALESCE(NULLIF($4, ''), role),
						    about = COALESCE(NULLIF($5, ''), about),
						    avatar = COALESCE(NULLIF($6, ''), avatar),
						    phone_number = COALESCE(NULLIF($7, ''), phone_number),
						    address = COALESCE(NULLIF($8, ''), address),
						    city = COALESCE(NULLIF($9, ''), city),
						    gender = COALESCE(NULLIF($10, ''), gender),
						    postcode = COALESCE(NULLIF($11, 0), postcode),
						    birthday = COALESCE(NULLIF($12, '')::date, birthday),
						    updated_at = now()
						WHERE user_id = $13
						RETURNING *
						`

	deleteUserQuery = `DELETE FROM users WHERE user_id = $1`

	getUserQuery = `SELECT user_id, first_name, last_name, email, role, about, avatar, phone_number, 
       				 address, city, gender, postcode, birthday, created_at, updated_at, login_date  
					 FROM users 
					 WHERE user_id = $1`

	getTotalCount = `SELECT COUNT(user_id) FROM users 
						WHERE first_name ILIKE '%' || $1 || '%' or last_name ILIKE '%' || $1 || '%'`

	findUsers = `SELECT user_id, first_name, last_name, email, role, about, avatar, phone_number, address,
	              city, gender, postcode, birthday, created_at, updated_at, login_date 
				  FROM users 
				  WHERE first_name ILIKE '%' || $1 || '%' or last_name ILIKE '%' || $1 || '%'
				  ORDER BY first_name, last_name
				  OFFSET $2 LIMIT $3
				  `

	getTotal = `SELECT COUNT(user_id) FROM users`

	getUsers = `SELECT user_id, first_name, last_name, email, role, about, avatar, phone_number, 
       			 address, city, gender, postcode, birthday, created_at, updated_at, login_date
				 FROM users 
				 ORDER BY COALESCE(NULLIF($1, ''), first_name) OFFSET $2 LIMIT $3`

	findUserByEmail = `SELECT user_id, first_name, last_name, email, role, about, avatar, phone_number, 
       			 		address, city, gender, postcode, birthday, created_at, updated_at, login_date, password
				 		FROM users 
				 		WHERE email = $1`
)

func NewAuthRepository(db *sqlx.DB) auth.Repository {
	return &authRepo{db: db}
}

func (r *authRepo) Register(
	ctx context.Context,
	user *models.User,
) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Register")
	defer span.Finish()

	userData := &models.User{}
	if err := r.db.QueryRowContext(
		ctx, createUserQuery,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.About,
		&user.PhoneNumber,
		&user.Address,
		&user.City,
		&user.Gender,
		&user.Postcode,
		&user.Birthday,
	).Scan(userData); err != nil {
		return nil, errors.Wrap(err, "authRepo.Register.StructScan")
	}

	return userData, nil
}

func (r *authRepo) Update(ctx context.Context, user *models.User) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Update")
	defer span.Finish()

	userData := &models.User{}

	if err := r.db.GetContext(
		ctx,
		userData,
		updateUserQuery,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Role,
		&user.About,
		&user.Avatar,
		&user.PhoneNumber,
		&user.Address,
		&user.City,
		&user.Gender,
		&user.Postcode,
		&user.Birthday,
		&user.UserID,
	); err != nil {
		return nil, errors.Wrap(err, "authRepo.Update.GetContext")
	}

	return userData, nil
}
func (r *authRepo) Delete(ctx context.Context, userID uuid.UUID) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.Delete")
	span.Finish()

	result, err := r.db.ExecContext(ctx, deleteUserQuery, userID)
	if err != nil {
		return errors.WithMessage(err, "authRepo Delete ExecContext")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "authRepo.Delete.RowsAffected")
	}

	if rowAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "authRepo.Delete.rowAffected")
	}

	return nil
}
func (r *authRepo) GetByID(ctx context.Context, userID uuid.UUID) (*models.User, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "authRepo.GetByID")
	span.Finish()
	userData := &models.User{}

	if err := r.db.QueryRowContext(
		ctx,
		getUserQuery,
		userID,
	).Scan(userData); err != nil {
		return nil, errors.Wrap(err, "authRepo.GetByID.QueryRowxContext")
	}

	return userData, nil
}
func (r *authRepo) FindByName(ctx context.Context, name string, query *utils.PaginationQuery) (*models.UsersList, error)
func (r *authRepo) FindByEmail(ctx context.Context, user *models.User) (*models.User, error)
func (r *authRepo) GetUsers(ctx context.Context, pq *utils.PaginationQuery) (*models.UsersList, error)
