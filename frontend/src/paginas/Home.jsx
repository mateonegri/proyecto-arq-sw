import React from 'react'
import {Hotel} from "../componentes/Hotel";
import Navbar from "../componentes/Navbar";
import Hotels from "../componentes/Hotels";
import '../hojas-de-estilo/Hotel.css'
export const Home = () => {



    return (
        <div className='Home'>
            <Navbar />
            <div className='contenedor-principal'>
                <Hotels />

            </div>

            </div>
    )
}

