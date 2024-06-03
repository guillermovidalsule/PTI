import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import {  useGaugeState, gaugeClasses, Gauge } from '@mui/x-charts/Gauge';
import { tokens,ColorModeContext, useMode } from '../theme';
import { Box, Button, CssBaseline, ThemeProvider,Typography } from "@mui/material";
import Header from './Header';
import Topbar from './Topbar';
import { Link } from 'react-router-dom';

import {APIProvider, Map, MapCameraChangedEvent, AdvancedMarker, Pin} from '@vis.gl/react-google-maps';



const RobotStatistics = ({ match }) => {
  const [theme, colorMode] = useMode(); 
  const [isSidebar, setIsSidebar] = useState(false);
  const colors = tokens(theme.palette.mode);
  const [map, setMap] = useState((null));
  const center = { lat: 41.38879, lng:  2.15899 }
  let mapaCenter = center;

  // Obtener el nombre del robot de match.params
  const { robotname } = useParams();
  console.log('Componente RobotStatistics renderizado' + {robotname});
  const [robotData, setRobotData] = useState([]);
  useEffect(() => {
    // Función para realizar la solicitud GET y escribir en el archivo .js
    const fetchData = async () => {
      try {
        const xhr = new XMLHttpRequest();
        const token = localStorage.getItem('token');
        xhr.open("GET", `http://nattech.fib.upc.edu:40342/robot/consulta/${robotname}`, true);
        xhr.setRequestHeader("token", token);
        xhr.setRequestHeader("Content-Type", "application/json");
      
        xhr.onreadystatechange = function () {
          if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
              const response = JSON.parse(xhr.responseText);
              console.log('Respuesta del servidor:', response);
              const robot = response.robot;
              const robotID = robot.ID;
              const robotName = robot.robotname;
              const robotIdle = robot.robotidle;
              const robotState = robot.robotstate;
              const createdAt = robot.created_at;
              const updatedAt = robot.updated_at;
              const robotID2 = robot.robot_id;
              const userID = robot.user_id;
              const robotInfo = robot.robot_info;
              const cpuFreq = robotInfo.cpu_freq;
              const temperature = robotInfo.temperature;
              const velocity = robotInfo.velocity;
              const ruta = robot.ruta;
  
              // Creamos un objeto con todos los campos del robot
              const robotDetails = {
                robotID,
                robotName,
                robotIdle,
                robotState,
                createdAt,
                updatedAt,
                robotID2,
                userID,
                cpuFreq,
                temperature,
                velocity,
                ruta
              };
  
              // Establecemos el objeto en el estado
              setRobotData(robotDetails);
              console.log('Dades del robot:', robotData);
            } else {
              console.error('Error al realizar la solicitud:', xhr.status);
            }
          }
        };
      
        xhr.send();
      } catch (error) {
       console.log('Failed: ' + error.message);
      }
    };
    fetchData();
  }, [robotname]);

  console.log('Centre mapa:', robotData.ruta);
  if (robotData.ruta && robotData.ruta.length > 0) {
    const firstPoint = robotData.ruta[0];
    mapaCenter = { lat: firstPoint.latitud, lng: firstPoint.longitud };
  }
  console.log('Mapacenter:', mapaCenter);

  return (
    <ColorModeContext.Provider value={colorMode}>
      <ThemeProvider theme={theme}> 
        <CssBaseline />
        <div className="pagina" style={{ display: 'flex' }}>
          <div className='stats' style={{ flex: '1' }}>
            <Topbar setIsSidebar={setIsSidebar} />
            <Box
              display="flex"
              justifyContent="center"
              alignItems="center"
              minHeight="20vh" 
            >
              <Header title={robotname} subtitle="Your robot statistics" />
            </Box>
            <Box mt={2} ml={5}>
              <Button component={Link} to="http://nattech.fib.upc.edu:40344/video_feed/video" variant="contained" color="primary"
                sx={{
                  backgroundColor: '#FF8C00', // Naranja
                  color: '#FFFFFF', // Blanco
                  fontSize: '1.2rem', // Tamaño de letra más grande
                  '&:hover': {
                    backgroundColor: '#FF4500', // Naranja más oscuro al pasar el ratón
                  },
                }}
              >
                Videostreaming
              </Button>
            </Box>
            <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
                minHeight="20vh" 
            >
              <div style={{ display: 'flex' }}>
                <div style={{ flex: '1', marginRight: '50px' }}>
                  <div className='cpufreq'>
                    <Gauge
                      width={200}
                      height={250}
                      startAngle={-110}
                      endAngle={110}
                      value={robotData.cpuFreq}
                      valueMax={5000}
                      text='CPU Freq'
                      sx={(theme) => ({
                        [`& .${gaugeClasses.valueArc}`]: {
                          fill: '#52b202',
                        },
                        [`& .${gaugeClasses.referenceArc}`]: {
                          fill: theme.palette.text.disabled,
                        },
                      })} 
                    >
                      CPU Freq
                    </Gauge>
                  </div>
                </div>
                <div style={{ flex: '1', marginRight: '50px' }}>
                  <div className='temp'>
                    <Gauge
                      width={200}
                      height={250}
                      startAngle={-110}
                      endAngle={110}
                      value={robotData.temperature}
                      valueMax={200}
                      text='Temperature'
                      sx={(theme) => ({
                        [`& .${gaugeClasses.valueArc}`]: {
                          fill: '#52b202',
                        },
                        [`& .${gaugeClasses.referenceArc}`]: {
                          fill: theme.palette.text.disabled,
                        },
                      })} 
                    >
                      Temperature
                    </Gauge>
                  </div>
                </div>
                <div style={{ flex: '1' }}>
                  <div className='velocitat'>
                    <Gauge
                      width={200}
                      height={250}
                      startAngle={-110}
                      endAngle={110}
                      value={robotData.velocity}
                      valueMax={200}
                      text='Velocity'
                      sx={(theme) => ({
                        [`& .${gaugeClasses.valueArc}`]: {
                          fill: '#52b202',
                        },
                        [`& .${gaugeClasses.referenceArc}`]: {
                          fill: theme.palette.text.disabled,
                        },
                      })} 
                    >
                      Velocity
                    </Gauge>
                  </div>
                </div>
              </div>
            </Box>
            <Box
              display="flex"
              justifyContent="center"
              alignItems="center"
              minHeight="30vh" 
            >
            </Box>
            <APIProvider apiKey={'AIzaSyCTqrFP6pfTUfbHaIfKDybhbZ4giqNDZTE'} onLoad={() => console.log('api loaded')}>
            <Map
              defaultZoom={5}
              mapId='Hamansito'
              defaultCenter={ mapaCenter }
              >
                
                {robotData.ruta && robotData.ruta.length > 0 && robotData.ruta.map((point, index) => (
                  
                  <AdvancedMarker
                    key={index}
                    position={{ lat: point.latitud, lng: point.longitud }}
                     
                  >
                    <Pin background={'#800080'} glyphColor={'#FFF'} borderColor={'#000'} />
                  </AdvancedMarker >
                  
                ))}
                
               
            </Map>
            </APIProvider>
          </div>
        </div>
      </ThemeProvider>
    </ColorModeContext.Provider>
  );
};

export default RobotStatistics;
