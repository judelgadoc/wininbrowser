from flask import Flask, request
import os

app = Flask(__name__)

@app.route('/<command>', methods=['GET'])
def execute_command(command):
    args = command.split()
    if len(args) == 0:
        return 'Comando vacío'
    elif args[0] == 'ls':
        files = os.listdir('.')
        return '\n'.join(files)
    elif args[0] == 'pwd':
        return os.getcwd()
    elif args[0] == 'mkdir':
        if len(args) < 2:
            return 'Debe especificar el nombre del directorio'
        os.mkdir(args[1])
        return f'Se ha creado el directorio {args[1]}'
    elif args[0] == 'rmdir':
        if len(args) < 2:
            return 'Debe especificar el nombre del directorio'
        os.rmdir(args[1])
        return f'Se ha eliminado el directorio {args[1]}'
    elif args[0] == 'touch':
        if len(args) < 2:
            return 'Debe especificar el nombre del archivo'
        with open(args[1], 'w') as f:
            pass
        return f'Se ha creado el archivo {args[1]}'
    elif args[0] == 'rm':
        if len(args) < 2:
            return 'Debe especificar el nombre del archivo'
        os.remove(args[1])
        return f'Se ha eliminado el archivo {args[1]}'
    elif args[0] == 'help':
        return 'Los siguientes comandos están disponibles: \nls -- listar el contenido del directorio actual \npwd -- mostrar el directorio actual \nmkdir -- crear un directorio \nrmdir -- eliminar un directorio \ntouch -- crear un archivo vacío \nrm -- eliminar un archivo \nhelp -- mostrar esta ayuda'
    else:
        return f'Comando desconocido: {args[0]}'

if __name__ == '__main__':
    app.run(debug=True)
