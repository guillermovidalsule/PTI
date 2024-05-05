from flask import Flask, render_template, request, redirect, url_for, send_file
import requests

app = Flask(__name__)

API_URL = "http://nattech.fib.upc.edu:40342"

@app.route('/')
def index():
     return redirect(url_for('robot'))



@app.route('/robot', methods=['GET', 'POST'])
def robot():
    if request.method == 'POST':
        # Obtener los datos del formulario
        username = request.form['username']
        password = request.form['password']

        # Crear el JSON
        data = {
            "username": username,
            "password": password
        }
       
        headers1 = {'Content-Type': 'application/json'}
        print (data)
    
        print("Hola desde Flask!")
        #me peta aqui, ya que no establece conexión
        response = requests.post('http://nattech.fib.upc.edu:40342/users/existe', json=data, headers=headers1)

        
        print("Adios desde Flask")

        if response.status_code == 200:
            #obtener token de la response y procesar la petición de creación del robot
            # Agregar el token a la cabecera de la solicitud
            request.headers['Authorization'] = f'Bearer {token}'
            headers1 = {
                 "Authorization": f"Bearer {token}",
                 "Content-Type": "application/json"  # Especificas el tipo de contenido si estás enviando datos JSON
            }

            robotname = request.form['robotname']
            data = {
            "robotname": robotname
             }
            
            response = request.post('http://nattech.fib.upc.edu:40342)/robot/alta', json=data, headers=headers1) 


            nombre_archivo = "info_robot.txt"
            with open(nombre_archivo, 'w') as archivo:
                archivo.write(response)

            return "Robot creado satisfactoriamente"

        else:
            return 'Error al enviar la solicitud'
        
        #return render_template('robot/robot.html')
    else:

         return render_template('robot/robot.html')

def pagina_no_encontrada(error):
    return render_template('404.html'), 404

if __name__ == '__main__':
    app.register_error_handler(404, pagina_no_encontrada)
    app.run(debug=True,host= '127.0.0.1', port=5000)
    #poner 127.0.0.1 que las de nattech no funcionan temporalmente, cuando vuelvan a funcionar poner las de abajo
    #172.16.4.34
    #8089