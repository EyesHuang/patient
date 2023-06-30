-- name: GetAllPatients :many
SELECT
    patient.id AS patient_id,
    patient.name AS patient_name,
    medical_order.id AS order_id,
    medical_order.message AS order_message
FROM
    patient
        JOIN medical_order ON patient.order_id = medical_order.id;

-- name: GetOrderByID :one
SELECT * FROM medical_order
WHERE id = @id;

-- name: UpdateOrder :exec
UPDATE medical_order
SET message = @message
WHERE id = @id;
