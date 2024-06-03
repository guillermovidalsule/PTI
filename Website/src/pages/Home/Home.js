import React from 'react'
import {Footer, Blog, Possibility, Features, Header, WhatTalos,Logo} from '../../containers';
import { CTA, Brand, Navbar} from '../../components';
import '../../app.css';


const Home = () => {
  return (
      <div className='App'>
        <div className='gradient__bg'>
          <Navbar />
          <Logo />
          <Header />
        </div>
        <Brand />
        <WhatTalos />
        <Features />
        <Possibility />
        <CTA />  
        <Blog />
        <Footer />
      </div>
  )
}

export default Home