-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
Create table IF NOT EXISTS products(
	id BIGSERIAL primary key,
	name text NOT NULL,
	price_in_centers INTEGER NOT NULL CHECK (price_in_centers > 0),
	quantity INTEGER NOT NULL DEFAULT 0,
	created_at TIMESTAMPTZ NOT NULL DEFAULT now()
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;
-- +goose StatementEnd
