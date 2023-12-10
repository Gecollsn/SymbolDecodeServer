package google

/*
sdk kits downloader
*/
type AndroidSDK struct {
}

type ISDKDownloadCallback interface {
	OnStart()
	OnProgress()
	OnFailed()
	OnComplete()
}

//check whether sdk is good to go
func (sdk *AndroidSDK) Check() bool {
	return false
}

func (sdk *AndroidSDK) Prepare(callback ISDKDownloadCallback) {
}
