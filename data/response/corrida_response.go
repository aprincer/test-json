package response

import (
	"database/sql"
	"time"
)

type CorridaResponse struct {
	//Col_name  string
	//Data_type string
	Id                         int           `json:"id"`
	Id_sondaje                 int           `json:"id_sondaje"`
	Secuencia                  int           `json:"secuencia"`
	Id_litologia_general       float32       `json:"id_litologia_general"`
	Cod_lit_general            string        `json:"cod_lit_general"`
	Id_litologia_especifica    int           `json:"id_litologia_especifica"`
	Cod_lit_especifica         string        `json:"cod_lit_especifica"`
	Longitud_triturada         float32       `json:"Longitud_triturada"`
	Id_tramo_triturado         *int          `json:"id_tramo_triturado"`
	Cod_origen_tramo_triturado *string       `json:"cod_origen_tramo_triturado"`
	Id_resistencia             *int          `json:"id_resistencia"`
	Cod_resistencia            *string       `json:"cod_resistencia"`
	Desde                      string        `json:"desde"`
	Hasta                      string        `json:"hasta"`
	Id_cond_subterranea        *int          `json:"id_cond_subterranea"`
	Cod_cond_subterranea       *string       `json:"cod_cond_subterranea"`
	Recuperado_m               string        `json:"recuperado_m"`
	Rqd_m                      string        `json:"rqd_m"`
	Num_fracturas              int           `json:"num_fracturas"`
	Id_jn                      *int          `json:"id_jn"`
	Cod_jn                     *string       `json:"cod_jn"`
	Id_ja                      sql.NullInt32 `json:"id_ja"`
	Cod_ja                     *string       `json:"cod_ja"`
	Id_jr                      *int          `json:"id_jr"`
	Cod_jr                     *string       `json:"cod_jr"`
	Estructura                 *string       `json:"estructura"`
	Surface_cond               *string       `json:"surface_cond"`
	Gsi_aprox                  *string       `json:"gsi_aprox"`
	Comentario                 *string       `json:"comentario"`
	Id_abertura                *int          `json:"id_abertura"`
	Id_rugosidad               *int          `json:"id_rugosidad"`
	Id_relleno                 *int          `json:"id_relleno"`
	Id_meteorizacion           *int          `json:"id_meteorizacion"`
	C_uuid                     string        `json:"c_uuid"`
	Suelo                      int           `json:"suelo"`
	Load_Date                  time.Time     `json:"Load_Date"`
	Loaded_By                  string        `json:"loaded_By"`
	Modified_Date              *time.Time    `json:"modified_Date"`
	Modified_By                *string       `json:"modified_By"`
	Ts                         *string       `json:"ts"`
}
