FROM dependencies_dev AS builder

WORKDIR /app

ENTRYPOINT CompileDaemon --build="go build -o mail_service_dev cmd/mail_service/start.go" --command=/app/mail_service_dev
