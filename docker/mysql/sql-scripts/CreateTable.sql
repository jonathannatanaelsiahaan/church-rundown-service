CREATE TABLE posts (
id int PRIMARY KEY AUTO_INCREMENT,
title  varchar(35),
content varchar(35)
);

CREATE TABLE organizers (
id int PRIMARY KEY AUTO_INCREMENT,
name varchar(300) NOT NULL UNIQUE,
display_name varchar(300) NOT NULL UNIQUE,
description varchar(35),
location_name varchar(300) NOT NULL,
location_lat varchar(100) NOT NULL,
location_lng varchar(100) NOT NULL,
location_address varchar(200) NOT NULL,
city varchar(100) NOT NULL,
province varchar(200) NOT NULL,
nation varchar(100) NOT NULL
);

CREATE TABLE users (
id int PRIMARY KEY AUTO_INCREMENT,
name varchar(100),
organizer_id int,
FOREIGN KEY (organizer_id) REFERENCES organizers(id)
);

CREATE TABLE accounts (
id int PRIMARY KEY AUTO_INCREMENT,
username varchar(30) NOT NULL UNIQUE,
password varchar(30) NOT NULL,
user_id int,
FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE rundowns (
id int PRIMARY KEY AUTO_INCREMENT,
title varchar(600),
subtitle varchar(600),
show_time datetime,
end_time datetime,
organizer_id int,
FOREIGN KEY (organizer_id) REFERENCES organizers(id)
);

CREATE TABLE rundown_items (
id int PRIMARY KEY AUTO_INCREMENT,
title text,
subtitle text,
text text,
rundown_id int,
FOREIGN KEY (rundown_id) REFERENCES rundowns(id)
);

CREATE TABLE concregation (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(350),
    age int,
    address varchar(250),
    organizer_id int,
    FOREIGN KEY (organizer_id) REFERENCES organizers(id)
);

CREATE TABLE device_inventory (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(400),
    purchase_date datetime,
    total int,
    organizer_id int,
    FOREIGN KEY (organizer_id) REFERENCES organizers(id)
);

CREATE TABLE service_schedule (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(400),
    text text,
    date datetime,
    organizer_id int,
    FOREIGN KEY (organizer_id) REFERENCES organizers(id)
);

CREATE TABLE sector_coordinator (
    id int PRIMARY KEY AUTO_INCREMENT,
    name varchar(400),
    concregation_id int,
    organizer_id int,
    FOREIGN KEY(organizer_id) REFERENCES organizers(id),
    FOREIGN KEY (concregation_id) REFERENCES concregation(id)
);

-- INSERT INTO organizers(name, description) values('Organizer 1', 'Description of organizer 1');
-- INSERT INTO organizers(name, description) values('Organizer 2', 'Description of organizer 2');
-- INSERT INTO organizers(name, description) values('Organizer 3', 'Description of organizer 3');

-- INSERT INTO users(name, organizer_id) values('User 1', 1);
-- INSERT INTO users(name, organizer_id) values('User 2', 2);
-- INSERT INTO users(name, organizer_id) values('User 3', 3);

-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 1', 'Subtitle', '2019-04-04 12:30:00', '2019-04-04 14:00:00', 1);
-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 2', 'Subtitle', '2019-04-04 13:30:00', '2019-04-04 14:00:00', 1);
-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 3', 'Subtitle', '2019-04-04 12:30:00', '2019-04-04 18:00:00', 1);
-- INSERT INTO rundowns(title, subtitle, show_time, end_time, organizer_id) values('Rundowns 4', 'Subtitle', '2019-04-04 15:30:00', '2019-04-04 18:00:00', 1);

-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 1', 'subs', 'text', 1);
-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 2', 'subs', 'text', 1);
-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 3', 'subs', 'text', 1);
-- INSERT INTO rundown_items(title, subtitle, text, rundown_id) values('Rundown item 4', 'subs', 'text', 1);