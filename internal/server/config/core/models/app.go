package models

type App struct {
	Mode          string
	MigrationPath string `yaml:"migration_path"`
}
