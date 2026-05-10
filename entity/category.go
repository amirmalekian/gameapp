package entity

type Category struct {
	ID          uint
	Name        string
	Description string
}

//type Category uint8
//
//const (
//	CategorySport Category = iota + 1
//	CategoryHistory
//	CategoryTech
//)
//
//func (c Category) String() string {
//	switch c {
//	case CategorySport:
//		return "Sport"
//	case CategoryHistory:
//		return "History"
//	case CategoryTech:
//		return "Tech"
//	default:
//		return "Unknown"
//	}
//}
