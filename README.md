[![Build Status](https://travis-ci.org/trigun117/tproxy-bot.svg?branch=master)](https://travis-ci.org/trigun117/tproxy-bot) [![codecov](https://codecov.io/gh/trigun117/tproxy-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/trigun117/tproxy-bot) [![Go Report Card](https://goreportcard.com/badge/github.com/trigun117/tproxy-bot)](https://goreportcard.com/report/github.com/trigun117/tproxy-bot)

# tproxy-bot

This is a telegram bot which fetch json with proxies and send to user

![example work of bot](https://github.com/trigun117/tproxy-bot/blob/master/image.JPG)

# Getting Started

For start using bot, build docker image from Dockerfile and run with this command
```
docker run \
-e TOKEN=set_your_bot_token \
-e LINK=site_link \
-e MTL=mtproto_link \
-e URL=site_json_url \
-e PASS=site_json_password \
-e FI=site_json_field \
--restart always \
-d your_image
```
Or with database
```
docker run \
-e TOKEN=set_your_bot_token \
-e HOST=set_your_database_host \
-e PORT=set_your_database_port \
-e USER=set_your_database_user \
-e PASSWORD=set_your_database_password \
-e DBNAME=set_your_database_name \
-e SSLMODE=set_your_database_sslmode(disable or enable, default disable) \
-e LINK=site_link \
-e MTL=mtproto_link \
-e URL=site_json_url \
-e PASS=site_json_password \
-e FI=site_json_field \
--restart always \
-d your_image
```

if you need to create table add
```
-e TABLE=yes
```
# Docker-Compose

Set environment variables in docker-compose.yml

```
services:

  db:
    image: postgres
    environment:
      POSTGRES_PASSWORD: set_your_database_password
  
  bot:
    image: trigun117/tproxy-bot
    environment:
      TOKEN: set_your_bot_token
      TABLE: "yes"
      HOST: db
      PORT: 5432
      USER: postgres
      PASSWORD: set_your_database_password
      DBNAME: postgres
      SSLMODE: disable
      LINK: site_link
      MTL: mtproto_link
      URL: site_json_url
      PASS: site_json_password
      FI: site_json_field
```
And launch with this command

```
docker-compose up --build
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
