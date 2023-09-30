SHOW TIMEZONE;
set TIMEZONE = 'Asia/Jakarta';

CREATE TABLE mock (
	id serial primary key,
	name VARCHAR ( 255 ) not null,
	method VARCHAR ( 50 ) not null,
	path VARCHAR ( 255 ) unique not null,
	response_code INT not null,
	request JSONB,
	response JSONB not null,
	created_at TIMESTAMP NOT null default NOW(),
	updated_at TIMESTAMP NOT null default NOW()
);
