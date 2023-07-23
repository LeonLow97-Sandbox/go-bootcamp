CREATE TABLE IF NOT EXISTS widgets (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    inventory_level INTEGER NOT NULL DEFAULT 100,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);

insert into widgets (name, description, inventory_level, price, created_at, updated_at) values ('Widget', 'A very nice widget.', 10, 1000, now(), now());


CREATE TABLE IF NOT EXISTS transaction_statuses (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);


INSERT INTO transaction_statuses (
    name
  )
VALUES  
 ('Pending'),
 ('Pending'),
 ('Cleared'),
 ('Declined'),
 ('Refunded'),
('Partially refunded');

CREATE TABLE IF NOT EXISTS transactions (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    amount INTEGER NOT NULL,
    currency VARCHAR(6),
    last_four VARCHAR(4),
    bank_return_code VARCHAR(255),
    -- transaction_status_id FOREIGN  KEY TO transaction table
    transaction_status_id INTEGER,
     CONSTRAINT `FK_transactions_transaction_status_id` 
    FOREIGN KEY (transaction_status_id) REFERENCES transaction_statuses(id) 
        ON DELETE SET NULL  ON UPDATE NO ACTION, 

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS statuses (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO statuses (
    name
)
VALUES
('Cleared'),
('Refunded'),
('Cancelled')
;


CREATE TABLE IF NOT EXISTS orders (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    widget_id INTEGER NOT NULL,
    FOREIGN  KEY (widget_id) REFERENCES widgets(id),

    transaction_id INTEGER,
    CONSTRAINT `FK_order_transaction_id`
    FOREIGN  KEY (transaction_id) REFERENCES transactions(id) 
        ON DELETE SET NULL ON UPDATE NO ACTION, 

    status_id INTEGER,
    CONSTRAINT `FK_statuses_status_id`
    FOREIGN  KEY (status_id) REFERENCES statuses(id) 
        ON DELETE SET NULL  ON UPDATE NO ACTION,

    quantity INTEGER,
    amount INTEGER,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
    updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS users (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  email VARCHAR(255),
  password VARCHAR(60),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
  updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO users (
    first_name,
    last_name,
    email,
    password
  )
VALUES  
('Admin',
'User',
'admin@example.com',
'$2a$12$VR1wDmweaF3ZTVgEHiJrNOSi8VcS4j0eamr96A/7iOe8vlum3O3/q'
);

ALTER TABLE widgets ADD COLUMN image VARCHAR(400);
update widgets set image = 'widget.png' where id = 1;

CREATE TABLE customers (
  id INTEGER PRIMARY KEY AUTO_INCREMENT,
  first_name VARCHAR(255),
  last_name VARCHAR(255),
  email VARCHAR(255),

  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,    
  updated_at TIMESTAMP  ON UPDATE CURRENT_TIMESTAMP
);

ALTER TABLE orders ADD COLUMN customer_id INTEGER NOT NULL;
ALTER TABLE orders ADD  CONSTRAINT `FK_order_customer_id`  FOREIGN KEY (customer_id)
 REFERENCES customers(id)   ON UPDATE NO ACTION; 

ALTER TABLE transactions ADD COLUMN expiry_month INT DEFAULT 0;
ALTER TABLE transactions ADD COLUMN expiry_year INT DEFAULT 0;

