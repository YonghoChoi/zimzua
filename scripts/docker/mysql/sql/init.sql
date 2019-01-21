DROP DATABASE IF EXISTS zimzua;
CREATE DATABASE zimzua;
USE zimzua;
CREATE TABLE account (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(24) NOT NULL,
  phone varchar(24) NOT NULL default '',
  email varchar(256) NOT NULL,
  password varchar(24) NOT NULL,
  loginType varchar(24) NOT NULL,
  token varchar(256) NOT NULL,
  created datetime default now(),
  updated datetime default now(),
  UNIQUE INDEX ux_email (email),
  PRIMARY KEY (id)
);

CREATE TABLE storage (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(24) NOT NULL,
  phone varchar(24) NOT NULL default '',
  address varchar(256) NOT NULL,
  location point NOT NULL,
  created datetime default now(),
  updated datetime default now(),
  UNIQUE INDEX ux_name (name),
  PRIMARY KEY (id)
);

CREATE SPATIAL INDEX `spaidx-storage-location` ON storage (location);

GRANT ALL PRIVILEGES ON zimzua.* TO 'zimzua'@'%'
    IDENTIFIED BY 'zimzua'
    WITH GRANT OPTION;
FLUSH PRIVILEGES;