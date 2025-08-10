INSERT INTO subscriptions(service_name, price, user_id, start_date, end_date)
VALUES (:service_name, :price, :user_id, :start_date, :end_date)
RETURNING *
