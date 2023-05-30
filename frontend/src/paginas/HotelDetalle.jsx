import React, {useEffect, useState} from 'react'
import {useParams} from "react-router-dom";
import Navbar from "../componentes/Navbar";
import '../hojas-de-estilo/HotelDetalle.css'
import 'react-calendar/dist/Calendar.css';
import { useNavigate } from "react-router-dom";
import FormInput from "../componentes/FormInput.jsx";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'
import Cookies from "universal-cookie";
const Cookie = new Cookies();


function convertirFecha(fecha) {
    let fechaString = fecha.toString()
    console.log(fechaString)

    let year = fechaString.substring(0,4)
    console.log(year)
    let month = fechaString.substring(5,7)
    console.log(month)
    let day = fechaString.substring(8,10)

    let yearPlusMonth = year.concat("",month)
    let fechaStringFinal = yearPlusMonth.concat("",day)

    console.log(fechaStringFinal)

    var fechaEntero = parseInt()

    return fechaEntero
}
const notifyBooked = () => {
    toast.success("Reservado!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}

const notifyError = () => {
    toast.error("Hotel no disponible para reserva en fecha seleccionada!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}

export const HotelDetalle = ( hotel_id ) => {
    const [hotel, setHotel] = useState()
    const {id} = useParams()
    const navigate = useNavigate();

    const infoHotel = `http://localhost:8090/hotel/${id}`



    const postBooking = "http://localhost:8090/booking"

    function goto(path){
        window.location = window.location.origin + path
    }

    const chequear =() => {
        if (Cookie.get("user_id") === -1 ) {
            goto("/login")
        }
    }
    async function insertBooking(jsonData) {

        const response = await fetch(postBooking, {
            method: "POST",
            headers:{"content-type":"application/json"},
            body: JSON.stringify(jsonData)
        }).then(response => {
            if (response.status === 400 || response.status === 401 || response.status === 403) {
                console.log("Error al reservar");

                notifyError();

                return response.json();

            } else {
                console.log("Booking added");

                notifyBooked();


                return response.json();
            }

        })

    }

    const getHotel = async () => {
        const response = await fetch(infoHotel);
        const hotel = await response.json();
        console.log(hotel)
        return hotel;
    }

    useEffect(() => {
        getHotel().then((hotel) => setHotel(hotel));
    }, [])


        const [values, setValues] = useState({
           "start_date":"",
            "end_date":"",
        })

        const inputs = [
            {
                id:1,
                name:"start_date",
                type:"date",
                placeholder:"Fecha Inicio",
                errorMessage:"La fecha de inicio no puede ser nula!",
                label:"Fecha Inicio",
                required:true,
            },
            {
                id:2,
                name:"end_date",
                type:"date",
                placeholder:"Fecha Fin",
                errorMessage:"La fecha de vuelta no puede ser nula!",
                label:"Fecha Fin",
                required:true,
            }

        ]


        const jsonData = {
            "start_date": convertirFecha(values.start_date),
            "end_date":convertirFecha(values.end_date),
            "user_booked_id": Cookie.get("user_id"),
            "booked_hotel_id": id //hotel_id
        }

        const handleSubmit = async (e) => {
            e.preventDefault();

            insertBooking(jsonData)
        }

        const onChange = (e) => {
            setValues({...values, [e.target.name]: e.target.value})
        }

        console.log(values);





    return (
        <div className='contenedor-principal'>

            <Navbar />
            <div className='contenedor-hoteldetalle'>


                <div className='nombre'>
                    <h1>{hotel?.hotel_name}</h1><br/>
                </div>
                <div className='hoteles-detalle'>
                    <div className='descripcion'>
                    <p>Esta ubicado en {hotel?.hotel_address}</p>
                    <p>{hotel?.hotel_description}</p>
                        <div className='reserva'>
                            <form className= 'reservaForm' onSubmit={handleSubmit}>
                                <h1>Elija la fecha en la que le gustaria ir:</h1>
                                {inputs.map((input ) => (
                                    <FormInput key={input.id} {...input} value={values[input.name]} onChange={onChange}/>

                                ))}
                                <button className="reservar-button" onClick={chequear}>RESERVAR</button>
                            </form>
                        </div>
                        <ToastContainer />
                    </div>
                    </div>
                </div>
        </div>
    )
}

