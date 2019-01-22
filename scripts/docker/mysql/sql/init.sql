DROP DATABASE IF EXISTS zimzua;
CREATE DATABASE zimzua DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;;
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
)ENGINE = InnoDB
COMMENT = '사용자 계정 테이블';

CREATE TABLE storage (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(256) NOT NULL,
  phone varchar(24) NOT NULL default '010-0000-0000',
  address varchar(256) NOT NULL,
  location point NOT NULL,
  created datetime default now(),
  updated datetime default now(),
  PRIMARY KEY (id)
)ENGINE = InnoDB
COMMENT = '업체 테이블';

CREATE SPATIAL INDEX `spaidx-storage-location` ON storage (location);

GRANT ALL PRIVILEGES ON zimzua.* TO 'zimzua'@'%'
    IDENTIFIED BY 'zimzua'
    WITH GRANT OPTION;
FLUSH PRIVILEGES;

-- 성능 개선 (참고 https://purumae.tistory.com/198)
DELIMITER //
CREATE PROCEDURE GetStorageList(loc POINT)
BEGIN
  SELECT name, phone, address, location, created, updated, ST_DISTANCE_SPHERE(loc, location) AS dist
  FROM zimzua.storage
  ORDER BY dist;
END
//

CREATE PROCEDURE GetStorageListOptimize(loc POINT)
BEGIN
  SET @MBR_length = 1000;

  SET @lon_diff = @MBR_length / 2 / ST_DISTANCE_SPHERE(POINT(@lon, @lat), POINT(@lon + IF(@lon < 0, 1, -1), @lat));
  SET @lat_diff = @MBR_length / 2 / ST_DISTANCE_SPHERE(POINT(@lon, @lat), POINT(@lon, @lat + IF(@lat < 0, 1, -1)));

  SET @diagonal = CONCAT('LINESTRING(', @lon -  IF(@lon < 0, 1, -1) * @lon_diff, ' ', @lat -  IF(@lon < 0, 1, -1) * @lat_diff, ',', @lon +  IF(@lon < 0, 1, -1) * @lon_diff, ' ', @lat +  IF(@lon < 0, 1, -1) * @lat_diff, ')');

  SELECT *, ST_DISTANCE_SPHERE(loc, location) AS dist
  FROM zimzua.storage FORCE INDEX FOR JOIN (`spaidx-storage-location`)
  WHERE MBRCONTAINS(ST_LINESTRINGFROMTEXT(@diagonal), location);
END
//
DELIMITER ;
;