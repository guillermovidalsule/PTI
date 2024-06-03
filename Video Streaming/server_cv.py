from flask import Flask, request, Response, jsonify
import base64
import cv2
import numpy as np
import io

app = Flask(__name__)

latest_frame = None

def generate_frames():
    global latest_frame
    while True:
        if latest_frame is not None:
            ret, buffer = cv2.imencode('.jpg', latest_frame)
            frame = buffer.tobytes()
            yield (b'--frame\r\n'
                   b'Content-Type: image/jpeg\r\n\r\n' + frame + b'\r\n')

@app.route('/video_feed', methods=['POST'])
def video_feed():
    global latest_frame

    # Obtener los datos JSON enviados por el cliente
    data = request.get_json()

    # Obtener el frame codificado en base64
    frame_base64 = data['frame']

    # Decodificar el frame de base64 a bytes
    try:
        frame_bytes = base64.b64decode(frame_base64)
    except binascii.Error as e:
        print(f"Error decoding base64 data: {e}")
        return jsonify({'status': 'error'})

    try:
        frame_np = np.frombuffer(frame_bytes, dtype=np.uint8)
        frame = cv2.imdecode(frame_np, cv2.IMREAD_COLOR)
    except cv2.error as e:
        print(f"Error decoding JPEG data: {e}")
        return jsonify({'status': 'error'})

    # Procesar el frame si es necesario (por ejemplo, para detectar objetos)
    # ...

    latest_frame = frame

    # Enviar una respuesta al cliente (opcional)
    response = {'status': 'success'}
    return jsonify(response)

@app.route('/video_feed/video')
def video_feed_video():
    global latest_frame

    if latest_frame is None:
        print("None")
        return Response(status=400)

    return Response(generate_frames(),
                    mimetype='multipart/x-mixed-replace; boundary=frame')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8084)
