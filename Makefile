
MOCKS := $(shell find . -type f -wholename '*mock.go')
FILES_TO_MOCK := $(shell find . -type f -wholename '*contract.go')

.PHONY: mocks
mocks: $(FILES_TO_MOCK)
	@rm -rf $(MOCKS)
	@for file in $^; do mockgen -source=$$file -destination=$${file//contract.go/mock.go}; done

run:
	go run content/cmd/main.go -c content/config.yml & \
	sleep 0.5 ; \
	go run session/cmd/main.go -c session/config.yml & \
	sleep 0.5 ; \
	go run user/cmd/main.go -c user/config.yml & \
	sleep 0.5 ; \
	go run favorites/cmd/main.go -c favorites/config.yml & \
	sleep 0.5 ; \
	go run cmd/main.go -c config.yml
