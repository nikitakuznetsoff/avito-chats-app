set names utf8;
set time_zone = '+03:00';

drop table if exists messages;
drop table if exists user_chat_relation;
drop table if exists chats;
drop table if exists users;

create table `users` (
    `id` int(16) not null auto_increment,
    `username` varchar(255) not null,
    `created_at` timestamp default current_timestamp,
    primary key(`id`)
) default charset = utf8;

create table `chats` (
    `id` int(16) not null auto_increment,
    `name` varchar(255) not null,
    `created_at` timestamp default current_timestamp,
    primary key(`id`)
) default charset = utf8;

create table `user_chat_relation` (
    `chat_id` int(16) not null,
    `user_id` int(16) not null,
    foreign key (`chat_id`) references chats(`id`),
    foreign key (`user_id`) references users(`id`)
);

create table `messages` (
    `id` int(16) not null auto_increment,
    `chat` int(16) not null,
    `author` int(16) not null,
    `text` varchar(255) not null,
    `created_at` timestamp default current_timestamp,
    foreign key (`chat`) references chats(`id`),
    foreign key (`author`) references users(`id`),
    primary key(`id`)
) default charset = utf8;

