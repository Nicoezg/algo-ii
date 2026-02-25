import tdas
import conversion_resultado

LIMITE_NAVEGACION = 20
ITERACIONES_LABELS = 50

def lectura(grafo, vertices_a_visitar):
	gr_ent = {}
	pila = tdas.Pila()
	cola = tdas.Cola()

	for v in vertices_a_visitar:
		gr_ent[v] = 0

	for v in vertices_a_visitar:
		for w in grafo.adyacentes(v):
			if w in vertices_a_visitar and w != v:
				gr_ent[w] += 1
	
	for v in vertices_a_visitar:
		if gr_ent[v] == 0:
			cola.encolar(v)
	
	while not cola.esta_vacia():
		v = cola.desencolar()
		pila.apilar(v)
		for w in grafo.adyacentes(v):			
			if w in vertices_a_visitar and w != v:
				gr_ent[w] -= 1
				if gr_ent[w] == 0:
					cola.encolar(w)

	return conversion_resultado.pila_a_lista(pila)
	

def distancia_a_vertices(grafo, origen, destino = None):
	orden = {}
	padres = {}
	visitados = set()
	cola = tdas.Cola()
	orden[origen] = 0
	padres[origen] = None
	visitados.add(origen)
	cola.encolar(origen)

	while not cola.esta_vacia():
		v = cola.desencolar()

		if v == destino:
			return padres, orden
		
		for w in grafo.adyacentes(v):
			if w in visitados:
				continue
			padres[w] = v
			orden[w] = orden[v] + 1
			visitados.add(w)
			cola.encolar(w)

	return padres, orden
	
def conectividad(grafo):
	resultados = []
	visitados = set()
	for v in grafo:
		if v not in visitados:
			obtener_cfcs(grafo, v, visitados, {}, {}, tdas.Pila(), set(), resultados, [0])

	return resultados

def obtener_cfcs(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global): 
	orden[v] = mas_bajo[v] = contador_global[0]
	contador_global[0] += 1
	visitados.add(v)
	pila.apilar(v)
	apilados.add(v)
	for w in grafo.adyacentes(v):
		if w not in visitados:
			obtener_cfcs(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)

		if w in apilados:
			mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])

	if orden[v] != mas_bajo[v]:
		return
	
	nueva_cfc = []
	while True:
		w = pila.desapilar()
		apilados.remove(w)
		nueva_cfc.append(w)
		if v == w:
			break

	cfcs.append(nueva_cfc)

def diametro(grafo):
    ordenes = {}
    padres = {}
    for v in grafo:
        padres[v], ordenes[v]= distancia_a_vertices(grafo,v)
    max = 0
    origen_max = None
    destino_max = None

    for origen, orden_dic in ordenes.items():
        for destino, valor in orden_dic.items():
            if valor >= max:
                max = valor
                origen_max = origen
                destino_max = destino

    res = conversion_resultado.reconstruir_camino(padres[origen_max], destino_max)
    return res, max

def comunidad(grafo, origen):
    cont = 0
    labels = {}
    for v in grafo:
        labels[v] = cont
        cont += 1

    orden = grafo.obtener_vertices()
    entrantes = aristas_entrantes(grafo)

    for _ in range(ITERACIONES_LABELS):
        for v in orden:
            labels[v] = _max_freq(entrantes[v],labels)

    comunidad = []
    comunidad_de_origen = labels[origen]

    for vertice in orden:
        if labels[vertice] == comunidad_de_origen:
            comunidad.append(vertice)

    return comunidad

def _max_freq(vertices,labels):
    cont = {}
    for vertice in vertices:
        frecuencia = cont.get(labels[vertice],0)
        cont[labels[vertice]] = frecuencia+1

    label_popular = None
    for label, cantidad in cont.items():
        if label_popular is None or cantidad > cont[label_popular]:
            label_popular = label

    return label_popular

def aristas_entrantes(grafo):
	entrantes = {}
	for v in grafo.obtener_vertices():
		entrantes[v] = []

	for v in grafo.obtener_vertices():
		for w in grafo.adyacentes(v):
			entrantes[w].append(v)

	return entrantes

def navegacion(grafo, origen):
	camino = []
	camino.append(origen)
	act = origen
	for i in range(LIMITE_NAVEGACION):
		if len(grafo.adyacentes(act)) == 0:
			break

		siguiente = grafo.adyacentes(act)[0]
		camino.append(siguiente)
		act = siguiente
	return camino

def coeficiente_clustering(grafo,vertice):
	aux = set()
	cont = 0.000
	adyacentes = len(grafo.adyacentes(vertice))

	for v in grafo.adyacentes(vertice):
		if v == vertice:
			adyacentes -= 1
			continue
		aux.add(v)
		
	if adyacentes < 2:
		return cont
	
	for v in grafo.adyacentes(vertice):
		if v == vertice:
			continue
		for w in grafo.adyacentes(v):
			if w in aux and w != v:
				cont += 1
				
	coeficiente = adyacentes *(adyacentes -1)
		
	return cont / coeficiente
