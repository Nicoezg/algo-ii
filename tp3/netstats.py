#!/usr/bin/python3
import sys
import comandos

ARGUMENTOS_ESPERADOS = 2
LIMITE_RECURSION = 75000
ERROR_ARGUMENTOS = "No hay suficientes argumentos"

def main():
    sys.setrecursionlimit(LIMITE_RECURSION)
    argumentos = sys.argv
    if len(argumentos) != ARGUMENTOS_ESPERADOS:
        raise ValueError(ERROR_ARGUMENTOS)
    grafo = comandos.inicializar_grafo(argumentos[1])
    conectados = {}
    cfcs = []
    diametro = []
    "Variables usadas en forma de cache"

    for entrada in sys.stdin:
        entrada = entrada.rstrip().split(" ")

        if len(entrada) == 0:
            continue

        comando = entrada[0]
        parametros = (" ").join(entrada[1:])
        parametros = parametros.split(",")

        if comando == "listar_operaciones":
            comandos.listar_operaciones()
        
        if comando == "camino":
            comandos.camino(grafo, parametros[0], parametros[1])
        
        if comando == "conectados":
            cfcs = comandos.conectividad(grafo, parametros[0], conectados, cfcs)
        
        if comando == "lectura":
            comandos.lectura(grafo, parametros)
        
        if comando == "diametro":
            comandos.diametro(grafo, diametro)
        
        if comando == "rango":
            comandos.rango(grafo, parametros[0], parametros[1])
        
        if comando == "comunidad":
            comandos.comunidad(grafo, parametros[0])

        if comando == "navegacion":
            comandos.navegacion(grafo, parametros[0])
        
        if comando == "clustering":
            if len(entrada) > 1:
                comandos.clustering(grafo,parametros[0])
            else:
                comandos.clustering(grafo)


if __name__ == "__main__":
    main()
        