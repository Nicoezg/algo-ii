def pila_a_lista(pila):
    lista = []
    while not pila.esta_vacia():
        lista.append(pila.desapilar())
    return lista

def vertices_a_distancia_n(n, orden):
	contador = 0
	for distancia in orden.values():
		if distancia == n:
			contador += 1

	return contador

def reconstruir_camino(padres, w):
	camino = []
	while w:
		camino.append(w)
		w = padres[w]

	return camino[::-1]