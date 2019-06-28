USE zimzua;

CREATE TABLE `tally` (
  `n` SMALLINT UNSIGNED NOT NULL COMMENT '1 ~ 10,000 정수',
  PRIMARY KEY (`n`))
ENGINE = InnoDB
COMMENT = 'mutex 용도의 테이블';

INSERT `tally` VALUES (1),(2),(3),(4),(5),(6),(7),(8),(9),(10);


SET @src = '{
"DESCRIPTION" : {"YCODE":"Y좌표","STOP_NM":"정류소명","STOP_NO":"정류소번호","XCODE":"X좌표"},
"DATA" : [
{"stop_nm":"종로2가사거리","ycode":"37.5697651251","stop_no":"01001","xcode":"126.9877498816"},
{"stop_nm":"창경궁.서울대학교병원","ycode":"37.5791830159","stop_no":"01002","xcode":"126.9965660023"},
{"stop_nm":"명륜3가.성대입구","ycode":"37.5826711749","stop_no":"01003","xcode":"126.9983401004"},
{"stop_nm":"종로2가.삼일교","ycode":"37.5685792736","stop_no":"01004","xcode":"126.9876130976"},
{"stop_nm":"혜화동로터리","ycode":"37.586243","stop_no":"01005","xcode":"127.001744"},
{"stop_nm":"서대문역사거리","ycode":"37.566137","stop_no":"01006","xcode":"126.966893"},
{"stop_nm":"서울역사박물관.경희궁앞","ycode":"37.569135","stop_no":"01007","xcode":"126.97038"}
]
}';

CREATE TEMPORARY TABLE t1
AS
SELECT JSON_EXTRACT(@src, CONCAT('$.DATA[', n - 1, ']')) AS DATA
FROM tally T
WHERE T.n <= JSON_LENGTH(JSON_EXTRACT(@src, '$.DATA'));

INSERT storage (name, address, location, lon, lat)
SELECT DATA ->> '$.stop_nm'
    , DATA ->> '$.stop_nm'
    , POINT(DATA ->> '$.xcode', DATA ->> '$.ycode')
    , CAST((DATA ->> '$.xcode') as DECIMAL(20,10)) as lon
    , CAST((DATA ->> '$.ycode') as DECIMAL(20,10)) as lat
FROM t1;

-- SELECT DATA ->> '$.stop_nm'
--     , DATA ->> '$.stop_nm'
--     , POINT(DATA ->> '$.xcode', DATA ->> '$.ycode')
--     , DATA ->> '$.xcode'
--     , DATA ->> '$.ycode'
-- FROM t1;

DROP TABLE t1;
DROP TABLE tally;

-- 버스 정류소 위치 정보 (http://data.seoul.go.kr/dataList/datasetView.do?infId=OA-15067&srvType=S&serviceKind=1)