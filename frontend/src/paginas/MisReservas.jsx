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

const bookingsURL = `http://localhost:8090/booking/user/${Cookie.get("user_id")}`

export const MisReservas=() => {

   if (Cookie.get("user_type") === "true"){
        notifyAdmin();
        goto("/admin/reservas")
    } else if (Cookie.get("user_id") === "-1"){
        notifyNotLoggedIn();
        goto("/login")
    }

    const [searchResults, setSearchResults] = useState([]);

    const [booking, setBooking] =  useState([]);

    const getBooking = async () => {
        const response = await fetch(bookingsURL);
        const booking = await response.json();
        return booking;
    }
    useEffect(() => {
        getBooking().then (booking => {
            setBooking(booking)
            console.log(booking)
            setSearchResults(booking)
        });
    },[])

    const handleSubmit = (e) => {
        e.preventDefault()
    }

    const handleSearchChange = (e) => {
        if (!e.target.value){
            return setSearchResults(booking)
        } 

        const resultArray = booking.filter(booking => booking.hotel_name.includes(e.
            target.value))
            
        setSearchResults(resultArray)
    }


    return (

        <>
        <Navbar/>
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
        <div className='contenedor-principal'>
            <h1>Tus reservas:</h1>
            {
                searchResults?.length ? searchResults.map(searchResults => <ReservasA key={searchResults.booking_id} id_booking={searchResults.booking_id} booking_startdate={searchResults.start_date} booking_enddate={searchResults.end_date}  booking_username={searchResults.user_name} booking_hotelname={searchResults.hotel_name} booking_hoteladdress={searchResults.hotel_address}/>
                 ): <p>Aun no tienes ninguna reserva</p>
            }
        </div>
        <ToastContainer/>
        </>
    )

}