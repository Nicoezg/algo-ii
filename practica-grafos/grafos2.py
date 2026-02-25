from collections import deque

#(★★★) Matías está en Barcelona y quiere recorrer un museo. Su idea es hacer un recorrido bastante lógico: 
#  empezar en una sala (al azar), luego ir a una adyacente a ésta, luego a una adyancente a la segunda 
#  (si no fue visitada aún), y así hasta recorrer todas las salas. Cuando no tiene más salas adyacentes para visitar 
#  (porque ya fueron todas visitadas), simplemente vuelve por donde vino buscando otras salas adyacentes. 
#  Teniendo un grafo no dirigido, que representa el mapa del museo (donde los vértices son salas, y las aristas 
#  (v, w) indican que las salas v y w se encuentran conectadas), implementar un algoritmo que nos devuelva una lista 
#  con un recorrido posible de la idea de Matías para visitar las salas del museo. Indicar el recorrido utilizado y 
#  el orden del algoritmo. Justificar.


def recorrer_mueso(grafo):
    padres = {}
    orden = {}
    visitados = set()
    origen = grafo.vertice_aleatorio()
    padres[origen] = None
    orden[origen] = 0
    visitados.add(origen)
    lista = []
    dfs(grafo, origen, visitados, padres, orden, lista)
    return lista



def dfs(grafo, origen, visitados, padres , orden, lista): ## DFS, O (V + E)
    for w in grafo.adyacentes(origen):
        
        if w not in visitados:
            lista.append(origen)
            padres[w] = origen
            orden[w] = orden[origen] + 1
            visitados.add(w)
            dfs(grafo, w, visitados, padres, orden)
    lista.append(origen)

#(★★★) Implementar un algoritmo que reciba un grafo dirigido, un vértice V y un número N, y 
# devuelva una lista con todos los vértices que se encuentren a exactamente N aristas de distancia 
# del vértice V. Indicar el tipo de recorrido utilizado y el orden del algoritmo. Justificar.

def distancia_n(grafo, v, n):
    lista = []
    visitados = set()
    orden = {}
    orden = bfs(grafo, visitados, orden, v, n) # O(V + E) en el peor de los casos

    for vertice, distancia in orden.items(): # O(K) siendo K los elementos que estan a distancia N del vertice V
        if distancia == n:
            lista.append(vertice)
    return lista
    



def bfs(grafo, visitados, orden, v, n): # O(V + E) el n que me pueden pasar puede ser muy grande si no, O(K + L) 
    # siendo K el numero de vertices hasta llegar a la distancia N y L el numero de aristas que veo en esos K vertices
    visitados.add(v)
    orden[v] = 0
    cola = deque()
    cola.append(v)

    while cola:
        v = cola.popleft()
        if orden[v] > n:
            continue

        for w in grafo.adyacentes(v):
            if w not in visitados:
                visitados.add(w)
                orden[w] = orden[v] + 1
                cola.append(w)
    
    return orden


# Implementar una función que reciba un grafo no dirigido y no pesado implementado con listas de adyacencia 
# (diccionario de diccionarios) y devuelva una matriz que sea equivalente a la representación de matriz de 
# adyacencia del mismo grafo. Indicar y justificar el orden del algoritmo implementado.

def matriz_adyacencia(grafo):
    matriz_adyacencia = []
    for v in grafo:
        fila_act = []
        for w in grafo:
            if grafo.estan_unidos(v, w):
                fila_act.append(1)
            else:
                fila_act.append(0)
        matriz_adyacencia.append(fila_act)
    return matriz_adyacencia

#(★★) Implementar una función que reciba un grafo no dirigido, y que compruebe la siguiente afirmación: 
# “La cantidad de vértices de grado IMPAR es PAR”. 
# Indicar y justificar el orden del algoritmo si el grafo está implementado como matriz de adyacencia.

def vertices_de_grado_par_impar(grafo): ## O(V + E)
    grados = obtener_grados(grafo) ## O(V + E)
    contador = 0
    for _, grado in grados.values(): ## O(V)
        if grado % 2 != 0:
            contador += 1
    return contador % 2 == 0
    

def obtener_grados(grafo): ## O(V)
    grados = {}
    for v in grafo:
        grados[v] = len(grafo.adyacentes(v))

#Dado un número inicial X se pueden realizar dos tipos de operaciones sobre el número:

#Multiplicar por 2
#Restarle 1.
#Implementar un algoritmo que encuentra la menor cantidad de operaciones a realizar para convertir el número 
# X en el número Y, con tan solo las operaciones mencionadas arriba (podemos aplicarlas la cantidad de veces que querramos).


def convertir_x_en_y(x, y):
# Si x < y, solo voy a tener que restar
# Si x > y, uso ambas operaciones
# Si x == y, no tengo que usar ninguna operacion
    if x == y:
        return 0
    grafo = Grafo(pesado = False, dirigido = False)
    grafo.agregar_vertice(x)
    grafo.agregar_vertice(y)


    while x != y:
        if x > y:
            grafo.agregar_vertice(x-1)
            grafo.agregar_arista(x, x-1)
            x = x-1
            continue
        if x < y:
            grafo.agregar_vertice(x*2)
            grafo.agregar_arista(x, x*2)
            x = x*2
    return len(grafo)-1

#Se tiene un arreglo de palabras de un lenguaje alienigena. Dicho arreglo se encuentra ordenado para dicho idioma 
# (no conocemos el orden de su abecedario). Implementar un algoritmo que reciba dicho arreglo y 
# determine un orden posible para las letras del abecedario en dicho idioma. Por ejemplo:
# {"caa", "acbd", "acba", "bac", "bad"} --> ['c', 'd', 'a', 'b']

def lenguaje_alienigena(arreglo):
    grafo = Grafo(dirigido= True)
    grados = {}
    for v in grafo:
        for w in grafo.adyecentes(w):
            grados[w] = grados.get(w, 0) + 1
    
    cola = deque()
    for v in grafo:
        if v not in grados:
            cola.append(v)
    
    result = []

    while cola:
        v = cola.popleft()
        result.append(v)
        for w in grafo.adyacentes(v):
            grados[w] = grados[w] + 1
            if grados [w] == 0:
                cola.append(w)
    
    return result





def grafo_desde_palabras(palabras):
    grafo = Grafo(esdirigido = True)
    for i in range(len(palabras) - 1):
        p1 = palabras[i]
        p2 = palabras[i+1]

        for letra in p1:
            if letra not in grafo:
                grafo.agregar_vertice(letra)
        
        for j in range (len(p1)-1):
            if p2[j] not in grafo:
                    grafo.agregar_vertice(p2[j])
            if p1[j] < p2[j]:
                grafo.agregar_arista(p1[j], p2[j], 1)
                continue
            if p1[j] > p2[j]:
                grafo.agrear_arista(p2[j], p1[j], 1)
                continue
    return grafo



# Implementar un algoritmo que reciba un grafo dirigido y nos devuelva la cantidad de componentes 
# débilmente conexas de este. Indicar y justificar la complejidad del algoritmo implementado.





#Contamos con un grafo dirigido que modela un ecosistema. En dicho grafo, cada vértice es una especie, y 
# cada arista (v, w) indica que v es depredador natural de w. Considerando la horrible tendencia del ser 
# humano por llevar a la extinción especies, algo que nos puede interesar es saber si existe alguna especie que, 
# si llegara a desaparecer, rompería todo el ecosistema: quienes la depredan no tienen un sustituto (y, por ende, 
# pueden desaparecer también) y/o quienes eran depredados por esta ya no tienen amenazas, por lo que crecerán 
# descontroladamente. Implementar un algoritmo que reciba un grafo de dichas características y devuelva una lista 
# de todas las especies que cumplan lo antes mencionado. Indicar y justificar la complejidad del algoritmo implementado.

def ecosistema(grafo):
    grados_s = grados_salida(grafo) # O (V + E)
    grados_e = grados_entrada(grafo)# O (V + E)
    res = []

    for presa, depredadores in grados_e.values(): # O (P) ; P < V
        if depredadores == 1 or depredadores == 0:
            res.append(presa)

    for depredador, presas in grados_s.values(): # O(D) ; D < V
        if presas == 1:
            res.append(depredador)

    return res



def grados_salida(grafo):
    gr = {}
    for v in grafo:
        for w in grafo.adyacentes(v):
            gr[v] = gr.get(v, 0) + 1
    return gr


def grados_entrada(grafo):
    gr = {}
    for v in grafo:
        for w in grafo.adyacentes(v):
            gr[w] = gr.get(w, 0) + 1
    return gr



        
        
