kitex_gen:
	kitex -module work idl/interact.thrift
	kitex -module work idl/relation.thrift
	kitex -module work idl/user.thrift
	kitex -module work idl/video.thrift
	kitex -module work idl/message.thrift

rpc_update: 
	cd ./rpc/interact/ && kitex -module work -service work.interact -use work/kitex_gen ../../idl/interact.thrift
	cd ./rpc/relation/ && kitex -module work -service work.relation -use work/kitex_gen ../../idl/relation.thrift
	cd ./rpc/user/ && kitex -module work -service work.user -use work/kitex_gen ../../idl/user.thrift
	cd ./rpc/video/ && kitex -module work -service work.video -use work/kitex_gen ../../idl/video.thrift
	cd ./rpc/message/ && kitex -module work -service work.message -use work/kitex_gen ../../idl/message.thrift

docker_start:
	docker compose --env-file=./docker/docker-compose-env.env --file=./docker/docker-compose.yaml --profile release up -d

docker_stop:
	docker compose down

facade:
	go run ./rpc/facade/main.go

user:
	go run ./rpc/user/*.go

video:
	go run ./rpc/video/*.go

interact:
	go run ./rpc/interact/*.go

message:
	go run ./rpc/message/*.go

relation:
	go run ./rpc/relation/*.go

jaeger:
	docker run -d --name jaeger -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 -p 5775:5775/udp -p 6831:6831/udp -p 6832:6832/udp -p 5778:5778 -p 16686:16686 -p 14268:14268 -p 14250:14250 -p 9411:9411 jaegertracing/all-in-one:latest