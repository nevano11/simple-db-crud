create table if not exists users
(
    id            serial not null primary key,
    name          text   not null,
    login         text   not null unique,
    password      text   not null,
    password_hash text   not null
);

create table if not exists cars
(
    id         serial  not null primary key,
    brand      text    not null,
    model      text    not null unique,
    generation integer not null,
    price      money   not null
);

create table if not exists usercar
(
    id      serial  not null primary key,
    user_id integer not null,
    car_id  integer not null,

    CONSTRAINT fk_users
        FOREIGN KEY(user_id)
            REFERENCES users(id),
    CONSTRAINT fk_cars
        FOREIGN KEY(car_id)
            REFERENCES cars(id)
)