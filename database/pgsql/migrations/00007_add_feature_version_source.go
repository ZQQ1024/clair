package migrations

import "github.com/remind101/migrate"

func init() {
	RegisterMigration(migrate.Migration{
		ID: 7,
		Up: migrate.Queries([]string{
			`CREATE TABLE IF NOT EXISTS featureversion_source (
        id SERIAL PRIMARY KEY,
        type TEXT,
        name TEXT,
        content TEXT
      );`,
			`ALTER TABLE featureversion ADD COLUMN source_id INT NOT NULL REFERENCES featureversion_source ON DELETE CASCADE;`,
		}),
		Down: migrate.Queries([]string{
			`DROP TABLE IF EXISTS featureversion_source;`,
			`ALTER TABLE featureversion DROP COLUMN source_id;`,
		}),
	})
}
