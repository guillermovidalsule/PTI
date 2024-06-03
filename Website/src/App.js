import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import SignIn from './pages/Signin/SignIn';
import Home from './pages/Home/Home';
import SignUp from './pages/Signup/SignUp';
//import Dashboard from './scenes/Dasboard2'; // Importa el componente del dashboard
import './app.css';
import Dashboard2 from './scenes/Dashboard2';
import RobotStatistics  from './scenes/RobotStatistics';
import { APIProvider } from '@vis.gl/react-google-maps';

const App = () => {
  return (
    <div>
      <BrowserRouter>
        <Routes>
          <Route index element={<Home />} />
          <Route path="/home" element={<Home />} />
          <Route path="/signin" element={<SignIn />} />
          <Route path="/signup" element={<SignUp />} />
          <Route path="/dash" element={<Dashboard2 />} />
          <Route path="/robots/:robotname" element={<RobotStatistics />} />
        </Routes>
      </BrowserRouter>
      
    </div>
  );
}

export default App;

