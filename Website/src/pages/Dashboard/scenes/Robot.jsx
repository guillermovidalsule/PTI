import { Box, Button, IconButton, Typography, useTheme } from "@mui/material";
import {DataGrid} from "@mui/x-data-grid"
import { tokens } from '../theme';
import Header from './Header'
import { mockDataTeam } from "../data/robots";
import React, { useEffect, useState } from 'react';
import axios from 'axios'; // Importa axios para hacer la solicitud GET
import RobotStatistics from "./RobotStatistics";
import { Link } from 'react-router-dom';



const Robots= () => {
  const [jsonData, setJsonData] = useState([]);
  useEffect(() => {
    // Función para realizar la solicitud GET y escribir en el archivo .js
    const fetchData = async () => {
      try {
        const xhr = new XMLHttpRequest();
        const token = localStorage.getItem('token');
        xhr.open("GET", "http://nattech.fib.upc.edu:40342/robot/listar", true);
        xhr.setRequestHeader("token", token);
        xhr.setRequestHeader("Content-Type", "application/json");
      
        xhr.onreadystatechange = function () {
          if (xhr.readyState === XMLHttpRequest.DONE) {
            if (xhr.status === 200) {
              const response = JSON.parse(xhr.responseText);
              console.log('Respuesta del servidor:', xhr.responseText);
              const formattedData = response.map(item => ({
                robotname: item.robotname,
                created_at: item.created_at,
                updated_at: item.updated_at,
                robotidle: item.robotidle,
                robot_id: item.robot_id,
                user_id: item.user_id,
              }));
              setJsonData(formattedData);
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
}, []);
  console.log(jsonData);

  const theme = useTheme();
  const colors = tokens(theme.palette.mode);
  const columns = [
    {
      field: "robotname",
      headerName: "Robot Name",
      flex: 1,
      renderCell: ({ row }) => (
        <Link to={`/robots/${row.robotname}`}>{row.robotname}</Link>
    )},
    {
      field: "created_at",
      headerName: "Created At",
      flex: 1,
    },
    {
      field: "updated_at",
      headerName: "Updated At",
      flex: 1,
    },
    {
      field: "robotidle",
      headerName: "Robot Idle",
      flex: 1,
      renderCell: ({ row: { robotidle } }) => (
        <Button variant="contained" style={{ backgroundColor: robotidle ? 'green' : 'gray' }}>
          {robotidle ? 'Active' : 'Inactive'}
        </Button>
      )
    },
    {
      field: "robot_id",
      headerName: "Robot ID",
      flex: 1,
    },
    {
      field: "user_id",
      headerName: "User ID",
      flex: 1,
    },
  ];

  return (
    <Box m="20px">
      <Header title="Robots" subtitle="Your robots" />
      <Box 
       m="40px 0 0 0"
       height="75vh"
       sx={{
        "& .MuiDataGrid-root": {
          border: "none",
        },
        "& .MuiDataGrid-cell": {
          borderBottom: "none",
        },
        "& .name-column--cell": {
          color: colors.greenAccent[300],
        },
        "& .MuiDataGrid-columnHeaders": {
          backgroundColor: colors.blueAccent[700],
          borderBottom: "none",
        },
        "& .MuiDataGrid-virtualScroller": {
          backgroundColor: colors.primary[400],
        },
        "& .MuiDataGrid-footerContainer": {
          borderTop: "none",
          backgroundColor: colors.blueAccent[700],
        },
        "& .MuiCheckbox-root": {
          color: `${colors.greenAccent[200]} !important`,
        },
      }}
       >
        {/*<DataGrid rows={mockDataTeam} columns={columns} getRowId={(row) => row.robot_id}  />*/}
        <DataGrid rows={jsonData} columns={columns} getRowId={(row) => row.robot_id}  />   
      </Box>    
    </Box>

  )
}

export default Robots
