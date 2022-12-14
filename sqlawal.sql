CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR(50) NOT NULL,
email VARCHAR(40) NOT NULL,
Phonenumber VARCHAR(15) NOT NULL,
Address text NOT NULL,
password varchar(50) NOT NULL
);

CREATE TABLE products (
id SERIAL PRIMARY KEY,
product_name VARCHAR(50) NOT NULL,
category VARCHAR(40) NOT NULL,
weight varchar(10) not null,
price int not null
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

SELECT * FROM products
SELECT * FROM users

INSERT INTO users (username,email,phonenumber,address,password) VALUES ('hamad','hamad@mail','082692762','jakarta','123') RETURNING id

DELETE FROM products;

ALTER TABLE products
DROP COLUMN category;

CREATE TABLE detail_user (
id SERIAL PRIMARY KEY,
user_id int NOT NULL,
email VARCHAR(40) NOT NULL,
Phonenumber VARCHAR(15) NOT NULL,
Address text NOT NULL,
	
CONSTRAINT fk_user 
	FOREIGN KEY (user_id)
	REFERENCES users(id)
);

CREATE TABLE categories (
id SERIAL PRIMARY KEY,
category_name Varchar(40) NOT NULL
);

CREATE TABLE products (
id SERIAL PRIMARY KEY,
category_id int NOT NULL,
product_name VARCHAR(50) NOT NULL,
weight varchar(10) not null,
price int not null,
	
CONSTRAINT fk_product_category 
	FOREIGN KEY (category_id)
	REFERENCES categories(id)
);

CREATE TABLE cart (
id SERIAL PRIMARY KEY,
user_id int,
product_id int,
qty int NOT NULL,
Total int not null,
created_at TIMESTAMP,
	
CONSTRAINT fk_user_cart
	FOREIGN KEY (user_id)
	REFERENCES users(id),
	
CONSTRAINT fk_product_cart 
	FOREIGN KEY (product_id)
	REFERENCES products(id)
);

drop table transactions;

CREATE TABLE transactions (
id SERIAL PRIMARY KEY,
cart_id int,
payment_id int,
created_at TIMESTAMP,
	
CONSTRAINT fk_cart_transaction 
	FOREIGN KEY (cart_id)
	REFERENCES cart(id),

CONSTRAINT fk_payment_transaction 
	FOREIGN KEY (payment_id)
	REFERENCES payments(id)
);

CREATE TABLE payments (
id SERIAL PRIMARY KEY,
payment_method VARCHAR(50) NOT NULL
);

select * from products

alter table users 
add column email VARCHAR(30) NOT NULL,
ADD COLUMN phone_number VARCHAR(15) not null,
add column address text;

alter table products
alter column category_id type varchar(40)

select * from users

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