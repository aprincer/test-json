package model

import (
	"time"
)

type Sondaje struct {
	Id             int
	Id_proyecto    int
	Nombre_sondaje string
	Fecha_inicio   string
	Fecha_fin      string
	Metros         string
	Este           string
	Norte          string
	Azimut         string
	Elevacion      string
	Inclinacion    string
	Id_geologo     string
	Id_supervisor  int
	S_uuid         string
	Load_Date      time.Time
	Loaded_By      string
	Modified_Date  *time.Time
	Modified_By    *string
	Ts             string
	Corridas       []Corrida
}
