import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import '../hojas-de-estilo/Reservas.css'
import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Cookies from "universal-cookie";

const ReservasA = ({ id_booking,  booking_startdate, booking_enddate, booking_username, booking_hotelname, booking_hoteladdress }) => {


    return (
       /*  <div className='reservas'>
            <h1>Reserva en hotel: {booking_hotelname}</h1>
            <p>Realizada por {booking_username} desde la fecha {booking_startdate} hasta {booking_enddate}</p>
        </div> */
        <>
            <Card sx={{width:'92%'}} className="cartaHotel">
                <CardContent>
                    <Typography gutterBottom variant="h5" component="div">
                        Reserva en hotel: {booking_hotelname}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                        {/* <p>{hotel_address}</p> */}
                        <p>Esta reserva fue realizada por {booking_username}</p>
                        <p>Fecha de inicio de reserva: {booking_startdate}</p>
                        <p>Fecha de final de reserva: {booking_enddate}</p>
                    </Typography>
                </CardContent>
            </Card>
        </>
    )

}

export default ReservasA;