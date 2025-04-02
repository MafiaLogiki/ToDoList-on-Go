module auth-service

go 1.22.2

replace github.com/MafiaLogiki/common/auth => ../../common/auth

replace github.com/MafiaLogiki/common/domain => ../../common/domain

replace github.com/MafiaLogiki/common/logger => ../../common/logger

replace github.com/MafiaLogiki/common/middleware => ../../common/middleware

require (
	github.com/MafiaLogiki/common/auth v0.0.0
	github.com/MafiaLogiki/common/domain v0.0.0-00010101000000-000000000000
	github.com/MafiaLogiki/common/logger v0.0.0
	github.com/MafiaLogiki/common/middleware v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.2.1
	github.com/lib/pq v1.10.9
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
