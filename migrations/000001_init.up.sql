-- CREATE SCHEMA public AUTHORIZATION sx_user;

CREATE TYPE account_statuses AS ENUM (
    'ACTIVATING',
    'ENABLED',
	'BLOCKED',
	'REMOVED');

CREATE TYPE user_types AS ENUM (
    'SYSTEM',
    'USER');

CREATE TYPE user_statuses AS ENUM (
    'DRAFT',
    'ENABLED',
    'BLOCKED',
    'REMOVED');

CREATE TABLE auth_accounts (
	account_id uuid NOT NULL DEFAULT gen_random_uuid(),
    datetime_created timestamp NOT NULL,
    account_status account_statuses NOT NULL,
    datetime_status timestamp DEFAULT NULL,
	CONSTRAINT accounts_pkey PRIMARY KEY (account_id)
);

CREATE TABLE auth_users (
	user_id integer NOT NULL GENERATED ALWAYS AS IDENTITY,
    account_id uuid NOT NULL,
    visitor_id bigint NOT NULL CHECK(visitor_id >= 0),
    user_type user_types NOT NULL,
    datetime_created timestamp NOT NULL,
    user_login varchar(32) NOT NULL,
    user_email varchar(128) NOT NULL,
    datetime_last_login timestamp DEFAULT NULL,
    user_last_login_ip integer DEFAULT NULL,
    datetime_last_visit timestamp DEFAULT NULL,
    user_status user_statuses NOT NULL,
    datetime_status timestamp NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (user_id),
    CONSTRAINT fk_auth_account FOREIGN KEY(account_id)
                               REFERENCES auth_accounts(account_id)
);