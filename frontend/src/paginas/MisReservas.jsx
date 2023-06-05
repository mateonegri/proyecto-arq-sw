import React, {useEffect, useState} from "react";
import Cookies from "universal-cookie";
import Navbar from "../componentes/Navbar";
import ReservasA from "../componentes/ReservasA";
const Cookie = new Cookies();
function goto(path){
    window.location = window.location.origin + path
}

const bookings = `http://localhost:8090/booking/${(Cookie.get("user_id"))}`
export const MisReservas=() => {

   if (Cookie.get("user_type") === "true"){
        goto("/admin/reservas")
    } else if (Cookie.get("user_id") === "-1"){
        goto("/login")
    }

    const [booking, setBooking] = useState([]);
    const getBooking = async () => {
        const response = await fetch(bookings);
        const booking = await response.json();

        console.log(booking)
        return booking;
    }
    useEffect(() => {
        getBooking().then ((booking) => setBooking(booking));
    },[])
    console.log(booking)
    return (


        <>
            <Navbar/>
        <div className='contenedor-principal'>

            <h1>Estas son tus reservas:</h1>
            {
                booking.length ? booking.map((booking) => <ReservasA key={booking.booking_id} id_booking={booking.booking_id} booking_startdate={booking.start_date} booking_enddate={booking.end_date}  booking_username={booking.user_name} booking_hotelname={booking.hotel_name} booking_hoteladdress={booking.hotel_address}/> ):null
            }
        </div>
        </>

    )

}