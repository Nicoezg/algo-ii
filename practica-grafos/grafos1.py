from collections import deque

#Implementar un algoritmo que, dado un grafo dirigido, 
#nos devuelva un ciclo dentro del mismo, si es que lo tiene. Indicar el orden del algoritmo.
def encontrar_ciclo(grafo):
  
  visitados = {}
  padre = {}
  for v in grafo: # O(V)

    if v not in visitados:
      padre[v] = None
      ciclo = dfs_ciclo(grafo, v, visitados, padre)

      if ciclo is not None:
        return ciclo
      
  return None


def dfs_ciclo(grafo, v, visitados, padre): # O (V + E)
  visitados[v] = True
  for w in grafo.adyacentes(v): # O (E)
    if w in visitados:
      
      # Si w fue visitado y es padre de v, entonces es la arista de donde
      # vengo (no es ciclo).
      # Si no es su padre, esta arista (v, w) cierra un ciclo que empieza
      # en w.

      if w != padre[v]:
        return reconstruir_ciclo(padre, w, v)
      
    else:
      padre[w] = v
      ciclo = dfs_ciclo(grafo, w, visitados, padre)

      if ciclo is not None:
        return ciclo

  # Si llegamos hasta acá es porque no encontramos ningún ciclo.
  return None

def reconstruir_ciclo(padre, w, v): # O(C)

  camino = []
  while w != v:
    camino.append(v)
    v = padre[v]
  camino.append(w)
  return camino[::-1]

#6(★★) Un árbol es un grafo no dirigido que cumple con las siguientes propiedades:

def es_arbol(grafo):
  return es_conexo(grafo) and (len(grafo) - 1 == contar_aristas(grafo))

def contar_aristas(grafo): #O(V + 2*E)
   contador = 0
   for v in grafo:
      for w in grafo.adyacentes(v):
         contador += 1
   return contador // 2

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


#7.

def grafo_traspuesto(grafo):
   nuevo = Grafo(dirigido = True)
   for v in grafo: ## O (V)
      nuevo.agregar_vertice(v)
   for v in grafo: ## O(V + E)
      for w in grafo:
         nuevo.agregar_arista(w,v)

## En caso de matriz de adyacencia es O(V**2) ---> E = V. O(V * E) == O (V ** 2). 
## En caso de mat


#8 (★★★) La teoría de los 6 grados de separación dice que cualquiera en la Tierra puede estar conectado a cualquier 
# otra persona del planeta a través de una cadena de conocidos que no tiene más de cinco intermediarios 
# conectando a ambas personas con solo seis enlaces). Suponiendo que se tiene un grafo G en el que cada vértice 
# es una persona y cada arista conecta gente que se conoce (el grafo es no dirigido):

#a. Implementar un algoritmo para comprobar si se cumple tal teoría para todo el conjunto de personas representadas 
# en el grafo G. Indicar el orden del algoritmo.

#b. Suponiendo que en el grafo G no habrán altas ni bajas de vértices, pero podrían haberla de aristas (la gente se 
# va conociendo, explicar las ventajas y desventajas que tendría implementar al grafo G con una matriz de adyacencia.

def bfs(grafo, visitados, orden,origen): ##O (V + E)
   q = deque()
   q.append(origen)
   visitados.add(origen)
   orden[origen] = 0

   while q:
      v = q.popleft()

      for w in grafo.adyacentes(v):
         if w not in visitados:
            q.append(w)
            visitados.add(w)

            if orden[v] + 1 > 6:
               return False
            
            orden[w] = orden[v] + 1
   return True, visitados



def separacion(grafo):
   if not es_conexo(grafo): # O(V + E)
      return False
   visitados = set()
   orden = {}

   for v in grafo: #  O(V)
      
      if v not in visitados:
        condicion, visitados = bfs(grafo, visitados, orden, v) #O(V + E)

        if condicion == False:
            return False
        
   return True

## O(Vx(V + E))

## Matriz de adyacencia : O(V ** 2), sacar aristas: O(1), ver si A-->B: O(1)