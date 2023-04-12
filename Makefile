
MOCKS := $(shell find . -type f -wholename '*mock.go')
FILES_TO_MOCK := $(shell find . -type f -wholename '*contract.go')

.PHONY: mocks
mocks: $(FILES_TO_MOCK)
	@rm -rf $(MOCKS)
	@for file in $^; do mockgen -source=$$file -destination=$${file//contract.go/mock.go}; done
