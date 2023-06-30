CREATE TABLE IF NOT EXISTS medical_order (
    id      SERIAL PRIMARY KEY,
    message varchar(255)
);

CREATE TABLE IF NOT EXISTS patient (
    id       SERIAL PRIMARY KEY,
    name     varchar(255),
    order_id int,
    CONSTRAINT patient_order_id FOREIGN KEY (order_id) REFERENCES medical_order(id)
);