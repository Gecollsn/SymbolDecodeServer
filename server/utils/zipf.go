package utils

type IZipToolbox interface {
	Zip(folder string, ozip string)
	Unzip(zfile string, ofolder string)
	List(zfile string)
	Extract(zfile string, depth int, name string, ofolder string)
}

type ZipToolbox struct{}

func (ZipToolbox) Zip(folder string, ozip string) {

}

func (ZipToolbox) Unzip(zfile string, ofolder string) {

}

func (ZipToolbox) List(zfile string) {

}

func (ZipToolbox) Extract(zfile string, depth int, name string, ofolder string) {

}
