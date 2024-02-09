CREATE EXTENSION pgcrypto;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

SET search_path TO public;

CREATE TABLE IF NOT EXISTS transactions (

                                            id uuid DEFAULT public.gen_random_uuid() NOT NULL,
                                            origin VARCHAR(255) NOT NULL,
                                            user_id uuid NOT NULL,
                                            transaction_type character varying,
                                            amount NUMERIC(10, 2) NOT NULL,
                                            created_at timestamp(6) without time zone NOT NULL
);
