CREATE TABLE users (
id SERIAL PRIMARY KEY,
username VARCHAR(50) NOT NULL,
password varchar(255) NOT NULL
email VARCHAR(30) NOT NULL,
phone_number VARCHAR(15) NOT NULL,
address text NOT NULL,

);

CREATE TABLE products (
id SERIAL PRIMARY KEY,
category_id INT NOT NULL,
product_name VARCHAR(50) NOT NULL,
weight varchar(10) not null,
price INT NOT NULL,
CONSTRAINT fk_category_id_products 
	FOREIGN KEY (category_id)
	REFERENCES categories(id)
);

CREATE TABLE categories (
id SERIAL PRIMARY KEY,
category_name VARCHAR(40) NOT NULL
)


CREATE TABLE transactions (
id SERIAL PRIMARY KEY,
user_id int,
payment_id int,
created_at TIMESTAMP,
total INT NOT NULL,
	
CONSTRAINT fk_user_transaction 
	FOREIGN KEY (user_id)
	REFERENCES users(id),

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
('BOLT',1,'1KG',23000),
('ORICAT',1,'1KG',23000),
('PROPLAN',1,'500GR',40000),
('ROYAL CANIN',1,'1KG',100000),
('WHISKAS CAN ADULT TUNA',4,'400GR',23000),
('WHISKAS CAN JUNIOR TUNA',4,'400GR',23000),
('LIFECAT CAN ADULT TUNA',4,'400GR',17000),
('WHISKAS POUCH ADULT TUNA',3,'75GR',7000),
('WHISKAS POUCH JUNIOR TUNA',3,'75GR',7000),
('LIFECAT POUCH ADULT TUNA',3,'75GR',6000),
('KALUNG',5,'-',20000),
('TONGKAT BULU',5,'-',20000),
('BAJU',5,'-',20000);

INSERT INTO payments (payment_method) VALUES
('CASH'),
('BCA'),
('BNI'),
('GOPAY');

INSERT INTO categories (category_name) VALUES 
('DRY FOOD'),
('WET FOOD'),
('WET POUCH'),
('WET CANNED'),
('ACCESORIES'),
('CAT LITTER');