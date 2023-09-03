#!/bin/bash
APP_NAME=gofinabot

mkdir -p $APP_NAME && cd $APP_NAME
wget https://raw.githubusercontent.com/Devil666face/gofinabot/main/Dockerfile
wget https://raw.githubusercontent.com/Devil666face/gofinabot/main/docker-compose.yaml
read -p "Create .env file? [y/n] " STATUS
if [[ "$STATUS" = "y" ]]; then
  read -p "TOKEN=" TOKEN
  echo "TOKEN=$TOKEN" >> .env
fi
docker-compose up -d