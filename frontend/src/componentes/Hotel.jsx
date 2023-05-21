import React from 'react';
import '../hojas-de-estilo/Hotel.css'
export function Hotel(props)  {
    return (
        <div className='contenedor-hotel'>
            <img
            className='imagen-hotel'
            src={require(`../imagenes/hotel_${props.imagen}.jpg`)}
            alt='Foto hotel'/>
            <div className='contenedor-informacion-hotel'>
                <p className='nombre-hotel'>{props.nombre}</p>
                <p className='ubicacion-hotel'>{props.barrio}, <strong>{props.ciudad}</strong></p>
                <p className='descripcion-hotel'>{props.descripcion} </p>
            </div>
        </div>
    );
}

