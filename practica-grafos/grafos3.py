from collections import deque

#Implementar una función que reciba un grafo no dirigido y determine si el mismo no tiene ciclos de una 
# cantidad impar de vértices. Indicar y justificar la complejidad de la función.

def tiene_ciclos_impares(grafo): 
  return es_bipartito(grafo)

# Un grafo bipartito NO puede tener ciclos impares.

def grados_entrada(grafo): ## O (V + E)
  g_ent= {}
  for v in grafo:
    g_ent[v] = 0
  for v in grafo:
    for w in grafo.adyacentes(v):
      g_ent[w] += 1
  return g_ent

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


def topologico_dfs(grafo):
  visitados = set()
  pila = Pila()
  for v in grafo:
    if v not in visitados:
      visitados.add(v)
      _dfs(grafo, v, visitados, pila)
  return pila_a_lista(pila)

def _dfs(grafo, v, visitados, pila):
  for w in grafo.adyacentes(v):
    if w not in visitados:
      visitados.add(w)
      _dfs(grafo, w, visitados, pila)
  pila.apilar(v)


def pila_a_lista(pila):
  lista = []
  while not pila.esta_vacia():
    lista.append(pila.desapilar())
  return lista


def dijkstra(grafo, origen, destino): # O(V + E log V)
  dist = {}
  padre = {}
  for v in grafo:
    dist[v] = float("inf")

  dist[origen] = 0
  padre[origen] = None
  q = Heap()
  q.encolar((0, origen))

  while not q.esta_vacio():
    _, v = q.desencolar()

    if v == destino:
      return padre, dist
    
    for w in grafo.adyacentes(v):
      posible_mejor_distancia = dist[v] + grafo.peso(v, w)

      if posible_mejor_distancia < dist[w]:
        dist[w] = posible_mejor_distancia
        padre[w] = v
        q.encolar((dist[w], w))

  return padre, dist


def bellman_ford(grafo, origen): # O (V x E)
  distancia = {}
  padre = {}
  for v in grafo:
    distancia[v] = float("inf")

  distancia[origen] = 0
  padre[origen] = None
  aristas = obtener_aristas(grafo)

  for i in range(len(grafo)): # Por cada vertice
    cambio = False
    for origen, destino, peso in aristas: # Veo todas las aristas

      if distancia[origen] + peso < distancia[destino]:
        cambio = True
        padre[destino] = origen
        distancia[destino] = distancia[origen] + peso

    if not cambio:
      return padre, distancia
    
  for v, w, peso in aristas:
    if distancia[v] + peso < distancia[w]:
      return None #Hay un ciclo negativo
  return padre, distancia


def prim(grafo): # O (V + E log V)
  v = grafo.obtener_vertice_aleatorio()
  visitados = set()
  visitados.add(v)
  q = Heap()

  for w in grafo.adyecentes(v): # Agrego todos las aristas al heap
    q.encolar((v,w), grafo.peso(v,w))
  arbol = Grafo(es_dirigido = False, lista_vertices= grafo.obtener_vertices() )

  while q: 
    (v, w), peso = q.desencolar()
    if w in visitados: # Si ya visite al punto destino, estaría completando un ciclo
      continue
    arbol.agregar_arista(v, w, peso) # Si no visite al punto destino, agrego la arista
    visitados.add(w)

    for x in grafo.adyacentes(w): # Encolo todas las aristas del vertice adyacente cuyos puntos destino no hayan sido visitados
      if x not in visitados:
        q.encolar((w,x), grafo.peso(w,x))

  return arbol    

def kruskal(grafo): # Depende del algoritmo de ordenamiento, en este caso, # O (E log V) (use un ordenamiento comparativo)
  conjuntos = UnionFind(grafo.obtener_vertices)
  aristas = sorted(obtener_aristas(grafo), key = lambda arista: arista[PESO]) # Ordeno las aristas por peso (menor a mayor)
  arbol = Grafo(False, grafo.obtener_vertices())

  for a in aristas: # O(alpha-1)
    v, w, peso = a
    if conjuntos.find(v) == conjuntos.find(w): # Si forman parte del mismo conjunto no hago nada
      continue
    arbol.arista(v,w, peso) #Si no forman parte del mismo conjunto, agrego la arista
    conjuntos.union(v, w) # Uno los vertices en el mismo conjunto.

  return arbol  #nos da el bosque de tendido mínimo, a diferencia de prim.

def dfs_puntos_articulacion(grafo, v, visitados, padre, orden, mas_bajo, ptos, es_raiz= False):
  hijos = 0
  mas_bajo[v] = orden[v]
  for w in grafo.adyacentes(v):
    if w not in visitados:
      hijos += 1
      orden[w] = orden[v] + 1
      padre[w] = v
      visitados.add(w)
      dfs_puntos_articulacion(grafo, w, visitados, padre, orden, mas_bajo, ptos, es_raiz)
    
      if mas_bajo[w] >= orden[v] and not es_raiz:
        # No hubo forma de pasar por arriba a este vértice, es punto de articulacion
        # se podria agregar como condicion "and v not in ptos" (ya que podria darse por mas de una rama)

        ptos.add(v)
      
      mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
    elif padre[v] != w: #evitamos considerar a la arista con el padre como una de retorno
    #Si uno ya es visitado, significa que puedo subir (si es que no podia ya ir mas arriba)
      mas_bajo[v] = min(mas_bajo[v], orden[w])
    
  if es_raiz and hijos > 1:
    ptos.add(v)


def ptos_articulacion(grafo): # O (V + E), un dfs con el que hago operaciones O(1) en el medio.
  origen = grafo.random()
  origen = "A"
  puntos_articulacion = set()
  dfs_puntos_articulacion(grafo, origen, {origen}, {origen: None}, {origen: 0}, {}, puntos_articulacion)


def cfcs_grafo(grafo): ## O(V + E), un dfs al que le hago operaciones O(1) en el medio
  resultados = []
  visitados = set()
  for v in grafo:
    if v not in visitados:
      dfs_cfc(grafo, v, visitados, {}, {}, Pila(), set(), resultados, [0])
  return resultados


def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global): 
  orden[v] = mas_bajo[v] = contador_global[0]
  contador_global[0] += 1
  visitados.add(v)
  pila.apilar(v)
  apilados.add(v)
  for w in grafo.adyacentes(v):
      if w not in visitados:
        # llamamos recursivamente
        dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, contador_global)
      if w in apilados:
        # Nos tenemos que fijar que esté entre los apilados
        # en otro dfs hecho antes --> no son parte de la misma componente
        mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
  if orden[v] == mas_bajo[w]:
    # Se cumple la condicion de cierre de CFC.
    nueva_cfc = []
    while True:
      w = pila.desapilar()
      apilados.remove(w)
      nueva_cfc.append(w)
      if w == v:
        break
    cfcs.append(nueva_cfc)






