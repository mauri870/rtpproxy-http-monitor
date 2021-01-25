.PHONY: build

NAME=rtpproxy-http-monitor

build:
	go build

install: build
	install -m755 $(NAME) /usr/local/bin/

init-d:
	install -m755 etc/init.d/$(NAME) /etc/init.d/

uninstall:
	-rm /usr/bin/$(NAME)

clean:
	-rm $(NAME)

