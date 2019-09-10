package repository

import (
	"cleanbase/app/user"
	"cleanbase/app/user/models"
	"context"
	"database/sql"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"time"
)

type mysqlUserRepository struct {
	DB *sql.DB
}

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

func (h *mysqlUserRepository) FetchUsers(ctx context.Context, cursor string, num int64) ([]*models.User, string, error) {
	query := `SELECT id_user, name FROM users WHERE created_at > ? ORDER BY name LIMIT ?`

	decodedCursor, err := DecodeCursor(cursor)
	if err != nil && cursor != "" {
		return nil, "", err
	}

	res, err := h.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	nextCursor := ""
	if len(res) == int(num) {
		nextCursor = EncodeCursor(time.Now())
	}

	return res, nextCursor, err
}

func (h *mysqlUserRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := h.DB.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error("err")
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.User, 0)
	for rows.Next() {
		if err := ctx.Err(); err != nil {
			return nil, err
		}
		t := new(models.User)
		err = rows.Scan(
			&t.Id,
			&t.Name,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (time.Time, error) {
	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(byt)
	t, err := time.Parse(timeFormat, timeString)

	return t, err
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

func (h *mysqlUserRepository) FetchUser(c context.Context, user int64) (*models.User, error) {
	a := &models.User{}
	query := `SELECT name FROM users WHERE id_user=?`
	row := h.DB.QueryRow(query, user)
	err := row.Scan(
		&a.Name,
	)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (h *mysqlUserRepository) Store(c context.Context, post *models.User) error {
	query := `INSERT users SET name=?`
	stmt, err := h.DB.PrepareContext(c, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(c, post.Name)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	post.Id = lastID
	return nil
}

func NewMysqlUserRepository(db *sql.DB) user.Repository{
	return &mysqlUserRepository{
		DB:db,
	}
}