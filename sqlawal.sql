CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR(50) NOT NULL,
password varchar(255) NOT NULL
email VARCHAR(30) NOT NULL,
Phonenumber VARCHAR(15) NOT NULL,
Address text NOT NULL,

);

CREATE TABLE products (
id SERIAL PRIMARY KEY,
category VARCHAR(40) NOT NULL,
product_name VARCHAR(50) NOT NULL,
weight varchar(10) not null,
price int not null
);


CREATE TABLE transactions (
id SERIAL PRIMARY KEY,
user_id int,
payment_id int,
created_at TIMESTAMP,
total INT NOT NULL,
	
CONSTRAINT fk_user_transaction 
	FOREIGN KEY (user_id)
	REFERENCES cart(id),

CONSTRAINT fk_payment_transaction 
	FOREIGN KEY (payment_id)
	REFERENCES payments(id)
);

CREATE TABLE payments (
id SERIAL PRIMARY KEY,
payment_method VARCHAR(50) NOT NULL
);

CREATE TABLE cart (
id SERIAL PRIMARY KEY,
user_id int,
product_id int,
qty int NOT NULL,
created_at TIMESTAMP,
total int not null,

	
CONSTRAINT fk_user_cart
	FOREIGN KEY (user_id)
	REFERENCES users(id),
	
CONSTRAINT fk_product_cart 
	FOREIGN KEY (product_id)
	REFERENCES products(id)
);


INSERT INTO products (product_name,category,weight,price)
VALUES 
('BOLT','DRY FOOD','1KG',23000),
('ORICAT','DRY FOOD','1KG',23000),
('PROPLAN','DRY FOOD','500GR',40000),
('ROYAL CANIN','DRY FOOD','1KG',100000),
('WHISKAS ADULT TUNA','WET FOOD','400GR',7000),
('WHISKAS JUNIOR TUNA','WET FOOD','400GR',7000),
('LIFECAT ADULT TUNA','WET FOOD','400GR',6000),
('KALUNG','ACCESORIES','-',20000),
('TONGKAT BULU','ACCESORIES','-',20000),
('BAJU','ACCESORIES','-',20000);

INSERT INTO payments (payment_method) VALUES
('CASH'),
('BCA'),
('BNI'),
('GOPAY');