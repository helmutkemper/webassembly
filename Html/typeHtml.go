package Html

type Html struct{}

func (el Html) NewImage(parent interface{}, propertiesList map[string]interface{}, waitLoad bool) interface{} {
	img := Image{
		SetProperty: propertiesList,
		WaitLoad:    waitLoad,
	}
	img.SetParent(parent)
	img.Create()

	return img.Get()
}
