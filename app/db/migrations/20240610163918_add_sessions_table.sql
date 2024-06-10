-- +goose Up
create table if not exists sessions(
	id serial primary key,
	user_id serial not null,
	ip_address text,
	user_agent text,
	expires_at timestamp with time zone,
	last_login_at timestamp with time zone
);

-- +goose Down
drop table if exists sessions;
