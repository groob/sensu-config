all: clean build

clean:
	rm -f ${PWD}/out/sensu-config
build:
	docker build -t sensu-config:build .	
	docker run -it --rm --name sensu-config-build -e CGO_ENABLED=0 -e GOOS=linux -v ${PWD}/out:/var/shared sensu-config:build go build -a -tags netgo -ldflags '-w' -o /var/shared/sensu-config
