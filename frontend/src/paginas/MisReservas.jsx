import React, {useEffect, useState} from "react";
import Cookies from "universal-cookie";
import Navbar from "../componentes/Navbar";
import ReservasA from "../componentes/ReservasA";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'
const Cookie = new Cookies();
const notifyNotLoggedIn = () => {
    toast.error("Recuerda que antes debes loggearte!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}
const notifyAdmin = () => {
    toast.dark("Redireccionandote a la vista de administrador!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}



function goto(path){
    setTimeout( () => {
        window.location = window.location.origin + path
    },2000)

}

const bookings = `http://localhost:8090/booking/user/${Cookie.get("user_id")}`
export const MisReservas=() => {

   if (Cookie.get("user_type") === "true"){
        notifyAdmin();
        goto("/admin/reservas")
    } else if (Cookie.get("user_id") === "-1"){
        notifyNotLoggedIn();
        goto("/login")
    }

    const [booking, setBooking] =  useState([]);
    const getBooking = async () => {
        const response = await fetch(bookings);
        const booking = await response.json();
        return booking;
    }
    useEffect(() => {
        getBooking().then ((booking) => setBooking(booking));
    },[])

    return (

        <>
            <Navbar/>

        <div className='contenedor-principal'>

        <div className='contenedor-reservas'>

            <h1>Tus reservas:</h1>
            {
                booking?.length ? booking.map((booking) => <ReservasA key={booking.booking_id} id_booking={booking.booking_id} booking_startdate={booking.start_date} booking_enddate={booking.end_date}  booking_username={booking.user_name} booking_hotelname={booking.hotel_name} booking_hoteladdress={booking.hotel_address}/> ): <p>Aun no tienes ninguna</p>
            }
        </div>
        </div>
            <ToastContainer/>
            </>
    )

}