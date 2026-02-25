package errores

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "Error: Cantidad incorrecta de parámetros"
}

type ErrorEntrada struct{}

func (e ErrorEntrada) Error() string {
	return "Error: Entrada inválida"
}

type ErrorUsuarioLoggeado struct{}

func (e ErrorUsuarioLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioInexistente struct{}

func (e ErrorUsuarioInexistente) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioNoLoggeado struct{}

func (e ErrorUsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorPost struct{}

func (e ErrorPost) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorPostInexistenteSinLikes struct{}

func (e ErrorPostInexistenteSinLikes) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorMostrarLikes struct{}

func (e ErrorMostrarLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
