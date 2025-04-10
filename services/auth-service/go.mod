module auth-service

go 1.24.1

replace github.com/MafiaLogiki/common/auth => ../../common/auth

replace github.com/MafiaLogiki/common/domain => ../../common/domain

replace github.com/MafiaLogiki/common/logger => ../../common/logger

replace github.com/MafiaLogiki/common/middleware => ../../common/middleware

require (
	github.com/MafiaLogiki/common/auth v0.0.0
	github.com/MafiaLogiki/common/domain v0.0.0
	github.com/MafiaLogiki/common/logger v0.0.0
	github.com/go-chi/chi/v5 v5.2.1
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/lib/pq v1.10.9
)

require (
	github.com/BurntSushi/toml v1.5.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/stretchr/testify v1.10.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)
