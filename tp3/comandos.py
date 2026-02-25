import biblioteca
import impresion
import grafo
import conversion_resultado

ERROR_NO_PERTENECE = "Pagina no pertenece a wikipedia"
CANTIDAD_DECIMALES_CLUSTETRING = 3

def listar_operaciones():
    impresion.imprimir_operaciones()

def camino(grafo, origen, destino):
    padres, orden = biblioteca.distancia_a_vertices(grafo, origen, destino)
    if destino not in orden:
        impresion.imprimir_error_camino()
        return
    
    camino = conversion_resultado.reconstruir_camino(padres, destino)
    impresion.imprimir_camino(camino, orden[destino])

def conectividad(grafo, pagina, dic_cfcs, cfc):
    if not grafo.verificar_vertice(pagina):
        impresion.imprimir_error_no_pertenece()
        return
    if len(cfc) == 0:
        cfc = biblioteca.conectividad(grafo)
        for i in range(len(cfc)):
            for vertice in cfc[i]:
                dic_cfcs[vertice] = i
                
    impresion.imprimir_orden(cfc[dic_cfcs[pagina]])
    return cfc

def lectura(grafo, paginas_a_visitar):
    resultado = biblioteca.lectura(grafo, paginas_a_visitar)
    if len(resultado) != len(paginas_a_visitar):
        impresion.imprimir_error_lectura()
        return
    
    impresion.imprimir_orden(resultado)

def diametro(grafo, diametro):
    if len(diametro) == 0:
        camino, max = biblioteca.diametro(grafo)
        diametro.append(camino)
        diametro.append(max)
    else:
        camino, max = diametro
    

    impresion.imprimir_camino(camino, max)

def rango(grafo, pagina, n):
    if not verificar_pagina_perteneciente(grafo, pagina):
        return
    _, orden = biblioteca.distancia_a_vertices(grafo, pagina)
    res = conversion_resultado.vertices_a_distancia_n(int(n), orden)
    impresion.imprimir_numero(res)


def comunidad(grafo, pagina):
    if not verificar_pagina_perteneciente(grafo, pagina):
        return
    resultado = biblioteca.comunidad(grafo, pagina)
    impresion.imprimir_orden(resultado)

def navegacion(grafo, origen):
    if not verificar_pagina_perteneciente(grafo, origen):
        return
    resultado = biblioteca.navegacion(grafo, origen)
    impresion.imprimir_navegacion(resultado)

def clustering(grafo, pagina = None):
    if pagina is not None:
       coeficiente = round(biblioteca.coeficiente_clustering(grafo, pagina),CANTIDAD_DECIMALES_CLUSTETRING)
       impresion.imprimir_numero(coeficiente)
       return
    suma = 0.000

    for v in grafo:
       suma += biblioteca.coeficiente_clustering(grafo,v)
    promedio = round(suma / len(grafo.obtener_vertices()), CANTIDAD_DECIMALES_CLUSTETRING)
    impresion.imprimir_numero(promedio)

def inicializar_grafo(ruta):
    res = grafo.Grafo(dirigido=True)
    with open(ruta) as wikipedia:
        for linea in wikipedia:
            linea=linea.rstrip("\n").split("\t")

            for i in range(len(linea)):
                if not res.verificar_vertice(linea[i]):
                    res.agregar_vertice(linea[i])

                if i == 0:
                    vertice=linea[i]

                else:
                    res.agregar_arista(vertice,linea[i])
    
    return res

def verificar_pagina_perteneciente(grafo, pagina):
    if not grafo.verificar_vertice(pagina):
       impresion.imprimir_error_no_pertenece()
       return False
    
    return True
