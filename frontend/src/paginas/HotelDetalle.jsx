import React, {useEffect, useState} from 'react'
import {useParams} from "react-router-dom";
import Calendar from "react-calendar";
import Navbar from "../componentes/Navbar";
import '../hojas-de-estilo/HotelDetalle.css'
import 'react-calendar/dist/Calendar.css';
export const HotelDetalle = () => {
    const [hotel, setHotel] = useState()
    const {id} = useParams()

    const [value, onChange] = useState(new Date());

    const infoHotel = `http://localhost:3000/hotel.json/${id}`

    const getHotel = async () => {
        const response = await fetch(infoHotel);
        const hotel = await response.json();
        console.log(hotel)
        return hotel;
    }

    useEffect(() => {
        getHotel().then((hotel) => setHotel(hotel));
    }, [])
    return (
        <div className='contenedor-hoteldetalle'>
            <Navbar />
            <div className='calendar'>
                <ul><p>Selecciona la fecha de estadia</p>

                <Calendar onChange={onChange} value={value} />
                </ul>
            </div>
            <h1>Detalle del hotel </h1>
            <div className='hoteles-detalle'>

                <div className='nombre'>
                    <p>{hotel?.hotel_name}</p>
                </div>
                <div className='descripcion'>
                    <p>Esta ubicado en {hotel?.hotel_address}</p>
                    <p>{hotel?.hotel_description}</p>
                 </div>
                <p>Apurate que a este hotel solo le quedan {hotel?.hotel_available_rooms} habitaciones disponibles </p>
            </div>
        </div>
    )
}

