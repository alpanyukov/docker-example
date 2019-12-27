# Практикум по Docker Client

## Скрипты

Собрать образ с приложением

```
docker build -t learn_docker .
```

Стянуть образ redis

```
docker pull redis:alpine
```

Создать сеть типа bridge

```
docker network create backend
```

Создать том для внешнего хранения

```
docker volume create --driver local --opt type=none --opt device=d/shared --opt o=bind file-data
```

Создать контейнер на основе образа redis

```
docker create --name app_redis --network backend redis:alpine
```

Создать контейнер на основе образа приложения

```
docker create --name app -p 3000:3000 --network backend -e "REDIS_URL=app_redis:6379" --mount source=file-data,target=/data learn_docker
```

Запустить контейнер с redis

```
docker start app_redis
```

Запустить контейнер с приложением

```
docker start app
```

## Docker Compose

Собрать все

```
docker-compose build
```

Запустить все

```
docker compose up -d
```
