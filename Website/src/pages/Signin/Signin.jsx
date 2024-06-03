import './signin.css';
import React, { useState } from 'react';
import {FaUser, FaLock} from "react-icons/fa";
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';


const Signin = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();
  const [error, setError] = useState('');


  const handleSubmit = async (e) => {
    e.preventDefault();


    try {
      const xhr = new XMLHttpRequest();
      xhr.open("POST", "http://nattech.fib.upc.edu:40342/users/login", true);
      xhr.setRequestHeader("Content-Type", "application/json");

      xhr.onreadystatechange = function () {
        if (xhr.readyState === XMLHttpRequest.DONE) {
          if (xhr.status === 200) {
            const data = JSON.parse(xhr.responseText);
            console.log(data);
            const token = data.token; // Suponiendo que el token se devuelve en la propiedad "token" del objeto de respuesta
            localStorage.setItem('token', token);
            toast.success('Registered successfully.');
            navigate('/dash');
          } else {
            setError('Invalid username or password.');
            
          }
        }
      };

      xhr.onerror = function () {
        toast.error('Failed to make request.');
      };

      xhr.send(JSON.stringify({ username, password }));
    } catch (error) {
      toast.error('Failed: ' + error.message);
    }
  };



  return (
    <div className='total' >
      <div className='signin'>
        <form action='' onSubmit={handleSubmit}>
          <h1>Sign in</h1>
          <div className='input-box'>
            <input type='text' placeholder='Username' 
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required />
            <FaUser className='icon' />
          </div>
          <div className='input-box'>
            <input type='password' placeholder='Password' 
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required />
            <FaLock  className='icon'/>
          </div>
          {error && <p className='error'>{error}</p>}
          <div className='remember-forgot'>
            <label><input type='checkbox'/>Remeber me</label>
            <a href='#'> Forgot password?</a>
          </div>
          <button type='submit'>Sign in</button>
          <div className='register-link'>
            <p>Don't have an account? <a href='/signup'>Register</a></p>
          </div>
        </form>
      </div>
    </div>
  );
}

export default Signin;