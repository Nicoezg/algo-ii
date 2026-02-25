import sys

OPERACIONES = ("camino", "conectados", "lectura", "diametro", "rango", "comunidad", "navegacion", "clustering")

def imprimir_operaciones():
    sys.stdout.write("\n".join(OPERACIONES))
    sys.stdout.write("\n")

def imprimir_camino(camino, distancia):
    sys.stdout.write((" -> ".join(camino)))
    sys.stdout.write(f"\nCosto: {distancia}\n")

def imprimir_numero(n):
    sys.stdout.write(str(n))
    sys.stdout.write("\n")

def imprimir_error_lectura():
    sys.stdout.write("No existe forma de leer las paginas en orden\n")

def imprimir_error_camino():
    sys.stdout.write("No se encontro recorrido\n")

def imprimir_orden(camino):
    sys.stdout.write(", ".join(camino))
    sys.stdout.write("\n")

def imprimir_navegacion(camino):
    sys.stdout.write((" -> ".join(camino)))
    sys.stdout.write("\n")

def imprimir_error_no_pertenece():
    sys.stdout.write("Pagina no pertenece a Wikipedia\n")