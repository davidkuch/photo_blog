package handlers

//all this for yet uncreated version:
// changes to be apllied include:
// gallery unique id's
// better permissions mechanism

type gallery struct {
	Id           int
	Owner        string
	Name         string
	Date_created string
	Last_updated string
	Size         int
	IsPublic     bool
	Pics         map[string]string
}

//func (gal gallery) Get_data() {

//}

//func (gal gallery) Get_pics_annotations(id int) {
//	gal.Pics = db.Get_pics_annotations(id, "temp")
//}
