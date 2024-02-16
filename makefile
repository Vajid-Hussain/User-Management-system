 

GOCMD:=go
BUILD_DIR:=build
BINARY_DIR:=$(BUILD_DIR)/bin

all: test build

run:
	$(GOCMD) run ./cmd/api


wire:
	cd pkg/di && wire

mock:
	mockgen -source=pkg/repository/interface/user.go -destination=pkg/mock/mockrepo/user_mock.go -package=mockrepo