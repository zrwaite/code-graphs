source .env
docker run --name redis -d -p 6379:6379 redis redis-server --requirepass "${REDIS_PASSWORD}"
