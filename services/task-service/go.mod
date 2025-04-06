module task-service

go 1.22.2

replace github.com/MafiaLogiki/common/auth => ../../common/auth
replace github.com/MafiaLogiki/common/domain => ../../common/domain
replace github.com/MafiaLogiki/common/logger => ../../common/logger
replace github.com/MafiaLogiki/common/middleware => ../../common/middleware

require (
	github.com/go-chi/chi/v5 v5.2.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
    github.com/MafiaLogiki/common/auth v0.0.0
    github.com/MafiaLogiki/common/domain v0.0.0
)
