create table public.users
(
    uuid       text primary key         not null default gen_random_uuid(),
    firstname  character varying(50)    not null,
    surname    character varying(50)    not null,
    middlename character varying(50),
    sex        character(1)             not null,
    birth_date timestamp with time zone not null,
    fio        text
);

CREATE OR REPLACE FUNCTION fill_fio() RETURNS TRIGGER AS
$$
declare
    t text;
BEGIN
    t = (SELECT CASE NEW.middlename IS NOT NULL
                    WHEN true THEN concat(' ', NEW.middlename)
                    WHEN false THEN ''
                    END
         from Users);
    UPDATE Users
    SET fio = concat(NEW.firstname, ' ', NEW.surname, t)
    WHERE uuid = NEW.uuid;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER fill_fio AFTER INSERT ON public.users FOR EACH ROW EXECUTE PROCEDURE fill_fio();

alter table users
    owner to postgres;

create table public.products
(
    uuid          text primary key not null default gen_random_uuid(),
    price_in_usd  float            NOT NULL DEFAULT 0,
    left_in_stock INT              NOT NULL DEFAULT 0,
    description   text
);

alter table products
    owner to postgres;

create table public.orders
(
    uuid       text primary key         not null default gen_random_uuid(),
    user_id    character varying(16)    NOT NULL REFERENCES users (uuid),
    created_at timestamp with time zone not null default now()
);

alter table orders
    owner to postgres;

create table public.OrderProducts
(
    order_id   character varying(16)    NOT NULL REFERENCES orders (uuid),
    product_id character varying(16)    NOT NULL REFERENCES products (uuid),
    PRIMARY KEY (order_id, product_id),
    created_at timestamp with time zone not null default now()
);

alter table OrderProducts
    owner to postgres;

create table public.currency
(
    name          character(3) not null primary key,
    course_to_usd float        not null
);

alter table currency
    owner to postgres;

create table public.friends
(
    first_user  character varying(16) NOT NULL REFERENCES users (uuid),
    second_user character varying(16) NOT NULL REFERENCES users (uuid),
    PRIMARY KEY (first_user, second_user)
);

alter table friends
    owner to postgres;
