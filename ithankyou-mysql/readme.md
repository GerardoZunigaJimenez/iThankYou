docker build -t ithankyou/database .
docker run -d -p 3306:3306 --name idb ithankyou/database
docker exec -it idb bash
mysql -u service -p
