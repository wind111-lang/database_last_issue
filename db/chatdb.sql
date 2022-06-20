create database if not exists chatdb character set utf8 collate utf8_general_ci;
use chatdb;

create user if not exists 'user'@'localhost' identified by 'hoge';
grant all privileges on chatdb.* to 'user'@'localhost';

create table if not exists members(
id bigint not null auto_increment,
username varchar(128),
password varchar(128),
birthday char(8),
primary key(id)
);

create table if not exists chat_log(
id bigint not null auto_increment,
username varchar(128),
text text,
primary key(id)
);

create table if not exists azureapi(
subkey varchar(128),
location varchar(16),
endpoint varchar(128),
uri varchar(32)
);

insert into azureapi (subkey,location,endpoint,uri) values('<your-subscription-key>','japaneast','https://api.cognitive.microsofttranslator.com/','/translate?api-version=3.0')