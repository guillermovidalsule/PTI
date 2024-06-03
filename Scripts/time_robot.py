import json
import datetime
import requests

file_path = "/robot.json"
server_url = "http://172.16.4.34:8082/users/login"


def read_json_file(file_path):
    try:
        with open(file_path, 'r') as file:
            json_data = json.load(file)
        return json_data
    except FileNotFoundError:
        print(f"Error: No s'ha trobat el fitxer '{file_path}'")
        return None
    except json.JSONDecodeError:
        print(f"Error: Hi ha un format invàlid en l'arxiu '{file_path}'")
        return None
    except Exception as e:
        print(f"Error: {str(e)}")
        return None
    

def sumar_dos_horas(time_from_json):
    dos_horas = datetime.timedelta(hours=2)
    nuevo_tiempo = time_from_json + dos_horas 
    return nuevo_tiempo


def get_user(username, password, server_url):
    try:
        json_data = {
            "username": username,
            "password": password
        }
        response = requests.post(server_url, json=json_data)
        response.raise_for_status()
        return response.json()
    
    except requests.exceptions.RequestException as e:
        print(f"Error: {str(e)}")
        return None
    
    except Exception as e:
        print(f"Error: {str(e)}")
        return None


def consultar_robot(token, robot_consulta_url):
    try:
        headers = {
            'token': token,
            'Content-Type': 'application/json'
        }
        response = requests.get(robot_consulta_url, headers=headers)
        response.raise_for_status()
        return response.json()
    
    except requests.exceptions.RequestException as e:
        print(f"Error: {str(e)}")
        return None
    
    except Exception as e:
        print(f"Error: {str(e)}")
        return None
    

def desactivar_robot(token, robot_desactivar_url):
    try:
        headers = {
            'token': token,
            'Content-Type': 'application/json'
        }
        response = requests.post(robot_desactivar_url, headers=headers)
        response.raise_for_status()
        return response.json()
    
    except requests.exceptions.RequestException as e:
        print(f"Error: {str(e)}")
        return None
    
    except Exception as e:
        print(f"Error: {str(e)}")
        return None


def activar_robot(token, robot_desactivar_url):
    try:
        headers = {
            'token': token,
            'Content-Type': 'application/json'
        }
        response = requests.post(robot_desactivar_url, headers=headers)
        response.raise_for_status()
        return response.json()
    
    except requests.exceptions.RequestException as e:
        print(f"Error: {str(e)}")
        return None
    
    except Exception as e:
        print(f"Error: {str(e)}")
        return None

json_object = read_json_file(file_path=file_path)

if __name__ == "__main__":
    if json_object is None:
        exit(1)
    username = json_object['username']
    password = json_object['password']
    robotname = json_object['robotname']
    user = get_user(username, password, server_url)

    if user is None:
        exit(1)

    token = user.get('token')
    robot_consulta_url=f"http://172.16.4.34:8082/robot/consulta/{robotname}"
    robot = consultar_robot(token, robot_consulta_url)
    if robot is None:
        exit(1)
    
    datos_robot = json.loads(json.dumps(robot))
    time = datos_robot['robot']['updated_at']

    #Conversión del formato json que generamos con GO al que trabaja Python
    #Necesitamos sumarle dos horas porque va con dos horas de retraso la funcion de obtener hora de GO
    time_json = datetime.datetime.strptime(time, '%Y-%m-%dT%H:%M:%S.%fZ')
    new_time = sumar_dos_horas(time_json)

    #Hora actual obtenida en Python
    hora_actual = datetime.datetime.now()
    #Diferencia entre los dos tiempos
    diferencia_tiempo = hora_actual - new_time
    print("La diferencia de tiempo es:", diferencia_tiempo)

    #Comparamos si la diferencia de timepo es mayor o menos de 30
    segundos_30 = datetime.timedelta(seconds=30)
    if diferencia_tiempo > segundos_30:
        print("Pondremos activate a FALSE")
        url_set_false = f"http://172.16.4.34:8082/robot/desactivar/{robotname}"
        respuesta = desactivar_robot(token,url_set_false)
        if respuesta is None:
            exit(1)
        print("Robot DESACTIVADO")
        
    else:
        url_set_true = f"http://172.16.4.34:8082/robot/activar/{robotname}"
        respuesta = activar_robot(token,url_set_true)
        if respuesta is None:
            exit(1)
        print("Robot ACTIVO")


    exit(0)