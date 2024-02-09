package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Drive string `mapstructure:"DB_DRIVE"`
	Host  string `mapstructure:"DB_HOST"`
	Port  int    `mapstructure:"DB_PORT"`
	User  string `mapstructure:"DB_USER"`
	Pass  string `mapstructure:"DB_PASS"`
	Name  string `mapstructure:"DB_NAME"`
}

type JwtConfig struct {
	Secret    string `mapstructure:"JWT_SECRET"`
	ExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	TokenAuth *jwtauth.JWTAuth
}

type Conf struct {
	ServerPort string         `mapstructure:"SERVER_PORT"`
	Database   DatabaseConfig `mapstructure:"DATABASE"`
	Jwt        JwtConfig      `mapstructure:"JWT"`
}

func Load(path string) (*Conf, error) {
	var env *Conf

	viper.SetConfigName("configs")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		panic(err)
	}

	setJwtConfig(env)
	setDatabaseConfig(env)

	env.Jwt.TokenAuth = jwtauth.New("HS256", []byte(env.Jwt.Secret), nil)

	return env, err
}

func setDatabaseConfig(env *Conf) {
	env.Database = DatabaseConfig{
		Drive: viper.GetString("DB_DRIVE"),
		Host:  viper.GetString("DB_HOST"),
		Port:  viper.GetInt("DB_PORT"),
		User:  viper.GetString("DB_USER"),
		Pass:  viper.GetString("DB_PASS"),
		Name:  viper.GetString("DB_NAME"),
	}
}

func setJwtConfig(env *Conf) {
	env.Jwt = JwtConfig{
		Secret:    viper.GetString("JWT_SECRET"),
		ExpiresIn: viper.GetInt("JWT_EXPIRES_IN"),
	}
}
