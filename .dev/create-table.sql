create schema if not exists greatest_users;

use greatest_users;

create table if not exists users (
    id varchar(100) not null, 
    legacy_login varchar(100) not null, 
    legacy_id numeric(10, 0) not null,
    legacy_node_id varchar(100) not null, 
    legacy_url varchar(300) not null, 
    legacy_html_url varchar(300) not null, 
    new_email varchar(100) not null, 
    created_at timestamp(6) not null,
    primary key (id)
);

create user if not exists user_service identified by '123';
grant all on users to user_service;
