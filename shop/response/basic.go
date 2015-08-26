package response

type Basic struct {

	Sid  string `json:"sid"`  //  必须：是	店铺ID
	Name string `json:"name"` //  必须：是	店铺名称
	Logo string `json:"logo"` //  必须：是	店铺LOGO

}
