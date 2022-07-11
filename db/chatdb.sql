create database if not exists chatdb character set utf8 collate utf8_general_ci;
use chatdb;

create user if not exists 'user'@'localhost' identified by 'hoge';
grant all privileges on chatdb.* to 'user'@'localhost';

create table if not exists members(
id bigint not null auto_increment,
username varchar(128) unique,
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