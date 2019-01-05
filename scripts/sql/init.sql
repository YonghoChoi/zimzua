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
  UNIQUE INDEX ux_email (email),
  PRIMARY KEY (id)
);