MAIN_SERVICE_BINARY=main_service
STAFF_BINARY=staff_service
SURVEY_BINARY=survey_service
MAIL_BINARY=mail_service

PROJECT_DIR := ${CURDIR}

DOCKER_PROD_DIR := ${CURDIR}/docker/production
DOCKER_DEV_DIR := ${CURDIR}/docker/development

## build: Build compiles project
build:
	go build -o ${MAIN_SERVICE_BINARY} cmd/main_service/start.go
	go build -o ${STAFF_BINARY} cmd/staff_service/start.go
	go build -o ${SURVEY_BINARY} cmd/survey_service/start.go
	go build -o ${MAIL_BINARY} cmd/mail_service/start.go

## build-docker: Builds all docker containers for production
build-docker:
	docker build -t dependencies -f ${DOCKER_PROD_DIR}/builder.Dockerfile .
	docker build -t main_service -f ${DOCKER_PROD_DIR}/main_service.Dockerfile .
	docker build -t staff_service -f ${DOCKER_PROD_DIR}/staff.Dockerfile .
	docker build -t survey_service -f ${DOCKER_PROD_DIR}/survey.Dockerfile .
	docker build -t mail_service -f ${DOCKER_PROD_DIR}/mail.Dockerfile .

## build-docker-dev: Builds all docker containers for development
build-docker-dev:
	docker build -t dependencies_dev -f ${DOCKER_DEV_DIR}/init.Dockerfile .
	docker build -t main_service_dev -f ${DOCKER_DEV_DIR}/main_service.Dockerfile .
	docker build -t staff_service_dev -f ${DOCKER_DEV_DIR}/staff.Dockerfile .
	docker build -t survey_service_dev -f ${DOCKER_DEV_DIR}/survey.Dockerfile .
	docker build -t mail_service_dev -f ${DOCKER_DEV_DIR}/mail.Dockerfile .

## run-and-build: Build and run docker
build-and-run: build-docker
	docker-compose up

## run: Build and run docker with new changes
run:
	docker rm -vf $$(docker ps -a -q) || true
	make build-docker
	docker-compose up --build --no-deps

## run-dev: Build and run docker with new changes in develop mode
run-dev:
	docker rm -vf $$(docker ps -a -q) || true
	make build-docker-dev
	docker-compose -f docker-compose-dev.yml up --build --no-deps

## test-coverage: get final code coverage
coverage:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	cat cover | fgrep -v "mock" | fgrep -v "pb.go" | fgrep -v "easyjson" | fgrep -v "start.go" > cover2
	go tool cover -func=cover2
	rm -rf cover
	rm -rf cover2

## coverage-html: generates HTML file with test coverage
test-html:
	go test -covermode=atomic -coverpkg=./... -coverprofile=cover ./...
	go tool cover -html=cover
	rm -rf cover

## run-background: run process in background(available after build)
run-background:
	docker rm -vf $$(docker ps -a -q) || true
	make build-docker
	docker-compose up -d --build --no-deps

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command to run:"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
