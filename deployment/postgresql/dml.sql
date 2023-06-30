BEGIN;
INSERT INTO medical_order (message) VALUES ('超過120請施打8u');
INSERT INTO medical_order (message) VALUES ('低於80請施打4u');
INSERT INTO medical_order (message) VALUES ('藥物劑量不得超過5u');

INSERT INTO patient (name, order_id) VALUES ('Nathan', 1);
INSERT INTO patient (name, order_id) VALUES ('Emily', 2);
INSERT INTO patient (name, order_id) VALUES ('Rebecca', 3);
INSERT INTO patient (name, order_id) VALUES ('Benjamin', 2);
INSERT INTO patient (name, order_id) VALUES ('Elizabeth', 1);
COMMIT;