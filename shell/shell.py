import os

def ls(args):
    files = os.listdir('.')
    for file in files:
        print(file)

def pwd(args):
    print(os.getcwd())

def mkdir(args):
    os.mkdir(args[1])

def rmdir(args):
    os.rmdir(args[1])

def touch(args):
    with open(args[1], 'w') as f:
        pass

def rm(args):
    os.remove(args[1])

def help(args):
    print('Los siguientes comandos están disponibles:')
    print('ls        -- listar el contenido del directorio actual')
    print('pwd       -- mostrar el directorio actual')
    print('mkdir     -- crear un directorio')
    print('rmdir     -- eliminar un directorio')
    print('touch     -- crear un archivo vacío')
    print('rm        -- eliminar un archivo')
    print('help      -- mostrar esta ayuda')
    print('exit      -- salir del shell')

def main():
    while True:
        command = input('> ')
        args = command.split()
        if len(args) == 0:
            continue
        if args[0] == 'ls':
            ls(args)
        elif args[0] == 'pwd':
            pwd(args)
        elif args[0] == 'mkdir':
            mkdir(args)
        elif args[0] == 'rmdir':
            rmdir(args)
        elif args[0] == 'touch':
            touch(args)
        elif args[0] == 'rm':
            rm(args)
        elif args[0] == 'help':
            help(args)
        elif args[0] == 'exit':
            break
        else:
            print(f"Comando desconocido: {args[0]}")

if __name__ == '__main__':
    main()