import React, {useEffect, useState} from 'react'
import Navbar from "../componentes/Navbar";
import ReservasA from "../componentes/ReservasA";
import 'react-date-range/dist/styles.css'; // main style file
import 'react-date-range/dist/theme/default.css'; // theme css file
import { DateRangePicker } from 'react-date-range';

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

    function numberToDate(input) {

        let inputString = input.toString()
        
        let year = Number(String(input).substring(0,4), 10)
        let month = Number(String(input).substring(5,7), 10)
        let day = Number(String(input).substring(8, 10), 10)

        let date = new Date(year, month - 1, day)

        return date
    }

    const handleSelect = (date) =>{
        console.log(date); 
      };
    
      const selectionRange = {
        startDate: new Date(),
        endDate: new Date(),
        key: 'selection',
      }
    

    return (
        <div className='contenedor-principal'>
            <Navbar />
            <div>
                <input type="date" placeholder="Fecha de inicio:" label="Fecha de inicio:" className='inputFechaReservas'></input>
                <input type="date" placeholder="Fecha de fin:" label="Fecha de fin:" className='inputFechaReservas'></input>
            </div>

                {
                    booking.length ? booking.map((booking) => <ReservasA key={booking.id} id_booking={booking.booking_id} booking_startdate={booking.start_date} booking_enddate={booking.end_date}  booking_username={booking.user_name} booking_hotelname={booking.hotel_name} booking_hoteladdress={booking.hotel_address}/> ):null
                }
        </div>
    )
}