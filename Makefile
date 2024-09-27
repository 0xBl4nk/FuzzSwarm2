BINARY_NAME = FuzzSwarm
SRC_DIR = ./src
MAIN_FILE = main.go

all: build

build:
	go build -o ${BINARY_NAME} ${MAIN_FILE}
