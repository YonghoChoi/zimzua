# NSOM Agent (Windows) 설치 가이드

1. [SVN 설치](https://sourceforge.net/projects/win32svn/) (설치가 되어있지 않은 경우에만)

   * 별도 설정 없이 Next 진행
   * 설치 진행 중 Subversion 설치 경로 확인 (default : C:\Program Files (x86)\Subversion)

2. nsom-windows-agent/nsom-windows-agent.exe.config 파일의 변수 설정

   ```
   ... 생략 ...
   <appSettings>
       <add key="REDIS_ADDR" value="[RedisHostAddr]:6379"/>
       <add key="SVN_BIN" value="C:\Program Files (x86)\Subversion\bin\svn.exe"/>
   </appSettings>
   ... 생략 ...
   ```

   * REDIS_ADDR의 RedisHostAddr 부분에 Redis 서버가 설치된 호스트 IP 주소 입력
   * SVN_BIN 경로에 앞서 확인한 Subversion 설치 경로 + bin\svn.exe 입력

3. nsom-windows-agent-installer\setup.bat 실행
