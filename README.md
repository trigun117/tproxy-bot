[![Build Status](https://travis-ci.org/trigun117/tproxy-bot.svg?branch=master)](https://travis-ci.org/trigun117/tproxy-bot) [![codecov](https://codecov.io/gh/trigun117/tproxy-bot/branch/master/graph/badge.svg)](https://codecov.io/gh/trigun117/tproxy-bot) [![Go Report Card](https://goreportcard.com/badge/github.com/trigun117/tproxy-bot)](https://goreportcard.com/report/github.com/trigun117/tproxy-bot)

# tproxy-bot

This is a telegram bot which fetch json with proxies and send to user

![example work of bot](https://github.com/trigun117/tproxy-bot/blob/master/image.JPG)

# Getting Started

For start using bot, build docker image from Dockerfile and run with this command
```
docker run \
-p 8443:8443 \
-e IP=server_ip_for_webhook \
-e TOKEN=set_your_bot_token \
--restart always \
-d your_image
```
Or with database
```
docker run \
-p 8443:8443 \
-e IP=server_ip_for_webhook \
-e TOKEN=set_your_bot_token \
-e HOST=set_your_database_host \
-e PORT=set_your_database_port \
-e USER=set_your_database_user \
-e PASSWORD=set_your_database_password \
-e DBNAME=set_your_database_name \
-e SSLMODE=set_your_database_sslmode(disable or enable, default disable) \
--restart always \
-d your_image
```

if you need to create table add
```
-e CREATE_TABLE=yes
```

# License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
