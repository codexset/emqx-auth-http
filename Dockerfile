FROM alpine:edge

COPY dist /app
WORKDIR /app

CMD [ "./emqx-auth-http" ]