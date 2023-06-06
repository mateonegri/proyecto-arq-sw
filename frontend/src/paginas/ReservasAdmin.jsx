import React, {useEffect, useState} from 'react'
import Navbar from "../componentes/Navbar";
import ReservasA from "../componentes/ReservasA";
import 'react-date-range/dist/styles.css'; // main style file
import 'react-date-range/dist/theme/default.css'; // theme css file
import { DateRangePicker } from 'react-date-range';

const bookings = "http://localhost:8090/booking"

export const ReservasAdmin = () => {
    const [booking, setBooking] = useState([]);
    const [searchResults, setSearchResults] = useState([]);

    const getBooking = async () => {
        const response = await fetch(bookings);
        const booking = await response.json();
        // setHotel(hotel)
        console.log(booking)
        return booking;
    }


    useEffect(() => {
        getBooking().then ((booking) => {
            setBooking(booking)
            setSearchResults(booking)
        });
    },[])

    function numberToDate(input) {

        let inputString = input.toString()
        
        let year = Number(String(input).substring(0,4), 10)
        let month = Number(String(input).substring(5,7), 10)
        let day = Number(String(input).substring(8, 10), 10)

        let date = new Date(year, month - 1, day)

        return date
    }

    const handleSubmit = (e) => {
        e.preventDefault()
    }

    const handleSearchChange = (e) => {
        if (!e.target.value) return setSearchResults(booking)

        const resultArray = booking.filter(booking => booking.hotel_name.includes(e.
            target.value))
            
        setSearchResults(resultArray)
    }
    

    return (
        <>
        <Navbar />
        <div className="inputReservas">
            <form onSubmit={handleSubmit}>
                <span>Busqueda por hotel</span>
                <input type="search" 
                    id = "search"
                    className="searchHotel" 
                    placeholder="Busqueda por hotel" 
                    onChange={handleSearchChange} />
                <button className="reservar-button">Buscar</button>
            </form>
        </div>
        <div className = "contenedor-principal">
            {
                searchResults?.length ? searchResults.map(searchResults => <ReservasA key={searchResults.booking_id} id_booking={searchResults.booking_id} booking_startdate={searchResults.start_date} booking_enddate={searchResults.end_date}  booking_username={searchResults.user_name} booking_hotelname={searchResults.hotel_name} booking_hoteladdress={searchResults.hotel_address}/>
                 ): <p>Aun no tienes ninguna</p>
            }
        </div>
        </>
    )
}