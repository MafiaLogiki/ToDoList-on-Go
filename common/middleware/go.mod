module middleware

go 1.22.2

replace github.com/MafiaLogiki/common/auth => ../auth/

replace github.com/MafiaLogiki/common/logger => ../logger/

require (
	github.com/MafiaLogiki/common/auth v0.0.0
	github.com/MafiaLogiki/common/logger v0.0.0
)

require (
	github.com/go-chi/chi/v5 v5.2.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.2 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
)
