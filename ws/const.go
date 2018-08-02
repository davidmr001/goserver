package ws

import (
	"github.com/gin-gonic/gin"
	"regexp"
)

const (
	PageSize    = 20
	KvGeekAdmin = "GeekAdmin"

	RpcUserHost = "127.0.0.1:9100"
)

var (
	WebGin = gin.New()

	UaSeo = regexp.MustCompile(`baiduspider|twitterbot|facebookexternalhit|rogerbot|linkedinbot|embedly|quora link preview|showyoubot|outbrain|pinterest|slackbot|vkShare|W3C_Validator`)
	UaH5  = regexp.MustCompile(`(MIDP)|(WAP)|(UP.Browser)|(Smartphone)|(Obigo)|(Mobile)|(AU.Browser)|(wxd.Mms)|(WxdB.Browser)|(CLDC)|(UP.Link)|(KM.Browser)|(UCWEB)|(SEMC\-Browser)|(Mini)|(Symbian)|(Palm)|(Nokia)|(Panasonic)|(MOT\-)|(SonyEricsson)|(NEC\-)|(Alcatel)|(Ericsson)|(BENQ)|(BenQ)|(Amoisonic)|(Amoi\-)|(Capitel)|(PHILIPS)|(SAMSUNG)|(Lenovo)|(Mitsu)|(Motorola)|(SHARP)|(WAPPER)|(LG\-)|(LG/)|(EG900)|(CECT)|(Compal)|(kejian)|(Bird)|(BIRD)|(G900/V1.0)|(Arima)|(CTL)|(TDG)|(Daxian)|(DAXIAN)|(DBTEL)|(Eastcom)|(EASTCOM)|(PANTECH)|(Dopod)|(Haier)|(HAIER)|(KONKA)|(KEJIAN)|(LENOVO)|(Soutec)|(SOUTEC)|(SAGEM)|(SEC\-)|(SED\-)|(EMOL\-)|(INNO55)|(ZTE)|(iPhone)|(Android)|(Windows CE)|(Wget)|(Java)|(curl)|(Opera)/`)
)

func ParamBase(db string) {

	ParamInit(EnvDev, `
UploadDir=./upload/
DownDir=./down/
DbDriver=mysql
DbDsn=root:root@tcp(127.0.0.1:3306)/`+ db+ `?timeout=30s&charset=utf8mb4&parseTime=true
ImgHost=http://127.0.0.1:9000
ImgMaxWidth=800
MultiSite=0
`)

	ParamInit(EnvProd, `
UploadDir=./upload/
DownDir=./down/
DbDriver=mysql
DbDsn=root:root@tcp(127.0.0.1:3306)/`+ db+ `?timeout=30s&charset=utf8mb4&parseTime=true
ImgHost=http://s.ecdiy.cn
ImgMaxWidth=800
MultiSite=0
`)

}
