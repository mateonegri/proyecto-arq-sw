import React, {useEffect, useState} from 'react'
import {useParams} from "react-router-dom";
import Navbar from "../componentes/Navbar";
import '../hojas-de-estilo/HotelDetalle.css'
import 'react-calendar/dist/Calendar.css';
import { useNavigate } from "react-router-dom";

export const HotelDetalle = ( hotel_id ) => {
    const [hotel, setHotel] = useState()
    const {id} = useParams()
    const navigate = useNavigate();

    const infoHotel = `http://localhost:8090/hotel/${hotel_id}`

    const getHotel = async () => {
        const response = await fetch(infoHotel);
        const hotel = await response.json();
        console.log(hotel)
        return hotel;
    }

    useEffect(() => {
        getHotel().then((hotel) => setHotel(hotel));
    }, [])

    const reservar = () => {
        navigate(`/home/hotel/reserva/${hotel_id}`);
    };


    return (
        <div className='contenedor-principal'>

            <Navbar />
            <div className='contenedor-hoteldetalle'>

            <div className='hoteles-detalle'>

                <div className='nombre'>
                    <p>{hotel?.hotel_name}</p>
                </div>
                <div className='descripcion'>
                    <p>Esta ubicado en {hotel?.hotel_address}</p>
                    <p>{hotel?.hotel_description}</p>
                 </div>
                <p>Apurate que a este hotel solo le quedan {hotel?.hotel_available_rooms} habitaciones disponibles </p>
                <div className='reservar-button'>
                    <button onClick={reservar}>
                        RESERVAR
                    </button>
                 </div>

                </div>
        </div>
        </div>
    )
}

