package repositories

import (
	"database/sql"

	"github.com/eviccari/multithread-test-go/internal/app/domain/dtos"
)

type GreatestUserMySQLRepository struct {
	db *sql.DB
}

func NewGreatestUserMySQLRepository(db *sql.DB) GreatestUserMySQLRepository {
	return GreatestUserMySQLRepository{
		db: db,
	}
}

func (gumr GreatestUserMySQLRepository) Create(dto dtos.GreatestUserDTO) (err error) {
	query := `insert into greatest_users.users (
		         id, 
				 legacy_login, 
				 legacy_id, 
				 legacy_node_id, 
				 legacy_url, 
				 legacy_html_url, 
				 new_email, 
				 created_at
			) values (
				?, ?, ?, ?, ?, ?, ?, ?
			)`

	_, err = gumr.db.Exec(
		query,
		dto.ID,
		dto.LegacyLogin,
		dto.LegacyID,
		dto.LegacyNodeID,
		dto.LegacyURL,
		dto.LegacyHTMLURL,
		dto.NewEmail,
		dto.CreatedAt,
	)
	return
}

func (gumr GreatestUserMySQLRepository) SetEmpty() (err error) {
	query := "truncate table greatest_users.users"
	_, err = gumr.db.Exec(query)
	return
}
