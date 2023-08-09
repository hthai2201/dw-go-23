

drop table if exists users;
create table users
(
    id         int auto_increment primary key,
    email      varchar(50)                                                                  not null,
    password   varchar(50)                                                                  not null,
    last_name  varchar(50)                                                                  not null,
    first_name varchar(50)                                                                  not null,
    phone      varchar(20)                                                                  null,
    roles      enum ('user', 'admin') default 'user'                                        not null,
    salt       varchar(50)                                                                  null,
    avatar     json                                                                         null,
    status     int                    default 1                                             not null,
    created_at timestamp              default current_timestamp                             null,
    updated_at timestamp              default current_timestamp on update current_timestamp null,
    unique (email)
);

