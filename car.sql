DROP DATABASE IF EXISTS car;
CREATE DATABASE car;
USE car;

CREATE TABLE car(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(40),
    price int,
    PRIMARY KEY(id));

INSERT INTO car VALUES(1, 'Fararri', 5000);
INSERT INTO car VALUES(2, 'TATA', 2000);    
