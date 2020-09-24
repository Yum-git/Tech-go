create database TechDojo;
use TechDojo;
create table userdata (id bigint(8) auto_increment, name varchar(200), index(id));

alter table userdata add primary key(id);

create table characterdata (characterid bigint(7) auto_increment, name varchar(100), rarity varchar(100), index(characterid));

alter table characterdata add primary key(characterid);

alter table characterdata rename column name to charactername;

create table possuser (id bigint(7), characterid bigint(7));

create table gachatable (rarity varchar(200), probability double);