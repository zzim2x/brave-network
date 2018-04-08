NAME=brave

.PHONY: depend clean default

default: $(NAME)
$(NAME): $(shell find . -name \*.go)
	go build -ldflags "$(VFLAG)" -o bin/$(NAME) .

dep:
	glide install
