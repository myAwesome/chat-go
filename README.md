# chat-go
Simple chat implementation

Golang + Mysql  
gin, gorm

docker-compose up -d --build &&
docker-compose ps &&
pwd &&
echo "wait 15 sec for Database"
sleep 15 &&
docker exec -i app_mysql_1 mysql -uroot -proot < sql.sql &&
go get -d -v ./... && go install -v ./... && go build -o server .
echo "&& (./server )

