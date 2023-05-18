
MOCKS := $(shell find . -type f -wholename '*mock.go')
FILES_TO_MOCK := $(shell find . -type f -wholename '*contract.go')

.PHONY: mocks
mocks: $(FILES_TO_MOCK)
	@rm -rf $(MOCKS)
	@for file in $^; do mockgen -source=$$file -destination=$${file//contract.go/mock.go}; done


DATE := $(shell TZ=Europe/Moscow date +'%Y_%m_%d__%H_%M_%S')

HASH_COMMIT := $(shell git rev-parse --short=8 HEAD)
# новое название актуальной директории

NEW_BUILD_DIR := build_$(DATE)_$(HASH_COMMIT)
# действующая актуальная директория

# ACTUAL_BUILD_DIR = $(shell ls -la | grep build | sort -r | awk 'NR==1 {print $9}')
ACTUAL_BUILD_DIR = $(shell find . -maxdepth 1 -type d -name 'build_*_*_*_*_*_*' | sort -r | head -n 1 | xargs basename)

build:
	rm -rf ./build && \
	go build -o ./${NEW_BUILD_DIR}/content/out 		content/cmd/main.go 		&& \
	go build -o ./${NEW_BUILD_DIR}/session/out 		session/cmd/main.go 		&& \
	go build -o ./${NEW_BUILD_DIR}/user/out 		user/cmd/main.go 			&& \
	go build -o ./${NEW_BUILD_DIR}/user_action/out 	user_action/cmd/main.go 	&& \
	go build -o ./${NEW_BUILD_DIR}/payment/out 		payment/cmd/main.go 	&& \
	cp config.yml ./${NEW_BUILD_DIR}/content 		&& \
	cp config.yml ./${NEW_BUILD_DIR}/session 		&& \
	cp config.yml ./${NEW_BUILD_DIR}/user 			&& \
	cp config.yml ./${NEW_BUILD_DIR}/user_action	&& \
	cp config.yml ./${NEW_BUILD_DIR}/payment	&& \
	go build -o ./${NEW_BUILD_DIR}/api_gateway/out cmd/main.go 	&& \
	cp config.yml ./${NEW_BUILD_DIR}/api_gateway				|| \
	(rm -rf ./${NEW_BUILD_DIR} && exit 1)

	sed 's/build_dir/.\/${NEW_BUILD_DIR}/ig' docker-compose-template.yml > docker-compose.yml

rollback:
	rm -rf ${ACTUAL_BUILD_DIR}
	make update_docker_compose_volumes

PREV_BUILD_DIR = $(shell find . -maxdepth 1 -type d -name 'build_*_*_*_*_*_*' | sort -r | head -n 1 | xargs basename)
update_docker_compose_volumes:
	sed "s/build_dir/.\/${PREV_BUILD_DIR}/ig" docker-compose-template.yml > docker-compose.yml
	
# измененные бинарники попадают в директорию с актуальными версиями бинарников build_0_...
build_api: 
	rm -rf ${ACTUAL_BUILD_DIR}/api_gateway
	go build -o ${ACTUAL_BUILD_DIR}/api_gateway/out cmd/main.go
	cp config.yml ${ACTUAL_BUILD_DIR}/api_gateway

build_fav:
	rm -rf ${ACTUAL_BUILD_DIR}/user_action
	go build -o ${ACTUAL_BUILD_DIR}/user_action/out user_action/cmd/main.go
	cp config.yml ${ACTUAL_BUILD_DIR}/user_action

build_user:
	rm -rf ${ACTUAL_BUILD_DIR}/user
	go build -o ${ACTUAL_BUILD_DIR}/user/out user/cmd/main.go
	cp config.yml ${ACTUAL_BUILD_DIR}/user

build_content:
	rm -rf ${ACTUAL_BUILD_DIR}/content
	go build -o ${ACTUAL_BUILD_DIR}/content/out content/cmd/main.go
	cp config.yml ${ACTUAL_BUILD_DIR}/content

