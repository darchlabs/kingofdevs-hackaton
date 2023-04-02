package env

type Env struct {
	ClientURL           string `envconfig:"client_url" required:"true"`
	DatabaseDSN         string `envconfig:"database_dsn" required:"true"`
	Port                string `envconfig:"port" required:"true"`
	MigrationDir        string `envconfig:"migration_dir" required:"true"`
	SynchronizersApiURL string `envconfig:"synchronizers_api_url" required:"true"`
}
