package DataStructure

type Container struct {
	CId string `json:"container_id"`
	//UpperId string `json:"container_upperGroupId"`
	CName string `json:"container_name"`
	//ps[]Property
	Ps []Property `json:"container_properties"`
}
