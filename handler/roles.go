package handler

import (
	"beverages-cli/entity"
	"context"
	"database/sql"
	"time"
)

func GetAllRoles(db *sql.DB) ([]entity.Roles, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := db.QueryContext(ctx, `
		SELECT *
		FROM Roles r 
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []entity.Roles
	for rows.Next() {
		var p entity.Roles
		if err = rows.Scan(&p.RoleID, &p.RoleName); err != nil {
			return nil, err
		}
		roles = append(roles, p)
	}

	return roles, nil
}

