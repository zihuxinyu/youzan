package response

type Regions struct {

	Regions []CommonRegion `json:"regions"` //区域列表

}

type CommonRegion struct {
	Id       int64     `json:"id"`      //  区域ID
	Name     string     `json:"name"`   //  区域全称
	ParentId int64   `json:"parent_id"` //  区域父级ID
}
