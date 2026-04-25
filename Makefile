

## Create a new microservice (usage: make create-service NAME=auth)
create-service:
	@if [ -z "$(NAME)" ]; then \
		echo "Usage: make create-service NAME=service-name"; \
		exit 1; \
	fi
	bash scripts/create_service.sh $(NAME)
