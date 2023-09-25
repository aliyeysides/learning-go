
status:
	goose -dir ./migrations postgres "$(shell cat .env.db)" status

create:
	goose -dir ./migrations postgres "$(shell cat .env.db)" create $(name) sql

migrate:
	goose -dir ./migrations postgres "$(shell cat .env.db)" up
