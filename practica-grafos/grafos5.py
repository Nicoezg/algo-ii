from collections import deque

#1. Como todos sabemos, Taller de Programación I implica programar mucho. El equipo de Federico esta desarrollando un
#Age of Empires y le asignaron el modulo de movimientos. Dispone de un mapa, podríamos decir una grilla, donde se
#puede ver para cada celda su contenido (por ejemplo si está libre, si es agua, un árbol, etc). Fede tiene que implementar
#un algoritmo que, a partir de una unidad (por ejemplo un soldado), en una celda en especifico encuentre el camino a
#una celda objetivo. Considerar que entre celdas puede haber diferencia de alturas por lo que puede costarle mas (o
#menos) a una unidad ir de una celda a otra. Por supuesto que no puede ser cualquier camino, si no el que haga que la
#unidad llegue más rápido a su objetivo.
#a. Modelar el problema usando Grafos especificando de forma completa todos sus componentes.
#b. Implementar un algoritmo que a partir de una celda de origen y una de destino, retorne el camino que tiene que
#hacer la unidad, indicando y justificando la complejidad final.

def age_of_empires(grafo, origen, destino):
    distancia = {}
    padre = {}
    for v in grafo: # O(V)
        distancia[v] = float("inf")

    h = Heap(b-a)
    h.encolar((0,origen))
    distancia[origen] = 0
    visitados.add(origen)
    padre[origen] = None
    
    while not h.EstaVacia(): # Encolamos y desencolamos E aristas con un heap => O (V + E log V) V es casi = E
        _, v = h.desencolar()
        if v == destino:
            return reconstruir_camino(padre, destino)
        
        for w in grafo.adyacentes(v):
            if w != obstaculo:
                posible_dist = distancia[v] + grafo.peso_arista(v,w)
                if posible_dist < distancia[w]:
                    padre[w] = v
                    distancia[w] = posible_dist
                    h.encolar(distancia[w], w)
    return reconstruir_camino(padre, destino)

# IMPORTANTE: Ejecutamos solo una vez ese O(C), por ende, no suma mucho a la complejidad.


def reconstruir_camino(padre, destino): # O (C), siendo C el numero de vertices presentes en el camino para llegar al destino.
    res = []
    res.append(destino)
    v = destino
    while padre[v]:
        res.append(padre[v])
        v = padre[v]
    return res


# 2. Grace, docente de un primer grado de un colegio primario, tiene como tarea armar la fila el acto del 9 de julio. Como
# no quiere que los chicos se distraigan, quiere que cada uno quede entre dos personas de las cuales no es amigo. Para eso
# dispone de un mapa de amistades donde se puede ver quién es amigo de quién.
# a. Modelar el problema usando Grafos especificando de forma completa todos sus componentes.
# b. Implementar un algoritmo que reciba como parametro el grafo del punto anterior y retorne una posible fila en
# formato de lista. Retornar lista vacia si no existe solución. Indicar y justificar la complejidad final.
# Pista: No pensar el ejercicio armando, inicialmente, la fila sino viendo como separar a cada uno de sus amigos.

def acto(grafo):
    res = []
    no_amigos1 = set()
    no_amigos2 = set()
    amistades = {}
    cola = deque()
    origen = grafo.vertice_aleatorio()
    cola.append(origen)
    amistades[origen] = 0

    while cola: # Me fijo si el grafo es bipartito
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w in amistades:
                if amistades[v] == amistades[w]:
                    raise ValueError("No hay solucion")
                else:
                    amistades[w] = 1 - amistades[v]
                    if amistades[w] == 1: # Armo mis dos conjuntos de "no amigos"
                        no_amigos1.add(w)
                    else:
                        no_amigos2.add(w)

                    cola.append(w)

    encontrado = False
    for persona in no_amigos1: # Busco el borde de no amigos 1
        if grafo.adyacentes(persona) != no_amigos2:

            for no_amigo in no_amigos2:

                if no_amigo not in grafo.adyacentes(persona):
                    borde2 = no_amigo
                    borde1 = persona
                    encontrado = True
                    break

            if encontrado == True:
                break
    
    if encontrado == False:
        raise ValueError("No se encontro un borde")
    

    for persona in no_amigos1: # Meto todas las personas que no se llevan de no_amigos1
        if persona == borde1:
            continue
        res.append(persona)

    res.append(borde1) # Meto el borde 1
    res.append(borde2) # Meto el borde 2

    for persona in no_amigos2: # Meto todas las personas que no se llevan de no_amigos2
        if persona == borde2:
            continue
        res.append(persona)

    return res

    


#3. Implementar un algoritmo que a partir de un arbol de tendido mínimo retorne verdadero o falso si al agregar una arista
# especifica (origen, destino peso) al grafo original el arbol de tendido minimo cambiaría o no. Indicar y justificar la
# complejidad asumiendo que el grafo esta implementado como matriz de adyacencias. La firma de la función debe ser
# sigue_siguendo_mst(mst, origen, destino, peso) donde mst es un grafo que es un arbol de tendido mínimo de
# un grafo mas completo. Se puede asumir que la arista nueva que se propone agregar no formaba parte del grafo original.

def sigue_siendo_mst(mst, origen, destino, peso):
    padres = {}
    visitados = set()
    nuevo = Grafo(no_dirigido, pesado)
    c = deque()
    visitados.add(origen)
    c.append(origen)
    padres[origen] = None

    while c:
        v = c.popleft()
        if v == destino:
            break
        for w in mst.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                visitados.add(w)
    
    aristas = reconstruir_aristas(mst, origen, destino, padres)
   
    for arista in aristas:
        v, w, peso_arista = arista
        if peso_arista > peso:
            return False
    return True


def reconstruir_aristas(mst, origen, destino, padres):
    res = []
    w = destino
    peso_total = 0
    while w != origen:
        v = padres[w]
        peso = mst.peso_arista(v, w)
        res.append((v, w), peso )
        peso_total += peso
        w = v
    return res, peso_total

def dfs(grafo):
    visitados = set()
    padres = {}
    orden = {}
    for v in grafo:
        if v not in visitados:
            _dfs(grafo, visitados, padres, orden, v)

def _dfs(grafo, visitados, padres, orden, v):
    for w in grafo.adyacentes(v):
        if w not in visitados:
            padres[w] = v
            orden[w] = orden[v] + 1
            visitados.add(w)
            _dfs(grafo, visitados, padres, orden, w)



    

            
            









    

    


    

    

        








