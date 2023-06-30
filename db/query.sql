-- name: GetAll :many
SELECT
    patient.id AS patient_id,
    patient.name AS patient_name,
    medical_order.message AS order_message
FROM
    patient
        JOIN medical_order ON patient.order_id = medical_order.id;
