CREATE TABLE `posts` (
  `id` int PRIMARY KEY Auto_Increment,
  `title` varchar(255),
  `content` mediumtext,
  `category` varchar(255),
  `created_at` timestamp default current_timestamp,
  `updated_at` timestamp default current_timestamp on update current_timestamp
);

CREATE TABLE `tags` (
  `id` int PRIMARY KEY Auto_Increment,
  `tag_name` varchar(255) UNIQUE NOT NULL
);

CREATE TABLE `post_tags` (
  `post_id` int NOT NULL,
  `tag_id` int NOT NULL,
  primary key (post_id, tag_id),
  foreign key (post_id) references posts(id) on delete cascade,
  foreign key (tag_id) references tags(id) on delete cascade
);
