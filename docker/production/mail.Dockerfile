FROM dependencies AS builder

WORKDIR /app

CMD /app/mail_service
