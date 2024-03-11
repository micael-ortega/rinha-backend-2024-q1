INSERT INTO transactions (value, description, kind, client_id)
VALUES($1, $2, $3, $4);


SELECT u.balance,
    u.account_limit,
    current_timestamp,
    t.value,
    t.kind,
    t.description,
    t.timestamp
FROM users u
    JOIN transactions t ON t.client_id = u.id
WHERE u.id = 1

ORDER BY t.timestamp DESC
LIMIT 10;