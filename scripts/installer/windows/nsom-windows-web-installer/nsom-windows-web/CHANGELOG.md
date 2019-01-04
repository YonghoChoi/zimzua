## v1.0.4 (2018-07-30)
### 기능 변경
* Agent 업데이트 정보에서 업데이트 체크 주기 제거

## v1.0.3 (2018-07-23)
### 기능 개선
* Agent 업데이트 방식 변경
    * 기존에는 Redis에 binary 저장
    * 변경된 버전에서는 NSOM WEB으로 부터 nsom-agent 다운로드
    * 실행 파일 뿐만 아니라 agent 디렉토리 전체 전달
    * 업데이트 기능 개선
        * 모든 호스트 업데이트
        * 특정 호스트 업데이트
        * 업데이트 경로 및 아카이브 파일명 변경 가능
        * 에이전트에서 업데이트 체크를 위한 주기 변경 가능

## v1.0.2 (2018-07-18)
### 기능 개선
* 파일 복사 시 압축 유형 추가 지원
    * 기존에는 Zip 만 지원
    * 7zip, bzip2 추가 지원
* 유형이 추가됨에 따라 라디오박스로 Y, N만 선택하던 UI를 셀렉트박스로 변경
* 압축 옵션 기능 추가 (Fastest/Fast/Normal/Maximum/Ultra)

## v1.0.1 (2018-05-28)
### 기능 추가
* Reverse Proxy 사용 여부 선택 가능하도록 수정
	* 현재 Reverse Proxy로 Nginx를 사용하고 있는데 정적파일 전송 기능만 사용하고 있기 때문에 굳이 사용하지 않아도 될 것으로 판단됨
	* 코드 상에 true/false로 사용 여부 지정
* 로깅 패키지 변경
	* log -> geb/log(logrus)
* NSOM WEB과 AGENT에서 공통으로 사용될 코드를 위해 common 디렉토리 생성
* 공통 모듈 작업
	* Interval 타이머 기능 추가
	* Custom Process 관련 struct 추가


### 기능 개선
* Group 페이지에서 Custom Process 인자 수정 중 일 경우 실행 방지
* Custom Process 기능 개선
	* exec 명령 실행과 프로세스 종료를 제외하고 Pub/Sub 방식 제거
	* Group 페이지 진입 시 글로벌 타이머 동작 (5초 주기로 Redis에 기록된 CustomProcess 체크)