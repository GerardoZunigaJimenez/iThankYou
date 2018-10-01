docker build -t service/database .
docker run -d -p 3306:3306 --name idb service/database
docker exec -it idb bash
mysql -u service -p



describe users;

insert into users(last_name, first_name, hashpassword, email, createdAt, updatedAt) values ('macias','Jesus','z','jmacias@itexico.net',CURDATE(),CURDATE());


select * from users;