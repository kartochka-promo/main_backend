FROM dependencies_dev AS builder

WORKDIR /app

ENTRYPOINT CompileDaemon --build="go build -o staff_service_dev cmd/staff_service/start.go" --command=/app/staff_service_dev
