package utils

type IDlProgress interface {
	OnStart()
	OnProcess(progress uint8)
	OnComplete()
	OnFailure(err error)
}

type INetdl interface {
	HttpDL(url string, dstPath string, callback IDlProgress)
}

type Netdl struct{}

func (ndl *Netdl) HttpDL(url string, dstPath string, callback IDlProgress) {

}
