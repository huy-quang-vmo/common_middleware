CREATE TABLE service_managements (
    id INT PRIMARY KEY,
    status VARCHAR(50) NOT NULL UNIQUE
);

INSERT INTO service_managements(id, status) VALUES
(1, 'Active');