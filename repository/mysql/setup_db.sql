CREATE TABLE users
(
    id  int PRIMARY KEY AUTO_INCREMENT,
    name  varchar(255) not null ,
    phone_number varchar(255) not null UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
--  created_at datetime DEFAULT CURRENT_TIMESTAMP   ->   TIMESTAMP, datetime => []uint8
);