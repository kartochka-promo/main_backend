FROM dependencies_dev AS builder

WORKDIR /app

ENTRYPOINT CompileDaemon --build="go build -o survey_service_dev cmd/survey_service/start.go" --command=/app/survey_service_dev
