import React from 'react'
import './whatTalos.css'
import { Feature } from '../../components'

const WhatTalos = () => {
  return (
    <div className='talos__whattalos section__margin' id="talos">
      <div className='talos__whattalos-feature'>
        <Feature title="What is Talos" text= "The aim of this project is to develop a security and reliable system that is capable of surveilling a determined perimeter and detecting potential trespassers."/>
      </div>
      <div className='talos__whattalos-heading'>
        <h1 className='gradient__text'>Possibilities byond imagination</h1>
        <p>Explore the library</p>
      </div>
      <div className='talos__whattalos-container'>
        <Feature title="Components" text= "COMPONENTS OF THE ROBOT" />
        <Feature title="Historical reasons" text= "This project, named after the mythological automaton giant Talos, who was in charge of protecting Europa and surveilling the coast of Crete comprises two main components that must be implemented: autonomous robots and a user-friendly interface that manages and monitors them."/>
        <Feature title="Further development" text= "Involve integrating a trained machine learning algorithm that distinguishes between a potential trespasser and authorized workers of the enterprise."/>
        {/* Minut video 2:08:21 */}
      </div>
    </div>
  )
}

export default WhatTalos
