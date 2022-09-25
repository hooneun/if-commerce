
CREATE TABLE users (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(255) NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,
  deleted_at TIMESTAMP NULL,
  UNIQUE KEY(email)
);

CREATE TABLE user_addresses (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  address VARCHAR(100) NOT NULL,
  detail VARCHAR(100) NOT NULL,

  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE points (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  point INT NOT NULL
);

create table point_transaction(
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  point_id BIGINT NOT NULL,
  title VARCHAR(20) NOT NULL,
  point INT NOT NULL,

  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,

  FOREIGN KEY (point_id) REFERENCES points(id)
);

CREATE TABLE orders (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  user_id BIGINT NOT NULL,
  user_address_id BIGINT NOT NULL,
  status VARCHAR(50) NOT NULL COMMENT '',
  product_amount INT NOT NULL,
  delivery_amount int NOT NULL,
  total_amount INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (user_address_id) REFERENCES user_addresses(id)
);

CREATE TABLE order_products (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  order_id BIGINT NOT NULL,
  product_id BIGINT NOT NULL,
  name VARCHAR(100) NOT NULL,
  status VARCHAR(50) NOT NULL COMMENT '',

  delivery_date DATE NOT NULL,

  amount int NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE payments (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  order_id BIGINT NOT NULL,
  product_amount INT NOT NULL,
  delivery_amount INT NOT NULL,
  total_amount INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (order_id) REFERENCES orders(id)
);

CREATE TABLE products (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  item_product_id BIGINT NOT NULL,
  name VARCHAR(100) NOT NULL,
  thumbnail VARCHAR(200) NOT NULL,
  detail TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,

  FOREIGN KEY (item_product_id) REFERENCES item_product(id)
);

CREATE TABLE product_comment(
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id BIGINT NOT NULL,
  product_id BIGINT NOT NULL,
  content TEXT NOT NULL,
  score INT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NULL,

  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE items (
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  quantity INT NOT NULL DEFAULT 0,
  price INT NOT NULL DEFAULT 0,
  type VARCHAR(20) NOT NULL COMMENT 'MAIN, OPTION',
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL
);

CREATE TABLE item_product(
  id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
  item_id BIGINT NOT NULL,
  product_id BIGINT NOT NULL,
  price INT NOT NULL,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL,

  FOREIGN KEY (item_id) REFERENCES items(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);
