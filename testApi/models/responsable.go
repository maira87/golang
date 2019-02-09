type Responsable struct {
   Id       int    `orm:"column(id);pk;auto"`
   Nombre   string `orm:"column(nombre)"`
   Correo string `orm:"column(correo)"`
   Telefono   int `orm:"column(telefono)"`
}
