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



up-deploy:
	@echo "Starting services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_LOCAL) && docker-compose --env-file $(ENV_FILE_LOCAL) up --build -d

down-deploy:
	@echo "Stopping services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_LOCAL) && docker-compose --env-file $(ENV_FILE_LOCAL) down

down-v-deploy:
	@echo "Stopping services with $(ENV_FILE_LOCAL)"
	@export ENV_FILE=$(ENV_FILE_LOCAL) && docker-compose --env-file $(ENV_FILE_LOCAL) down -v
