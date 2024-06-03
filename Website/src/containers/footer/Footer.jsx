import React from 'react'
import './footer.css'
import logo from '../../assets/prova2.svg';
import {Link} from 'react-router-dom';

const Footer = () => {
  return (
    <div className='talos__footer section__padding'>
      <div className='talos__footer-heading'>
        <h1 className='gradient__text'>Unlock endless possibilities. Join us today and embark on a journey towards innovation and growth.</h1>
      </div>
      <div className='talos__footer-btn'>
      <Link to="/signup">
        <p>Request early access</p>
      </Link>
      </div>
      <div className='talos__footer-links'>
        <div className='talos__footer-links_logo'>
          <img src={logo} alt="logo" />
          <p>Facultat d'informàtica de Barcelona(FIB), all rights reserved</p>
        </div>
        <div className='talos__footer-links_div'>
          <h4>Links</h4>
          <p>Social Media</p>
          <p>Contact</p>
        </div>
        <div className='talos__footer-links_div'>
          <h4>Company</h4>
          <p>Terms & Conditions</p>
          <p>Privacy Policy</p>
          <p>Contact</p>
        </div>
      </div>
      <div className='talos__footer-copyright'>
        <p>© 2024 Talos Group. All rights reserved</p>
      </div>
      
    </div>
  )
}

export default Footer
