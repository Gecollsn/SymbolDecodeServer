package dlabt

type DLProgress interface {
	OnStart()
	OnProcess(progress uint8)
	OnComplete()
	OnFailure(err error)
}

type INetdl interface {
	HttpDL(url string, dlpath string, callback DLProgress)
}

type Netdl struct{}

func (ndl *Netdl) HttpDL(url string, dlpath string, callback DLProgress) {

}
