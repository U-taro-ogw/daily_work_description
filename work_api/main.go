package main

//import "time"

//const location = "Asia/Tokyo"

func main() {
	//loc, err := time.LoadLocation(location)
	//if err != nil {
	//	loc = time.FixedZone(location, 9*60*60)
	//}
	//time.Local = loc

	a := App{}
	a.Initialize()
	a.Run(":8080")
}

//func init() {
//	loc, err := time.LoadLocation(location)
//	if err != nil {
//		loc = time.FixedZone(location, 9*60*60)
//	}
//	time.Local = loc
//}