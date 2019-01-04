# ZIM ZUA

## 디렉토리 구조
* bin : 빌드 출력물
* cmd : 어플리케이션 패키지
* internal : 내부 패키지
* scripts : 스크립트
    * docker : 테스트 환경을 위한 Dockerfile
    * kubernetes : 테스트 환경을 위한 Kubernetes 설정 파일
    * installer : 설치 자동화 스크립트



## 프로젝트 빌드
* api
    ```
    make.bat /app api /platform linux build
    ```

* Debug 모드
    ```
    make.bat /app api /platform linux /debug true build
    ```


