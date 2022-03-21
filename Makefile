# This included makefile should define the 'custom' target rule which is called here.
# The (-) sign before `include` will do an implicit check if the file exists.
-include $(INCLUDE_MAKEFILE)

.PHONY: release
release: custom

.PHONY: release-local
release-local: custom-local

.PHONY: test
test: custom-test

# build images from all services
.PHONY: build-docker-local
build-docker-local:
	./tools/build-docker-local.sh

build-image-local:
	./tools/build-image-local.sh ${service} ${tag}

.PHONY: up-local
up-local:
	docker-compose -f docker-compose.yml -f docker-compose.db.yml --env-file ./.env.local up --build -d

.PHONY: down-local
down-local:
	docker-compose -f docker-compose.yml -f docker-compose.db.yml --env-file ./.env.local down