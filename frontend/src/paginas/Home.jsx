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


                {/*<Hotel*/}
                {/*    nombre='The Lotte Palace'*/}
                {/*    barrio='Manhattan'*/}
                {/*    imagen='lotte'*/}
                {/*    ciudad='New York'*/}
                {/*    descripcion='The Lotte New York Palace Hotel, a grand edifice once known as the Helmsley Palace Hotel and later the New York Palace Hotel, incorporates the Gilded Age’s famed Villard Houses into a 55-story luxury hotel. Today, the hotel combines timeless extravagance with 21st-century luxury in Midtown East.'/>*/}
                {/*<Hotel*/}
                {/*    nombre='Gran Melia'*/}
                {/*    barrio='Parque Nacional Iguazu'*/}
                {/*    ciudad='Argentina'*/}
                {/*    imagen='melia'*/}
                {/*    descripcion='Al nordeste de la República Argentina, el río Iguazú cae desde el río Paraná formando las cataratas más famosas del mundo. Junto a él se encuentra este excepcional hotel, el único dentro del Parque Nacional Iguazú. Las privilegiadas vistas a las cataratas de Iguazú desde sus habitaciones y suites, su piscina infinita y sus propuestas gastronómicas, envuelven una experiencia única: alojarse ante una de las Siete Maravillas Naturales del Mundo.'/>*/}
                {/*<Hotel*/}
                {/*    nombre='Los Cauquenes'*/}
                {/*    barrio='Ushuaia'*/}
                {/*    ciudad='Argentina'*/}
                {/*    imagen='cauquenes'*/}
                {/*    descripcion='Los Cauquenes Resort + Spa + Experiences ofrece diferentes actividades para vivir una experiencia única: Turismo activo y de aventura, experiencias gastronómicas, tratamientos de spa y bienestar, piscina in-out y un entorno natural inigualable.'*/}
                {/*/>*/}
            </div>
    )
}

