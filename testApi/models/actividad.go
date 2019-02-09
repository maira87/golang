type Actividad struct {
   Id       int    `orm:"column(id);pk;auto"`
   Descripcion   string `orm:"column(descripcion)"`
   Fecha_Resgistro string `orm:"column(fecha_resgistro)"`
   Fecha_Limite       string    `orm:"column(fecha_limite)"`
   Soporte   string `orm:"column(soporte)"`
   Estado string `orm:"column(estado)"`
}
