
IMPORT_PATH := ./
CPP_DST_DIR := ./cpp
GO_DST_DIR := ./go
JAVA_DST_DIR := ./java
PYTHON_DST_DIR := ./python
RUBY_DST_DIR := ./ruby
OBJC_DST_DIR := ./objc

PROTOC3 := $(shell which protoc)

default: all

.DEFAULT:
	for dir in cpp go java python ruby objc;do \
	    mkdir -p $$dir; \
	done
	$(PROTOC3) --proto_path=$(IMPORT_PATH) --cpp_out=$(CPP_DST_DIR) --java_out=$(JAVA_DST_DIR) --python_out=$(PYTHON_DST_DIR) --go_out=$(GO_DST_DIR) --ruby_out=$(RUBY_DST_DIR) --objc_out=$(OBJC_DST_DIR)  example.proto

.PHONY: clean
clean:
	rm -rf *.h *cc *.go *.java *.rb *.py *.m
	rm -rf cpp go java python ruby objc
