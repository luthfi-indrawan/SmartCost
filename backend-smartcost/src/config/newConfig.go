package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App      AppConfig
		Database DatabaseConfig
		Redis    RedisConfig
		Supabase SupabaseConfig
	}

	AppConfig struct {
		Env            string   `env:"APP_ENV" envDefault:"development"`
		Name           string   `env:"APP_NAME" envDefault:"my_app"`
		Port           string   `env:"APP_PORT" envDefault:"8080"`
		TimeoutSeconds int      `env:"APP_TIMEOUT_SECONDS" envDefault:"30"`
		AllowedOrigins []string `env:"APP_ALLOWED_ORIGINS" envSeparator:","` // otomatis di-parsing jadi slice/array
		SecretKey      string   `env:"APP_SECRET_KEY" required:"true"`       // Wajib diisi di prod, kalau kosong aplikasi akan error saat start
		LogLevel       string   `env:"APP_LOG_LEVEL" envDefault:"info"`
	}

	DatabaseConfig struct {
		Drive           string `env:"DB_DRIVE" envDefault:"postgres"`
		User            string `env:"DB_USER" envDefault:"postgres"`
		Host            string `env:"DB_HOST" envDefault:"localhost"`
		Port            string `env:"DB_PORT" envDefault:"5432"`
		Name            string `env:"DB_NAME" envDefault:"mydb"`
		Password        string `env:"DB_PASSWORD" envDefault:"secret"`
		SSLMode         string `env:"DB_SSL_MODE" envDefault:"disable"`
		MaxOpenConns    int    `env:"DB_MAX_OPEN_CONNS" envDefault:"25"`
		MaxIdleConns    int    `env:"DB_MAX_IDLE_CONNS" envDefault:"25"`
		ConnMaxLifetime int    `env:"DB_CONN_MAX_LIFETIME" envDefault:"300"`
		ConnMaxIdleTime int    `env:"DB_CONN_MAX_IDLE_TIME" envDefault:"180"`
	}

	RedisConfig struct {
		Addr         string `env:"REDIS_ADDR" envDefault:"localhost:6379"`
		Password     string `env:"REDIS_PASSWORD" envDefault:""`
		DB           int    `env:"REDIS_DB" envDefault:"0"`
		PoolSize     int    `env:"REDIS_POOL_SIZE" envDefault:"10"`
		MinIdleConns int    `env:"REDIS_MIN_IDLE_CONNS" envDefault:"3"`
		DialTimeout  int    `env:"REDIS_DIAL_TIMEOUT" envDefault:"5"`
		ReadTimeout  int    `env:"REDIS_READ_TIMEOUT" envDefault:"3"`
		WriteTimeout int    `env:"REDIS_WRITE_TIMEOUT" envDefault:"3"`
		PoolTimeout  int    `env:"REDIS_POOL_TIMEOUT" envDefault:"4"`
	}

	SupabaseConfig struct {
		URL             string `env:"SUPABASE_URL"`
		AnonKey         string `env:"SUPABASE_ANON_KEY"`
		AccessKeyID     string `env:"SUPABASE_STORAGE_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"SUPABASE_STORAGE_SECRET_ACCESS_KEY"`
		Region          string `env:"SUPABASE_STORAGE_REGION" envDefault:"ap-southeast-1"`
		BucketPublic    string `env:"SUPABASE_STORAGE_BUCKET_PUBLIC"`
		BucketPrivate   string `env:"SUPABASE_STORAGE_BUCKET_PRIVATE"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	opts := env.Options{RequiredIfNoDef: true}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		log.Printf("Gagal memparsing environment variables: %v", err)
		return nil, err
	}

	return cfg, nil
}