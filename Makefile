
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

# make -B build
.PHONY:
build:
	rm -rf ./build && \
	go build -o ./build/content/out content/cmd/main.go 		&& \
	go build -o ./build/session/out session/cmd/main.go 		&& \
	go build -o ./build/user/out user/cmd/main.go 				&& \
	go build -o ./build/user_action/out user_action/cmd/main.go 	&& \
	go build -o ./build/payment/out payment/cmd/main.go && \
	cp config.yml ./build/content
	cp config.yml ./build/session
	cp config.yml ./build/user
	cp config.yml ./build/user_action
	cp config.yml ./build/payment
	
	go build -o ./build/api_gateway/out cmd/main.go
	cp config.yml ./build/api_gateway	

build_api:
	rm -rf ./build/api_gateway && \
	go build -o ./build/api_gateway/out cmd/main.go
	cp config.yml ./build/api_gateway

build_fav:
	rm -rf ./build/user_action && \
	go build -o ./build/user_action/out user_action/cmd/main.go
	cp config.yml ./build/user_action

build_user:
	rm -rf ./build/user && \
	go build -o ./build/user/out user/cmd/main.go
	cp config.yml ./build/user

build_content:
	rm -rf ./build/content && \
	go build -o ./build/content/out content/cmd/main.go
	cp config.yml ./build/content

move_config:
	cp config.yml ./build/content
	cp config.yml ./build/session
	cp config.yml ./build/user
	cp config.yml ./build/favorites
	cp config.yml ./build/api_gateway	
