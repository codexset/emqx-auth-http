FROM alpine:edge

COPY dist /app
WORKDIR /app

EXPOSE 8080

CMD [ "./emqx-auth-http" ]