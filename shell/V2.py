from flask import Flask, request
import shutil
import os

app = Flask(__name__)

history = []

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
        dir_name = args[1]
        if os.path.exists(dir_name):
            return f'El directorio "{dir_name}" ya existe'
        try:
            os.mkdir(dir_name)
        except Exception as e:
            return f'Error al crear el directorio "{dir_name}": {e}'
        return f'Se ha creado el directorio "{dir_name}"'
    
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
       return "Los siguientes comandos están disponibles: <br>ls -- listar el contenido del directorio actual <br>pwd -- mostrar el directorio actual <br>mkdir  -- crear un directorio <br>rmdir  -- eliminar un directorio <br>touch  -- crear un archivo vacío <br>rm  -- eliminar un archivo <br>help  -- mostrar esta ayuda <br>cd  -- cambiar el directorio actual <br>mv -- mover o renombrar archivos <br>cp -- copiar archivos <br>history -- mostrar historial de comandos"

    elif args[0] == 'cd':
        if len(args) < 2:
            return 'Debe especificar el directorio al que desea cambiar'

        try:
            print(args[1])
            os.chdir(args[1]) 
        except FileNotFoundError:
            return f'Directorio "{args[1]}" no encontrado'
        return f'Directorio cambiado a "{args[1]}"'
    
    elif args[0] == 'mv':
        if len(args) < 3:
            return 'Debe especificar el archivo/directorio y la ubicación de destino'
        try:
            os.rename(args[1], args[2])
            return f'"{args[1]}" movido/renombrado a "{args[2]}"'
        except FileNotFoundError:
            return f'No se pudo mover o renombrar "{args[1]}"'
        except FileExistsError:
            return f'"{args[2]}" ya existe'
        except OSError as e:
            return f'Error al mover/renombrar "{args[1]}": {e}'

    
    elif args[0] == 'cp':
        if len(args) < 3:
            return 'Debe especificar el archivo/directorio y la ubicación de destino'
        try:
            if os.path.isdir(args[1]):
                shutil.copytree(args[1], args[2])
            else:
                shutil.copy2(args[1], args[2])
        except FileNotFoundError:
            return f'No se pudo encontrar el archivo/directorio "{args[1]}"'
        except shutil.Error as e:
            return f'Error al copiar "{args[1]}": {e}'
        return f'"{args[1]}" copiado a "{args[2]}"'
    
    elif args[0] == 'history':
        return "\n".join(history)
    
    else:
        return f'Comando desconocido: {args[0]}'

if __name__ == '__main__':
    app.run(debug=True)
