CREATE TABLE products (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255),
    body_html TEXT,
    vendor VARCHAR(255),
    product_type VARCHAR(255),
    handle VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    published_at DATETIME,
    published_scope VARCHAR(255),
    tags VARCHAR(255),
    status VARCHAR(255),
    template_suffix VARCHAR(255),
    metafields_global_title_tag VARCHAR(255),
    metafields_global_description_tag VARCHAR(255),
    admin_graphql_api_id VARCHAR(255)
);

CREATE TABLE product_options (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT UNSIGNED,
    name VARCHAR(255),
    value VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE variants (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT UNSIGNED,
    title VARCHAR(255),
    sku VARCHAR(255),
    position INT,
    option1 VARCHAR(255),
    option2 VARCHAR(255),
    option3 VARCHAR(255),
    price DECIMAL(10, 2),
    inventory_quantity INT,
    weight DECIMAL(10, 2),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE images (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT UNSIGNED,
    src VARCHAR(255),
    width INT,
    height INT,
    position INT,
    alt VARCHAR(255),
    FOREIGN KEY (product_id) REFERENCES products(id)
);

CREATE TABLE metafields (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    product_id BIGINT UNSIGNED,
    namespace VARCHAR(255),
    key VARCHAR(255),
    value TEXT,
    value_type VARCHAR(50),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
