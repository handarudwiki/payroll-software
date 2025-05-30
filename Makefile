# Makefile untuk membaca file .env untuk URL database

# Memuat variabel lingkungan dari file .env
ENV_FILE ?= .env

ifneq ("$(wildcard $(ENV_FILE))","")
    include $(ENV_FILE)
    export
endif

MIGRATIONS_DIR = database/migrations

# Perintah untuk menjalankan migrasi
up:
	goose -dir $(MIGRATIONS_DIR) postgres $(DATABASE_URL) up

up_test : ENV_FILE = .env.test
up_test : 
	goose -dir ${MIGRATIONS_DIR} postgres ${DATABASE_URL} up

down:
	goose -dir $(MIGRATIONS_DIR) postgres $(DATABASE_URL) down



# Perintah untuk membuat migrasi baru
# Menggunakan variabel `name` yang dikirimkan melalui perintah `make`
new:
	goose -dir $(MIGRATIONS_DIR) create $(name) sql

test:
	@go test -v ./tests/e2e/main_test.go

run:
	@go run ./cmd/main.go

seed:
	@go run ./cmd/seed/seed.go

