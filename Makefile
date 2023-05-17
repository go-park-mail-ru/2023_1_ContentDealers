
MOCKS := $(shell find . -type f -wholename '*mock.go')
FILES_TO_MOCK := $(shell find . -type f -wholename '*contract.go')

.PHONY: mocks
mocks: $(FILES_TO_MOCK)
	@rm -rf $(MOCKS)
	@for file in $^; do mockgen -source=$$file -destination=$${file//contract.go/mock.go}; done

# netstat -tulpn - проверка  
# ctrl C останавливает все
run:
	go run content/cmd/main.go -c config.yml & \
	sleep 0.5 ; \
	go run session/cmd/main.go -c config.yml & \
	sleep 0.5 ; \
	go run user/cmd/main.go -c config.yml & \
	sleep 0.5 ; \
	go run favorites/cmd/main.go -c config.yml & \
	sleep 0.5 ; \
	go run cmd/main.go -c config.yml

# ctrl C ничего не останавливает, далее kill -9
run_bin:
	./build/content/out -c 		./build/content/config.yml & \
	sleep 0.5 ; \
	./build/session/out -c 		./build/session/config.yml & \
	sleep 0.5 ; \
	./build/user/out -c 		./build/user/config.yml & \
	sleep 0.5 ; \
	./build/favorites/out -c 	./build/favorites/config.yml & \
	sleep 0.5 ; \
	./build/api_gateway/out -c 	./build/api_gateway/config.yml & \

DATE := $(shell date +'%Y%m%d_%H%M%S')
HASH_COMMIT := 
BUILD_DIR := build_0_$(DATE)_$(HASH_COMMIT)
BUILD_DIR_0 := $(shell find . -maxdepth 1 -type d -name 'build_0_*' -print -quit)

# make -B build
# 1. икремент версии директорий build_(version)_(date)_(hashcommit)
# 2. сборка бинарников и сохранение в build_0_...
# 3. изменение названия директории build в docker-compose.yml в volume
# 4. если сборка завершилась в ошибкой, то версия декрементируется, в docker-compose.yml ничего не изменяется
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

