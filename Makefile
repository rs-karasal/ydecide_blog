ENV_FILE_LOCAL=.env.local
ENV_FILE_DEPLOY=.env.local

up:
	@echo "Starting services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_LOCAL) && docker-compose --env-file $(ENV_FILE_LOCAL) up --build -d

down:
	@echo "Stopping services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_LOCAL) && docker-compose --env-file $(ENV_FILE_LOCAL) down

down-v:
	@echo "Stopping services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_LOCAL) && docker-compose --env-file $(ENV_FILE_LOCAL) down -v



up-d:
	@echo "Starting services with $(ENV_FILE_DEPLOY)"
	@export ENV_FILE=$(ENV_FILE_DEPLOY) && docker-compose --env-file $(ENV_FILE_DEPLOY) up --build -d

down-d:
	@echo "Stopping services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_DEPLOY) && docker-compose --env-file $(ENV_FILE_DEPLOY) down

down-d-v:
	@echo "Stopping services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_DEPLOY) && docker-compose --env-file $(ENV_FILE_DEPLOY) down -v
