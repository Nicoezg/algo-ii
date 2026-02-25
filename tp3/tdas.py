class _Nodo:
    def __init__(self, dato=None, prox=None):
        self.dato = dato
        self.prox = prox

class Cola:
    def __init__(self):
        """Crea una cola vacía."""
        self.primero = None
        self.ultimo = None

    def encolar(self, x):
        """Encola el elemento x."""
        nuevo = _Nodo(x)
        if self.ultimo is not None:
                self.ultimo.prox = nuevo
                self.ultimo = nuevo
        else:
            self.primero = nuevo
            self.ultimo = nuevo

    def desencolar(self):
        """Desencola el primer elemento y devuelve su valor.
        Si la cola está vacía, levanta ValueError."""
        if self.primero is None:
            raise ValueError("La cola está vacía")
        valor = self.primero.dato
        self.primero = self.primero.prox
        if not self.primero:
            self.ultimo = None
        return valor
    def esta_vacia(self):
        """Devuelve True si la cola esta vacía, False si no."""
        return self.primero is None
    
class Pila:
    def __init__(self):
        '''
        Inicializa una nueva pila, vacía
        '''
        self.tope = None
        self.cantidad = 0

    def __len__(self):
        return self.cantidad

    def apilar(self, dato):
        '''
        Agrega un nuevo elemento a la pila
        '''
        nodo = _Nodo(dato, self.tope)
        self.tope = nodo
        self.cantidad += 1

    def desapilar(self):
        '''
        Desapila el elemento que está en el tope de la pila
        y lo devuelve.
        Pre: la pila NO está vacía.
        Pos: el nuevo tope es el que estaba abajo del tope anterior
        '''
        if self.esta_vacia():
            raise ValueError("pila vacía")
        dato = self.tope.dato
        self.tope = self.tope.prox
        self.cantidad -= 1
        return dato

    def ver_tope(self):
        '''
        Devuelve el elemento que está en el tope de la pila.
        Pre: la pila NO está vacía.
        '''
        if self.esta_vacia():
            raise ValueError("pila vacía")
        return self.tope.dato

    def esta_vacia(self):
        '''
        Devuelve True o False según si la pila está vacía o no
        '''
        return self.tope is None