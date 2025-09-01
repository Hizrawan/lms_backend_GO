package config

type MigrationConfig struct {
	Migrate         bool `mapstructure:"migrate"`
	Version         int  `mapstructure:"version"`
	RollbackOnError bool `mapstructure:"rollback_on_error"`
	AllowDrop       bool `mapstructure:"allow_drop"`
}
