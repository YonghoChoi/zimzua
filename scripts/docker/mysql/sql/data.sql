insert into storage(name, phone, address, location) values('디캠프', '010-0000-0000', '서울특별시 강남구 역삼동 선릉로 551 새롬빌딩', POINT(127.0396348,37.5063471));
insert into storage(name, phone, address, location) values('에이원스크린골프', '010-0000-0000', '서울특별시 강남구 역삼동 683-26', POINT(127.0437843,37.5080175));
insert into storage(name, phone, address, location) values('컴퓨터구조대', '010-0000-0000', '서울특별시 강남구 역삼1동 683-13', POINT(127.0430344,37.5087952));
insert into storage(name, phone, address, location) values('샵식스플러스', '010-0000-0000', '서울특별시 강남구 역삼1동 693-30', POINT(127.043474,37.5069385));
insert into storage(name, phone, address, location) values('삼성역', '010-0000-0000', '삼성역', POINT(127.060955,37.5088652));
insert into storage(name, phone, address, location) values('강남역', '010-0000-0000', '강남역', POINT(127.032834,37.4969117));

-- 선릉 박물관 37.5084632,127.043695
SET @lon = 127.043695;
SET @lat = 37.5084632;

SELECT *, ST_DISTANCE_SPHERE(POINT(@lon, @lat), location) AS dist
FROM storage
ORDER BY dist;


-- 성능 개선 (참고 https://purumae.tistory.com/198)
SET @lon = 127.043695;
SET @lat = 37.5084632;

SET @MBR_length = 5000;

SET @lon_diff = @MBR_length / 2 / ST_DISTANCE_SPHERE(POINT(@lon, @lat), POINT(@lon + IF(@lon < 0, 1, -1), @lat));
SET @lat_diff = @MBR_length / 2 / ST_DISTANCE_SPHERE(POINT(@lon, @lat), POINT(@lon, @lat + IF(@lat < 0, 1, -1)));

SET @diagonal = CONCAT('LINESTRING(', @lon -  IF(@lon < 0, 1, -1) * @lon_diff, ' ', @lat -  IF(@lon < 0, 1, -1) * @lat_diff, ',', @lon +  IF(@lon < 0, 1, -1) * @lon_diff, ' ', @lat +  IF(@lon < 0, 1, -1) * @lat_diff, ')');

SELECT *, ST_DISTANCE_SPHERE(POINT(@lon, @lat), location) AS dist
FROM storage FORCE INDEX FOR JOIN (`spaidx-storage-location`)
WHERE MBRCONTAINS(ST_LINESTRINGFROMTEXT(@diagonal), location);
