FROM dependencies_dev AS builder

WORKDIR /app

ENTRYPOINT CompileDaemon --build="go build -o main_service_dev cmd/main_service/start.go" --command=/app/main_service_dev
