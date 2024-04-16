package migrations

import (
	"database/sql"
	"fmt"
)

const (
	ordersTableCreate = "create table if not exists orders (id varchar(255), price decimal(10,2), tax decimal(10,2), final_price decimal(10,2), primary key (id));"
)

var migrations = map[string]string{
	"orders_table_create": ordersTableCreate,
}

func RunMigrations(db *sql.DB) error {
	for name, migration := range migrations {
		_, err := db.Exec(migration)
		if err != nil {
			return err
		}

		fmt.Printf("Migration %s ran successfully\n", name)
	}

	return nil
}
