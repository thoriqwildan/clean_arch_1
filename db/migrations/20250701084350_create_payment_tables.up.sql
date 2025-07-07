CREATE TABLE payment_methods (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE,
    description VARCHAR(255) NULL,
    order_num INT DEFAULT 1,
    user_action VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    code VARCHAR(255) NULL
);

CREATE TABLE payment_channels (
    id BIGSERIAL PRIMARY KEY,
    payment_method_id BIGINT,
    code VARCHAR(255) UNIQUE,
    name VARCHAR(255) UNIQUE,
    icon_url VARCHAR(255) NULL,
    order_num BIGINT DEFAULT 1 NULL,
    lib_name VARCHAR(255) NULL,
    user_action VARCHAR(255),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    mdr VARCHAR(255) DEFAULT '0',
    fixed_fee DOUBLE PRECISION DEFAULT 0,
    CONSTRAINT fk_payment_methods_channels FOREIGN KEY (payment_method_id) REFERENCES payment_methods (id) ON UPDATE CASCADE ON DELETE CASCADE
);