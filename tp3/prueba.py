def main():
    with open("comunidad-chile.txt") as chile:
        contador = 0
        for linea in chile:
            linea=linea.rstrip("\n").split(",")
            for palabra in linea:
                contador += 1
        print(contador)
main()