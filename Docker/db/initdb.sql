use chatdb;

create table if not exists members(
id bigint not null auto_increment,
username varchar(128),
password varchar(128),
birthday char(10),
primary key(id)
);

create table if not exists chat_log(
id bigint not null auto_increment,
username varchar(128),
text text,
primary key(id)
);
