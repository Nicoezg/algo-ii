# Algoritmos y Programación II — Buchwald (2023, 2° Cuatrimestre)

Repositorio académico correspondiente a la materia **Algoritmos y Programación II** de la **Facultad de Ingeniería, UBA**, cursada durante el segundo cuatrimestre de 2023, con el profesor **Buchwald**.

---

## Estructura del Repositorio

### 📦 `tdas/` — Tipos de Datos Abstractos

Implementaciones en **Go** de los principales TDAs utilizados durante la materia:

| TDA | Descripción |
|-----|-------------|
| `pila/` | Pila dinámica (stack) |
| `cola/` | Cola enlazada (queue) |
| `lista/` | Lista enlazada (linked list) |
| `cola_prioridad/` | Heap / Cola de prioridad |
| `diccionario/` | Diccionario con tabla de hash cerrado y ABB (Árbol Binario de Búsqueda) |

Cada TDA incluye su interfaz, implementación y tests unitarios.

---

### 🗳️ `tp1/` — ReRePolez

**Lenguaje:** Go

Trabajo Práctico 1: simulación de un sistema electoral. El programa procesa un padrón de votantes y un archivo de partidos, recibe comandos por entrada estándar (`ingresar`, `votar`, `fin-votar`, `deshacer`) y muestra los resultados de las elecciones.

**Estructuras utilizadas:** Cola enlazada (para la fila de votantes), Pila (para deshacer votos), búsqueda binaria sobre el padrón ordenado.

---

### 📸 `tp2/` — AlgoGram

**Lenguaje:** Go

Trabajo Práctico 2: simulación de una red social estilo Instagram. Permite a los usuarios hacer login/logout, publicar posts, ver su feed personalizado, likear publicaciones y mostrar los likes de cada post.

**Estructuras utilizadas:** Tabla de hash (diccionario) para usuarios y publicaciones, Cola de prioridad para el feed.

---

### 🌐 `tp3/` — NetStats

**Lenguaje:** Python

Trabajo Práctico 3: análisis de redes usando grafos sobre datos de Wikipedia. Implementa comandos para:

- `camino`: encontrar el camino más corto entre dos artículos
- `conectados`: hallar la componente fuertemente conexa de un artículo
- `lectura`: obtener un orden de lectura válido (sort topológico)
- `diametro`: calcular el diámetro del grafo
- `rango`: contar artículos a exactamente *n* saltos de distancia
- `comunidad`: detectar comunidades (Label Propagation)
- `navegacion`: simular navegación por PageRank
- `clustering`: calcular el coeficiente de clustering

---

### 🧩 `practica-grafos/` — Práctica de Grafos

**Lenguaje:** Python

Ejercicios de práctica sobre algoritmos de grafos:

- Detección de ciclos (DFS)
- Verificación de árbol
- Teoría de los 6 grados de separación (BFS)
- Sort topológico
- Componentes fuertemente conexas (Kosaraju)
- Camino mínimo (Dijkstra, BFS)
- Árbol de tendido mínimo (Prim, Kruskal)

---

### 📝 `practica-parcial2/` — Práctica Parcial 2

**Lenguaje:** Go

Ejercicios de práctica para el segundo parcial de la materia. Incluye resolución de problemas algorítmicos usando los TDAs implementados (heaps, ABBs, tablas de hash, listas, pilas y colas).

---

## Tecnologías Utilizadas

- **Go** — TDAs, TP1, TP2 y práctica de parciales
- **Python 3** — TP3 y práctica de grafos

## Autor

**Nicolás Ezequiel García** — Estudiante FIUBA, 2023
