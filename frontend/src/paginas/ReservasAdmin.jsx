import React, {useEffect, useState} from 'react'
import Navbar from "../componentes/Navbar";
import ReservasA from "../componentes/ReservasA";
import 'react-date-range/dist/styles.css'; // main style file
import 'react-date-range/dist/theme/default.css'; // theme css file
import { DateRangePicker } from 'react-date-range';
import FormInput from "../componentes/FormInput.jsx";

const bookings = "http://localhost:8090/booking"

function convertirFecha(fecha) {
    let fechaString = fecha.toString()
    

    let year = fechaString.substring(0,4)
    
    let month = fechaString.substring(5,7)
    
    let day = fechaString.substring(8,10)

    let yearPlusMonth = year.concat("",month)
    let fechaStringFinal = yearPlusMonth.concat("",day)

    return fechaStringFinal
}

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

    const [values, setValues] = useState({
        "search": "",
        "date": "",
    })

    const inputs = [
        {
            id: 1,
            name: "search",
            type: "search",
            placeholder: "Busqueda por hotel" ,
        },
        {
            id: 2,
            name: "date",
            type: "date",
            label:"Busqueda por fecha",
            placeholder: "Busqueda por fecha",
        }
    ]

    const handleSubmit = (e) => {

        e.preventDefault();

        if (!values.search && !values.date){
            return setSearchResults(booking)
        } 

        if (!values.date) {

            const resultArray = booking.filter(booking => booking.hotel_name.includes(values.search))

            return setSearchResults(resultArray)
        }

        if (!values.search) {
        
            let valorFecha = convertirFecha(values.date)

            const resultArray = booking.filter( function (booking) {
                return Number(booking.start_date) <= Number(valorFecha)
                && Number(booking.end_date) >= Number(valorFecha)
            } )

            return setSearchResults(resultArray)
        }

        let valorFecha = convertirFecha(values.date)

        const resultArray = booking.filter(booking => booking.hotel_name.includes(values.search)
        && Number(booking.start_date) <= Number(valorFecha)
        && Number(booking.end_date) >= Number(valorFecha))

        console.log(resultArray)
        
        setSearchResults(resultArray)
    }

    const onChange = (e) => {
        setValues({...values, [e.target.name]: e.target.value})
    }
    
    

    return (
        <>
        <Navbar />
        <div className="inputReservas">
        <form >
            {inputs.map((input ) => (
            <FormInput key={input.id} {...input} value={values[input.name]}  onChange={onChange}/>
            ))}
            <button className="reservar-button" onClick={handleSubmit} >Buscar</button>
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