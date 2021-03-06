package config

// AdminUser represents the default admin user.
type AdminUser struct {
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}
