
MOCKS := $(shell find . -type f -wholename '*mock.go')
FILES_TO_MOCK := $(shell find . -type f -wholename '*contract.go')

.PHONY: mocks
mocks: $(FILES_TO_MOCK)
	@rm -rf $(MOCKS)
	@for file in $^; do mockgen -source=$$file -destination=$${file//contract.go/mock.go}; done


DATE := $(shell TZ=Europe/Moscow date +'%Y_%m_%d__%H_%M_%S')

HASH_COMMIT := $(shell git rev-parse --short=8 HEAD)
# новое название актуальной директории

NEW_BUILD_DIR := build_versions/build_$(DATE)_$(HASH_COMMIT)
# действующая актуальная директория

# ACTUAL_BUILD_DIR = $(shell ls -la | grep build | sort -r | awk 'NR==1 {print $9}')
ACTUAL_BUILD_DIR = build_versions/$(shell find ./build_versions -maxdepth 1 -type d -name 'build_*_*_*_*_*_*' | sort -r | head -n 1 | xargs basename)

build:
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

	rm -rf build
	mkdir -p build/
	cp -r ${NEW_BUILD_DIR}/* build


rollback:
	rm -rf ${ACTUAL_BUILD_DIR}
	make update_build_dir

PREV_BUILD_DIR = build_versions/$(shell find ./build_versions -maxdepth 1 -type d -name 'build_*_*_*_*_*_*' | sort -r | head -n 1 | xargs basename)
update_build_dir:
	rm -rf build
	mkdir -p build
	cp -r ${PREV_BUILD_DIR}/* build/
	
build_api: 
	rm -rf build/api_gateway
	go build -o build/api_gateway/out cmd/main.go
	cp config.yml build/api_gateway

build_fav:
	rm -rf build/user_action
	go build -o build/user_action/out user_action/cmd/main.go
	cp config.yml build/user_action

build_user:
	rm -rf build/user
	go build -o build/user/out user/cmd/main.go
	cp config.yml build/user

build_content:
	rm -rf build/content
	go build -o build/content/out content/cmd/main.go
	cp config.yml build/content

