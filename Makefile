.PHONY: all clean
all: api swagger dao cleancache

api: worker-api master-api
worker-update: worker-api worker-swagger-update dao
master-update: master-api master-swagger-update dao cleancache
swagger-update: worker-swagger-update master-swagger-update
dao: mysql

worker-api-port=8888
worker-swagger-port=8083

master-api-port=8080
master-swagger-port=8084

mkfile_path := $(shell pwd)

cleancache:
	rm -rf tmp

#>>>>>>>>>>>>>>>>>>>>>>>> Worker Command <<<<<<<<<<<<<<<<<<<<<< 
worker-api:
	goctl api format --dir ./api
	goctl api go -api ./api/worker.api -dir ./worker -style goZero --home ./template

worker-swagger-update: worker-api
	goctl api plugin -plugin goctl-swagger='swagger -filename swag.json --host 127.0.0.1:$(worker-api-port)' -api ./api/worker.api -dir ./worker

worker-swagger-run: worker-swagger-update
	docker run --rm --privileged -d -p $(worker-swagger-port):8080 -e SWAGGER_JSON=/app/worker.json -v $(mkfile_path)/worker/swagger:/app swaggerapi/swagger-ui

worker-run: 
	go run worker/worker.go -f worker/etc/worker-api.yaml

#>>>>>>>>>>>>>>>>>>>>>>>> Master Command <<<<<<<<<<<<<<<<<<<<<< 
master-api:
	goctl api format --dir ./api
	goctl api go -api ./api/master.api -dir ./master -style goZero --home ./template

master-swagger-update: master-api master.api
	goctl api plugin -plugin goctl-swagger='swagger -filename swag.json --host 127.0.0.1:$(master-api-port)' -api ./tmp/master.api -dir ./master
	rm -rf tmp

master-swagger-run: master-swagger-update
	docker run --rm --privileged -d -p $(master-swagger-port):8080 -e SWAGGER_JSON=/app/master.json -v $(mkfile_path)/master/swagger:/app swaggerapi/swagger-ui

master-run: 
	go run master/master.go -f master/etc/master-api.yaml

master.api: master-api
	mkdir tmp
	cat api/global.api >> tmp/master.api	
	echo "" >> tmp/master.api
	cat api/cluster.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/container.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/cronjob.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/deployment.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/namespace.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/job.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/node.api >> tmp/master.api
	echo "" >> tmp/master.api
	cat api/token.api >> tmp/master.api
	sed s/"global.api"//g tmp/master.api > tmp/master.api

#>>>>>>>>>>>>>>>>>>>>>>>> Other Command <<<<<<<<<<<<<<<<<<<<<< 
mysql:
	go run ./script/sql

clean:
	rm -rf tmp