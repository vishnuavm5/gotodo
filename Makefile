include .env

create_migrations:
	@echo "starting migrations"
		cd sql/schema && goose postgres postgres://user:password@localhost:5040/test up

run_db_container:
	@echo "starting db container at 5040"
		docker run --name DB_DOCKER_CONTAINER -p 5040:5432 --network=todonet  -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -d postgres:12-alpine
run_api_container:
	@echo "starting api container"
		docker run --name mytodo -p 8000:8000 --network=todonet  --env-file .env gotodoapi
stop_all_containers:
	@echo "stopping all..."
		docker stop $(docker ps -a -q)
create_db:
	@echo "create db"
		docker exec -it DB_DOCKER_CONTAINER createdb --username=user --owner=user test