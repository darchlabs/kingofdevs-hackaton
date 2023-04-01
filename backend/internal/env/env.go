package env

type Env struct {
	DatabaseDSN  string `envconfig:"database_dsn" required:"true"`
	Port         string `envconfig:"port" required:"true"`
	ClientURL    string `envconfig:"client_url" required:"true"`
	MigrationDir string `envconfig:"migration_dir" required:"true"`
}
