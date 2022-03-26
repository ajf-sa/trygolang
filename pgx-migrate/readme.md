export POSTGRESQL_URL='postgres://admin:secret@localhost:5432/pgx-migrate-v1?sslmode=disable'

migrate create -ext sql -dir db/migrations -seq create_users_table

migrate -database ${POSTGRESQL_URL} -path db/migrations up

migrate -database ${POSTGRESQL_URL} -path db/migrations down

migrate create -ext sql -dir db/migrations -seq add_mood_to_users

migrate create -ext sql -dir db/migrations -seq add_name_to_users
