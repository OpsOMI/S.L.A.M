package models

type App struct {
	Mode          string `yaml:"mode"`
	MigrationPath string `yaml:"migration_path"`
}
