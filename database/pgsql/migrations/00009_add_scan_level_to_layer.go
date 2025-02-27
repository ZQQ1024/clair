package migrations

import "github.com/remind101/migrate"

func init() {
	RegisterMigration(migrate.Migration{
		ID: 9,
		Up: migrate.Queries([]string{
			// Add new column scan_level with default value
			`ALTER TABLE public.layer ADD COLUMN scan_level INTEGER NOT NULL DEFAULT 1;`,
			// Add new column package_map
			`ALTER TABLE public.layer ADD COLUMN package_map TEXT;`,
			// Add comment on the layer table
			`COMMENT ON TABLE public.layer IS '不能覆盖同步';`,
		}),
		Down: migrate.Queries([]string{
			// Remove added columns
			`ALTER TABLE public.layer DROP COLUMN scan_level;`,
			`ALTER TABLE public.layer DROP COLUMN package_map;`,
			// Remove comment on the table
			`COMMENT ON TABLE public.layer IS NULL;`,
		}),
	})
}
