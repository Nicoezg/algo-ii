from collections import deque


# Se cuenta con un grafo en el que sus aristas tiene peso 1 o 2, únicamente. 
# Implementar un algoritmo que permita obtener el camino mínimo de un vértice hacia todos los demás, en tiempo 
# O (V + E)


def camino_minimo_limitado(grafo):
    nuevo = Grafo(pesado = False)
    peso_total = 0
    visitados = {}
    for v in grafo: # O (V)
        nuevo.agregar(v)
        
    for v in grafo: # O(V + E) dicc de dicc, todo constante
        visitados.add(v)

        for w in grafo.adyacentes(v):

            if w not in visitados:
                act = v

                if grafo.peso(v,w) == 2:
                    nuevo_vertice = 0
                    nuevo.agregar_vertice(nuevo_vertice)
                    nuevo.agregar_arista(act, nuevo_vertice)
                    act = nuevo_vertice
                nuevo.agregar_arista(act, w)
    distancias = bfs(nuevo) # O(V + E)
    return distancias

# Se resuelve en O (V + E), Por cada vertice veo sus aristas.

# Implementar una función que permita determinar si un grafo puede ser no dirigido. 
# Determinar el orden del algoritmo implementado.

def dirigido_o_no(grafo):
    aristas = obtener_aristas(grafo)
    for arista in aristas:
        if arista[::-1] not in aristas:
            return False
    return True
    
def obtener_aristas(grafo):
    resultado = set()
    for v in grafo:
        for w in grafo.adyacentes(v):
                resultado.add([v,w])
    return resultado


# Escribir una función que detecte y retorne los puntos de articulación de un árbol pasado por parámetro 
# (es decir, un grafo con características de árbol). Indicar y justificar la complejidad de la función.

def puntos_art_arbol(grafo):
    res = []
    for v in grafo:
        if len(grafo.adyacentes(v)) >= 2:
            res.append(v)
    return res


# José se quiere subir al hype de Marvel entonces quiere ver todas las películas. Como le dijeron que no se pueden ver en
#cualquier orden, preguntó y le indicaron cuales referencian a otras para que las vea en un orden que entienda 
# las referencias. Por ejemplo: Avengers: End Game tiene referencias de, entre otras, las películas de Iron Man y 
# Capitán America, entonces tiene que ver estas ultimas antes de End Game. Ademas de querer entender todas, quiere 
# que si da igual cual película ver primero entonces verlas en orden temporal en el cual sucede la película. 
# Por ejemplo para ver Los Guardianes de la Galaxia no es necesario ver Capitana Marvel ni viseversa, 
#entonces quiere ver Capitana Marvel primero que sucede antes que Los Guardianes de la Galaxia. 
# Suponer que se cuenta con un Diccionario que dada una película nos indica la fecha en la que transcurre.

def marvel(grafo, dicc):
    heap = Heap(b - a) #Heap de minimos de fecha de salida
    res = []
    g_ent = {}

    for v in grafo: # O ((V + E) log V )
        g_ent[v] = len(grafo.adyacentes(v))
        if len(grafo.adyacentes(v)) == 0:
            _, año = dicc[v]
            heap.Encolar((v, año))

    while not heap.EstaVacia():
        v, _ = heap.Desencolar()
        res.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                _, año = dicc[w]
                heap.Encolar((w, año))
    return res

    

def topologico_grados(grafo):
  g_ent = grados_entrada(grafo)
  q = deque()
  for v in grafo:
    if g_ent[v] == 0:
      q.append(v)
  resultado = []

  while q: # O(V + E) SOLO SI NO TIENE CICLOS.
    v = q.popleft()
    for w in grafo.adyacentes(v):
      g_ent[w] -= 1
      if g_ent[w] == 0:
        q.append(w)
  if len(resultado) != len(grafo):
    raise ValueError("Hay un ciclo")
  return resultado



# 3. Implementar un algoritmo que dado un grafo y un árbol verifique si el árbol es un MST del grafo. 
# Considerar todas las condiciones que son necesarias para que eso sea el caso. Indicar y justificar la complejidad.

def vertificar_arbol_mst(grafo, arbol): # O(V + E log V)
    visitados = set()
    peso_arbol = 0
    peso_mst = 0
    if len(arbol) != len(grafo): # O(1)
        return False
    
    if not es_conexo(arbol): # O (V + E)
        return False
    
    if len(arbol)-1 != obtener_aristas(arbol): # O (V + E)
        return False 

    
    for v in arbol: # O (V + E)
        visitados.add(v)
        for w in arbol.adyacentes(v):
            if not grafo.estan_unidos(v,w):
                return False
        
            peso_arbol += arbol.peso_arista(v,w)

    v = grafo.obtener_vertice_aleatorio()
    visitados = set()
    visitados.add(v)
    q = Heap()

    for w in grafo.adyecentes(v): # Agrego todos las aristas al heap  O(V + E log V)
        q.encolar((v,w), grafo.peso(v,w))
    mst = Grafo(es_dirigido = False, lista_vertices= grafo.obtener_vertices() )

    while q: 
        (v, w), peso = q.desencolar()
        if w in visitados: # Si ya visite al punto destino, estaría completando un ciclo
            continue
        mst.agregar_arista(v, w, peso) # Si no visite al punto destino, agrego la arista
        peso_mst += peso
        visitados.add(w)

        for x in grafo.adyacentes(w): # Encolo todas las aristas del vertice adyacente cuyos puntos destino no hayan sido visitados
            if x not in visitados:
                q.encolar((w,x), grafo.peso(w,x))
    
    return peso_arbol == peso_mst

    

def es_conexo(grafo): #O(V + E)
    if len(grafo) == 0:
        return True
    visitados = set()
    cola = deque()
    origen = grafo.vertice_aleatorio()
    visitados.add(origen)
    cola.append(origen)
    while len(cola) != 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                cola.append(w)
                visitados.add(w)

    return len(visitados) == len(grafo)

def obtener_aristas(arbol):
    visitados = {}
    contador = 0
    for v in arbol:
        for w in arbol.adyacentes(v):
            contador += 1
    return contador // 2

def bfs(grafo): # O (V + E)
    orden = {}
    padres = {}
    visitados = set()
    q = deque()
    origen = grafo.vertice_aleatorio()
    orden[origen] = 0
    padres[origen] = None
    visitados.add(origen)
    q.append(origen)

    while q:
        v = q.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                q.append(w)
    return padres

def dfs(grafo): # O (V + E)
    orden = {}
    padres = {}
    visitados = set()
    origen = grafo.vertice_aleatorio()
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)

    _dfs(grafo, orden, padres, visitados, origen)

def _dfs(grafo, orden, padres, visitados, v):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            orden[w] = orden[v] + 1
            padres[w] = v
            visitados.add(w)
            _dfs(grafo, orden, padres, visitados, w)

def bfs_topologico(grafo): # O (V + E)
    orden = {}
    cola = deque()
    res = []
    for v in grafo:
        orden[v] = len(grafo.adyacentes(v))
        if orden[v] == 0:
            cola.append(v)
    
    while cola:
        v = cola.popleft()
        res.append(v)

        for w in grafo.adyacentes(v):
            orden[w] -= 1
            if orden[w] == 0:
                cola.append(w)
    return res
        
def dijkstra(grafo): # O(V + E log V)
    dist = {}
    padres = {}
    heap = Heap(b-a)
    origen = grafo.obtener_aleatorio()
    for v in grafo:
        dist[v] = float("inf")
    
    dist[v] = 0
    padres[v] = None
    heap.Encolar((v, 0))

    while not heap.esta_vacio():
        v, _ = heap.Desencolar()
        for w in grafo.adyacentes(v):
            posible_dist = dist[v] + grafo.peso_arista(v,w)
            if posible_dist < dist[w]:
                dist[w] = posible_dist
                padres[w] = v
                heap.Encolar((w, dist[w]))

    return dist

def bellman_ford(grafo): # O(V x E)
    distancia = {}
    padres = {}
    for v in grafo:
        distancia[v] = float("inf")

    origen = grafo.vertice_aleatorio()
    distancia[origen] = 0
    padres[origen] = None

    aristas = obtener_aristas(grafo)
    for i in range(len(grafo)):
        cambio = False
        for origen, destino, peso in aristas:
            if distancia[origen] + peso < distancia[destino]:
                cambio = True
                distancia[destino] = distancia[origen] + peso
                padres[destino] = origen
            
        if not cambio:
            return distancia, padres
            
    for v, w, peso in aristas:
        if distancia[v] + peso < distancia[w]:
            return None
    return distancia, padres

def prim(grafo): # O (V + E log V)
    heap = Heap()
    visitados = set()
    nuevo = Grafo()
    for v in grafo:
        nuevo.agregar_vertice(v)
    
    origen = grafo.obtener_aleatorio()
    visitados.add(origen)

    for w in grafo.adyacentes(origen):
        heap.Encolar(((origen, w) , grafo.peso_arista(v,w)))
    
    while not heap.EstaVacia():
        v, w, dist = heap.Desencolar()
        if w not in visitados:
            nuevo.agregar_arista((v,w),dist)

            for x in grafo.adyacentes(w):
                if x not in visitados:
                    heap.Encolar(((w, x), grafo.peso_arista(w, x)))
            visitados.add(w)

    return nuevo

def kruskal(grafo): # O(V + E log V) depende del algoritmo usado para ordenar.
    nuevo = Grafo()
    conjuntos = UnionFind()
    aristas = []
    for v in grafo:
        nuevo.agregar_vertice(v)
        for w in grafo.adyacentes(v):
            aristas.append((v,w), grafo.peso_arista(v,w))
    
    aristas = sorted(aristas)

    for arista in aristas:
        v, w, peso = arista
        if conjuntos.find(v) != conjuntos.find(w):
            nuevo.agregar_arista((v,w), peso)
            conjuntos.union(v,w)
    
    return nuevo

        






        
    

