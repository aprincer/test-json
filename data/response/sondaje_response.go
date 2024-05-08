package response

import (
	"test-json/model"
	"time"
)

type SondajeResponse struct {
	Id             int             `json:"id"`
	Id_proyecto    int             `json:"id_proyecto"`
	Nombre_sondaje string          `json:"nombre_sondaje"`
	Fecha_inicio   string          `json:"fecha_inicio"`
	Fecha_fin      string          `json:"fecha_fin"`
	Metros         string          `json:"metros"`
	Este           string          `json:"este"`
	Norte          string          `json:"norte"`
	Azimut         string          `json:"azimut"`
	Elevacion      string          `json:"elevacion"`
	Inclinacion    string          `json:"inclinacion"`
	Id_geologo     string          `json:"id_geologo"`
	Id_supervisor  int             `json:"id_supervisor"`
	S_uuid         string          `json:"s_uuid"`
	Load_Date      time.Time       `json:"load_date"`
	Loaded_By      string          `json:"loaded_by"`
	Modified_Date  *time.Time      `json:"modified_date"`
	Modified_By    *string         `json:"modified_by"`
	Ts             string          `json:"ts"`
	Corridas       []model.Corrida `json:"corridas"`
}
