CREATE TABLE public.weather_requests
(
    id bigserial NOT NULL,
    lat numeric(19,5) NOT NULL,
    lon numeric(19,5) NOT NULL,
    created timestamp without time zone NOT NULL,
    is_complete boolean NOT NULL,
    is_in_progress boolean NOT NULL,
    CONSTRAINT weather_requests_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.weather_requests
    OWNER to postgres;

CREATE TABLE public.weather_responses
(
    id bigserial NOT NULL,
    request_id bigint NOT NULL,
    temperature integer NOT NULL,
    humidity integer NOT NULL,
    pressure integer NOT NULL,
    is_succeeded boolean NOT NULL,
    CONSTRAINT weather_responses_pkey PRIMARY KEY (id),
    CONSTRAINT "weather_responses_request_id_FK" FOREIGN KEY (request_id)
        REFERENCES public.weather_requests (id) MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.weather_responses
    OWNER to postgres;