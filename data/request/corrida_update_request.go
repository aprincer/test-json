package request

type CorridaUpdateRequest struct {
	Id        int
	Secuencia int `validate :"required min=1, max=1000" json:"name"`
}
