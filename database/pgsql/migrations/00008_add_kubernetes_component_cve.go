package migrations

import "github.com/remind101/migrate"

func init() {
	RegisterMigration(migrate.Migration{
		ID: 8,
		Up: migrate.Queries([]string{
			`CREATE TABLE IF NOT EXISTS kubernetes_component_cve (
        cve_id TEXT NOT NULL,
        component_name TEXT NOT NULL,
        description TEXT NOT NULL,
        vector_string TEXT NOT NULL,
        severity TEXT NOT NULL,
        score DOUBLE PRECISION NOT NULL,
        hyperlink JSON NOT NULL,
        affected_in JSON NOT NULL
      );`,
		}),
		Down: migrate.Queries([]string{
			`DROP TABLE IF EXISTS kubernetes_component_cve;`,
		}),
	})
}
