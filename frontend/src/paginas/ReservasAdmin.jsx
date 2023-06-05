import React, {useEffect, useState} from 'react'
import Navbar from "../componentes/Navbar";
import ReservasA from "../componentes/ReservasA";

const bookings = "http://localhost:8090/booking"
export const ReservasAdmin = () => {
    const [booking, setBooking] = useState([]);

    const getBooking = async () => {
        const response = await fetch(bookings);
        const booking = await response.json();
        // setHotel(hotel)
        console.log(booking)
        return booking;
    }
    useEffect(() => {
        getBooking().then ((booking) => setBooking(booking));
    },[])
    return (
        <div className='contenedor-principal'>
            <Navbar />
            <h1> Todas las reservas realizadas en la pagina web: </h1>
            <p>(Solo tu puedes ver esto)</p>

                {
                    booking.length ? booking.map((booking) => <ReservasA key={booking.id} id_booking={booking.booking_id} booking_startdate={booking.start_date} booking_enddate={booking.end_date}  booking_username={booking.user_name} booking_hotelname={booking.hotel_name} booking_hoteladdress={booking.hotel_address}/> ):null
                }

        </div>
    )
}