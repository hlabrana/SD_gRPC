package main


func main() {
	var opcion int
	var listaTorres []Torre
	for true {
		opcion = IngresarOpcion()
		if opcion == 1 {
			var torreEscogida Torre
			torreEscogida = SeleccionarTorre(listaTorres)
			ConsultarTorre(torreEscogida)
		} else{
			listaTorres = append(listaTorres,AgregarTorre())
		}
	}
}