CREATE SEQUENCE public.seq_group_barang


create table group_barang 
(
	id bigint not null primary key default nextval('seq_group_barang'::regclass),
	code character varying(50) COLLATE pg_catalog."default",
    name character varying(500) COLLATE pg_catalog."default",
    updated_by character varying(100) COLLATE pg_catalog."default",
    updated_at timestamp with time zone
)
