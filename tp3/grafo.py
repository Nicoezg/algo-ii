import random

ERROR_VERTICE_INEXISTENTE = "El vertice no pertenece al grafo"
ERROR_ARISTA_INEXISTENTE = "La arista no pertenece al grafo"
ERROR_VERTICES_INEXISTENTES = "Algun vertice no pertenece al grafo"
ERROR_GRAFO_VACIO = "El grafo está vacio"

class Grafo:

    def __init__(self, dirigido = False, pesado = False, vertices_iniciales = []):
        self.datos = dict()
        self.vertices = []
        for vertice in vertices_iniciales:
            self.datos[vertice] = {}
            self.vertices.add(vertice)
        self.dirigido = dirigido
        self.pesado = pesado
    
    def __len__(self):
        return len(self.vertices)
    
    def __iter__(self):
        return iter(self.vertices)

    def agregar_vertice(self, v):
        self.datos[v] = {}
        self.vertices.append(v)
   
    def borrar_vertice(self, v):
        if not self.verificar_vertice(v):
            raise ValueError(ERROR_VERTICE_INEXISTENTE)

        
        for aristas in self.datos.values():
            if v in aristas:
                aristas.pop(v)
        
        self.datos.pop(v)
        self.vertices.pop(v)

    def agregar_arista(self, v, w,peso=1):
        if not self.verificar_arista(v, w):
            raise ValueError(ERROR_ARISTA_INEXISTENTE)
        
        self.datos[v][w] = peso

        if not self.dirigido:
            self.datos[w][v]=peso

    def borrar_arista(self, v, w):
        if not self.verificar_arista(v, w):
            raise ValueError(ERROR_ARISTA_INEXISTENTE)
        
        self.datos[v].pop(w)

        if not self.dirigido:
            self.datos[w].pop(v)

    def peso_arista(self, v, w):
        if not self.verificar_arista(v, w):
            raise ValueError(ERROR_ARISTA_INEXISTENTE)
        
        return self.datos[v][w]

    def estan_unidos(self, v, w):
        if not self.verificar_arista(v, w):
            raise ValueError(ERROR_VERTICES_INEXISTENTES)
        
        return w in self.datos[w]
    
    def obtener_vertices(self):
        return self.vertices

    def vertice_aleatorio(self):
        if len(self) == 0:
            raise ValueError(ERROR_GRAFO_VACIO)
        return random.choice(self.vertices)
    
    def adyacentes(self, v):
        lista = []
        for w in self.datos[v].keys():
            lista.append(w)
        return lista
    
    def verificar_vertice(self, v):
        return v in self.datos

    def verificar_arista(self, v, w):
        return v in self.datos and w in self.datos