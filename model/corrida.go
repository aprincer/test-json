package model

import (
	"database/sql"
	"time"
)

type Corrida struct {
	//Col_name  string
	//Data_type string

	Id                         int
	Id_sondaje                 int
	Secuencia                  int
	Id_litologia_general       float32
	Cod_lit_general            string
	Id_litologia_especifica    int
	Cod_lit_especifica         string
	Longitud_triturada         float32
	Id_tramo_triturado         *int
	Cod_origen_tramo_triturado *string
	Id_resistencia             *int
	Cod_resistencia            *string
	Desde                      string
	Hasta                      string
	Id_cond_subterranea        *int
	Cod_cond_subterranea       *string
	Recuperado_m               string
	Rqd_m                      string
	Num_fracturas              int
	Id_jn                      *int
	Cod_jn                     *string
	Id_ja                      sql.NullInt32
	Cod_ja                     *string
	Id_jr                      *int
	Cod_jr                     *string
	Estructura                 *string
	Surface_cond               *string
	Gsi_aprox                  *string
	Comentario                 *string
	Id_abertura                *int
	Id_rugosidad               *int
	Id_relleno                 *int
	Id_meteorizacion           *int
	C_uuid                     string
	Suelo                      int
	Load_Date                  time.Time
	Loaded_By                  string
	Modified_Date              *time.Time
	Modified_By                *string
	Ts                         *string
	Fracturas                  []Fractura
}
