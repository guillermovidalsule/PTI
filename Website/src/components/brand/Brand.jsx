import React from 'react'
import './brand.css'
import {google, linkedin, amazon, fib, slack} from './imports';

const Brand = () => {
  return (
    <div className='talos__brand section__padding'>
      <div>
        <img src={google} alt="google"/>
      </div>
      <div>
        <img src={slack} alt="slack"/>
      </div>
      <div>
        <img src={amazon} alt="amazon"/>
      </div>
      <div>
        <img src={fib} alt="fib"/>
      </div>

    </div>
  )
}

export default Brand