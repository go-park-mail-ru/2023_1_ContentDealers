
MOCKS := $(shell find . -type f -wholename '*mock.go')
FILES_TO_MOCK := $(shell find . -type f -wholename '*contract.go')

.PHONY: mocks
mocks: $(FILES_TO_MOCK)
	@rm -rf $(MOCKS)
	@for file in $^; do mockgen -source=$$file -destination=$${file//contract.go/mock.go}; done


DATE := $(shell date +'%Y%m%d_%H%M%S')

HASH_COMMIT := $(shell git rev-parse --short=8 HEAD)
# новое название актуальной директории

BUILD_DIR := build_0_$(DATE)_$(HASH_COMMIT)
# действующая актуальная директория

BUILD_DIR_0 := $(shell find . -maxdepth 1 -type d -name 'build_0_*' -print -quit | xargs basename)

# make -B build

# ======[ СБОРКА ]======
# 1. икремент версии директорий build_(version)_(date)_(hashcommit)
# 2. сборка бинарников и сохранение в build_0_...
# 3. если сборка завершилась в ошибкой
# 	- версия в директориях декрементируется
# 	- в docker-compose.yml ничего не изменяется
# 4. если сборка завершилась успешно
# 	- происходит изменение названия директории build в docker-compose.yml в volume
build:
	make increment_versions
	rm -rf ./build && \
	go build -o ./${BUILD_DIR}/content/out 		content/cmd/main.go 		&& \
	go build -o ./${BUILD_DIR}/session/out 		session/cmd/main.go 		&& \
	go build -o ./${BUILD_DIR}/user/out 		user/cmd/main.go 			&& \
	go build -o ./${BUILD_DIR}/user_action/out 	user_action/cmd/main.go 	&& \
	cp config.yml ./${BUILD_DIR}/content 		&& \
	cp config.yml ./${BUILD_DIR}/session 		&& \
	cp config.yml ./${BUILD_DIR}/user 			&& \
	cp config.yml ./${BUILD_DIR}/user_action	&& \
	go build -o ./${BUILD_DIR}/api_gateway/out cmd/main.go 	&& \
	cp config.yml ./${BUILD_DIR}/api_gateway				|| \
	(rm -rf ./${BUILD_DIR} && make decrement_versions && exit 1)

	sed 's/build_dir/.\/${BUILD_DIR}/ig' docker-compose-template.yml > docker-compose.yml

# ======[ ОТКАТ ]======
# 1. удаляется старая актуальная директория
# 2. декремент версий в названиях директорий (build_1 -> build_0)
# 3. обновляется директория в docker-compose.yml в volume с учетом отката
rollback:
	rm -rf ${BUILD_DIR_0}
	make decrement_versions
	make update_docker_compose_volumes

BUILD_DIR_0_NEW := $(shell find . -maxdepth 1 -type d -name 'build_0_*' -print -quit | xargs basename)
update_docker_compose_volumes:
	sed "s/build_dir/.\/${BUILD_DIR_0_NEW}/ig" docker-compose-template.yml > docker-compose.yml

increment_versions:
	for dir in ./build_*_*; do \
        if echo "$$dir" | grep -qE '^\.\/build_[0-9]+_.*$$'; then \
            num=$$(echo "$$dir" | sed -E 's/^\.\/build_([0-9]+)_.*$$/\1/'); \
            suffix=$$(echo "$$dir" | sed -E 's/^\.\/build_[0-9]+_(.*)$$/\1/'); \
            new_num=$$(expr $$num + 1); \
            new_dir="build_$$new_num"_"$$suffix"; \
            mv "$$dir" "$$new_dir"; \
        fi \
    done

decrement_versions:
	for dir in ./build_*_*; do \
        if echo "$$dir" | grep -qE '^\.\/build_[0-9]+_.*$$'; then \
            num=$$(echo "$$dir" | sed -E 's/^\.\/build_([0-9]+)_.*$$/\1/'); \
            suffix=$$(echo "$$dir" | sed -E 's/^\.\/build_[0-9]+_(.*)$$/\1/'); \
            new_num=$$(expr $$num - 1); \
            new_dir="build_$$new_num"_"$$suffix"; \
            mv "$$dir" "$$new_dir"; \
        fi \
    done

# ======[ СБОРКА ОТДЕЛЬНЫХ СЕРВИСОВ ]======
# измененные бинарники попадают в директорию с актуальными версиями бинарников build_0_...
build_api: 
	rm -rf ${BUILD_DIR_0}/api_gateway
	go build -o ${BUILD_DIR_0}/api_gateway/out cmd/main.go
	cp config.yml ${BUILD_DIR_0}/api_gateway

build_fav:
	rm -rf ${BUILD_DIR_0}/user_action
	go build -o ${BUILD_DIR_0}/user_action/out user_action/cmd/main.go
	cp config.yml ${BUILD_DIR_0}/user_action

build_user:
	rm -rf ${BUILD_DIR_0}/user
	go build -o ${BUILD_DIR_0}/user/out user/cmd/main.go
	cp config.yml ${BUILD_DIR_0}/user

build_content:
	rm -rf ${BUILD_DIR_0}/content
	go build -o ${BUILD_DIR_0}/content/out content/cmd/main.go
	cp config.yml ${BUILD_DIR_0}/content

