package version

var AppVersion string // AppVersion Application의 Makefile에서 빌드시 git 리비전과 빌드일자를 조합해서 설정
func GetVersion() string {
	return "0.0.1"
}
