.PHONY: all clean
all: api swagger dao cleancache

api: worker-api master-api
worker: worker-api worker-swagger dao
master: master-api master-swagger dao cleancache
swagger: worker-swagger master-swagger
dao: mysql
cleancache:
	rm -rf tmp

worker-api:
	goctl api format --dir ./api
	goctl api go -api ./api/worker.api -dir ./worker -style goZero --home ./template

master-api:
	goctl api format --dir ./api
	goctl api go -api ./api/master.api -dir ./master -style goZero --home ./template


worker-api-port=8888
worker-swagger-port=8083

worker-swagger: worker-api
	goctl api plugin -plugin goctl-swagger='swagger -filename swag.json --host 127.0.0.1:$(worker-api-port)' -api ./api/worker.api -dir ./worker


master-api-port=8080
master-swagger-port=8084
master-swagger: master-api master.api
	goctl api plugin -plugin goctl-swagger='swagger -filename swag.json --host 127.0.0.1:$(master-api-port)' -api ./tmp/master.api -dir ./master


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

mysql:
	go run ./script/sql

clean:
	rm -rf tmp