from collections import deque

#1. Escribir una función ‘tieneCicloImpar que detecte si un grafo pasado por parametro tiene al menos
# un ciclo de longitud impar. Aclaración: NO se pide retornar el camino, solamente si existe o no.

def tieneCicloImpar(grafo):
    # Me tengo que fijar si es bipartito, un grafo bipartito NO tiene cilos de longitud impar.
    colores = {}
    origen = grafo.obtener_vertice()
    colores[origen] = 0
    cola = deque()
    cola.append(origen)

    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w in colores:
                if colores[w] == colores[v]:
                    return False
            else:
                colores[w] = 1 - colores[v]
                cola.append(w)

    return True

# 3. Implementar un algoritmo que encuentre todas las componentes debilmente conexas de un grafo
# dirigido. Indicar y justificar la complejidad del algoritmo asumiendo que el grafo esta implementado
# como matriz de adyacencias,.

def compDebiles(grafo):
    nuevo = Grafo(nodirigido = True)
    visitados = {}
    res = []
    for v in grafo: # O(V)
        nuevo.agregar_vertice(v)

    for v in grafo: # O(V ** 2)
        visitados.add(v)
        for w in grafo.adyacentes(v):
            if w not in visitados:
                nuevo.agregar_arista(v,w)
    visitados = {}
    for v in grafo: # O(V)
        if v not in visitados:
            res.append(bfs(visitados, grafo, v))
    return res


def bfs(visitados, grafo, v): # O(V + E)
    cola = deque()
    visitados.add(v)
    cola.append(v)
    comp = []
    comp.append(v)

    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                comp.append(w)
    return comp

def bfs(grafo):
    origen = grafo.obtener_vertice()
    cola = deque()
    visitados = set()
    padres = {}
    orden = {}
    cola.append(origen)
    visitados.add(origen)
    padres[origen] = None
    orden[origen] = 0

    while cola:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                visitados.add(w)
                cola.append(w)
    return padres, orden

def dfs(grafo):
    origen = grafo.obtener_vertice()
    visitados = set()
    padres = {}
    orden = {}
    visitados.add(origen)
    orden[origen] = 0
    padres[origen] = None
    return _dfs(grafo, origen, visitados, padres, orden)

def _dfs(grafo, origen, visitados, padres, orden):
    for w in grafo.adyacentes(origen):
        if w not in visitados:
            padres[w] = origen
            orden[w] = origen + 1
            _dfs(grafo, w, visitados, padres, orden)

def bfs_topologico(grafo):
    cola = deque()
    gr_ent = {}
    lista = []

    for v in grafo:
        gr_ent[v] = 0

    for v in grafo:
        for w in grafo.adyacentes(v):
            gr_ent[w] += 1

    for v in grafo:
        if gr_ent[v] == 0:
            cola.append(v)
    
    while cola:
        v = cola.popleft()
        lista.append(v)
        for w in grafo.adyacentes(v):
            gr_ent[w] -= 1
            if gr_ent[w] == 0:
                cola.append(w)
    return lista

def dijkstra(grafo):
    distancia = {}
    origen = grafo.obtener_aleatorio()
    heap = Heap(b-a)

    for v in grafo:
        distancia[v] = float("inf")
    
    distancia[origen] = 0
    heap.encolar(0,origen)

    while heap:
        _, v = heap.desencolar()
        for w in grafo.adyacentes(v):
            dist_posible = distancia[v] + grafo.peso_arista(v,w)
            if dist_posible < distancia[w]:
                distancia[w] = dist_posible
                heap.encolar(distancia[w], w)
    return distancia


def bellmanford(grafo):
    distancia = {}
    visitados = set()
    origen = grafo.obtener_aleatorio()
    aristas = []
    
    for v in grafo:
        distancia[v] = float("inf")
    
    for v in grafo:
        for w in grafo:
            aristas.append(grafo.obtener_arista(v,w))

    for i in range (len(grafo)):
        cambio = False
        for v,w, peso in aristas:
            if peso + distancia[v] < distancia[w]:
                distancia[w] = peso + distancia[v]
                cambio = True
        if not cambio:
            return aristas
    
    for v,w, peso in aristas:
        if peso + distancia[v] < distancia[w]:
            raise ValueError("Hay un ciclo negativo!")
    return aristas

def prim(grafo):
    heap = Heap(b-a)
    nuevo = Grafo()
    visitados = set()
    origen = grafo.vertice_aleatorio()
    visitados.add(origen)
    for v in grafo:
        nuevo.agregar_vertice(v)
    for w in grafo.adyacentes(origen):
        heap.Encolar((origen,w), grafo.peso_arista(origen,w))
    
    while heap:
        origen, destino, peso = heap.desencolar()
        if destino not in visitados:
            nuevo.agregar_arista((origen, destino), peso)
            visitados.add(destino)
            for x in grafo.adyacentes(destino):
                if x not in visitados:
                    heap.Encolar((destino,x), grafo.peso_arista(origen,w))
    return nuevo

def kruskal(grafo):
    conjuntos = UnionFind()
    aristas = []
    nuevo = Grafo()
    for v in grafo:
        nuevo.agregar_vertice(v)
        for w in grafo.adyacentes(v):
            aristas.append((v,w), grafo.peso_arista(v,w))
    
    aristas = HeapSort(aristas)
    
    for arista in aristas:
        v,w,peso = arista
        if conjuntos.find(v) != conjuntos.find(w):
            nuevo.agregar_arista((v,w), peso)
            conjuntos.union(v,w)

    return nuevo

    

    

            





    



    


    


    

