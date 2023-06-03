import React from 'react'
import {Hotel} from "../componentes/Hotel";
import Navbar from "../componentes/Navbar";
import Hotels from "../componentes/Hotels";
import Cookies from "universal-cookie";
import '../hojas-de-estilo/Hotel.css'

const Cookie = new Cookies()


export const Home = () => {

    if (Cookie.get("user_type") == false){
        console.log("Anda")
    }

    console.log(Cookie.get("user_type"))

    return (
        <div className='Home'>
            <Navbar />
            <div className='contenedor-principal'>
                <Hotels />

            </div>

            </div>
    )
}

