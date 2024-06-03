import React, { useState } from 'react';
import { useLocation } from 'react-router-dom';
import withAuth from '../pages/Dashboard/withAuth'; // Importa el HOC de autenticación
import { ColorModeContext, useMode } from '../theme'; 
import { CssBaseline, ThemeProvider } from "@mui/material";
import Sidebar from './Sidebar';
import Topbar from './Topbar';
import IndexDash from './index';
import Robots from './Robot';

const Dashboard2 = () => {
  const [theme, colorMode] = useMode(); 
  const [isSidebar, setIsSidebar] = useState(true);

  return (
    <ColorModeContext.Provider value={colorMode}>
      <ThemeProvider theme={theme}> 
        <CssBaseline />
        <div className="app" style={{ display: 'flex' }}>
          <Sidebar isSidebar={isSidebar} />
          <div className="content" style={{ flex: '1' }}>
            <Topbar setIsSidebar={setIsSidebar} />
            <IndexDash/> 
            <Robots />
          </div>
        </div>
      </ThemeProvider>
    </ColorModeContext.Provider>
  );
};

export default withAuth(Dashboard2);