import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import '../hojas-de-estilo/Reservas.css'
const ReservasA = ({ id_booking,  booking_startdate, booking_enddate, booking_username, booking_hotelname, booking_hoteladdress }) => {
    return (
        <div className='reservas'>
            <h1>Reserva en hotel: {booking_hotelname}</h1>
            <p>Realizada por {booking_username} desde la fecha {booking_startdate} hasta {booking_enddate}</p>
        </div>

    )

}

export default ReservasA;