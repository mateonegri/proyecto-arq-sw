import React, {useEffect, useState} from 'react'
import {useParams} from "react-router-dom";
import Navbar from "../componentes/Navbar";


export const HotelDetalle = () => {
    const [hotel, setHotel] = useState()
    const {id} = useParams()

    const infoHotel = `http://localhost:3000/hotel.json/:${id}`

    const getHotel = async () => {
        const response = await fetch (infoHotel)
        const hotel = await response.json()
        return hotel;
    }
    console.log(hotel)
    useEffect(() => {
        getHotel().then((hotel) => setHotel(hotel));
    }, [])
    return (
        <div className='contenedor-principal'>
            <Navbar />
            <h1>Detalle del hotel </h1>
            <div className='hotelesDetalle'>

                <div className='cadaDetalle'>
                    <p>{hotel?.hotel_name}</p>
                    <p>{hotel?.hotel_description}</p>
                    <p>{hotel?.hotel_address}</p>

                    <p>Apurate que a este hotel solo le quedan {hotel?.hotel_available_rooms} habitaciones disponibles </p>

                </div>



            </div>
        </div>
    )
}

