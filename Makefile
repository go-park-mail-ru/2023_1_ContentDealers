
MOCKS_DESTINATION := internal/mocks
FILES_IN_REPO := $(wildcard, internal/repository/*)
TEST_FILES_IN_REPO := $(wildcard, internal/repository/*_test.go)
FILES_TO_MOCK := $(filter-out $FILES_IN_REPO, $(TEST_FILES_IN_REPO))
.PHONY: mocks


mocks: $(FILES_TO_MOCK)
	echo "Generating mocks..."
	rm -rf $(MOCKS_DESTINATION)
	for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done