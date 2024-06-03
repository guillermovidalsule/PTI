import React from 'react'
import './cta.css'
import {Link} from 'react-router-dom';

const CTA = () => {
  return (
    <div className='talos__cta'>
      <div className='talos__cta-content'>
        <p>Request early Access to get started</p>
        <h3>Register today & start exploring the endless possibilities.</h3>
      </div>
      <div className='talos__cta-btn'>
      <Link to="/signup">
        <button type='button'>Get Started</button>
      </Link>
      </div>
      
    </div>
  )
}

export default CTA