package request

type CorridaCreateRequest struct {
	//Secuencia int `validate :"required min=1, max=100" json:"name"`
	Id        int
	Secuencia int
}
