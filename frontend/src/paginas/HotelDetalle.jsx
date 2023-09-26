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
import { CardContent, CardMedia } from '@mui/material';
 import Card from '@mui/material/Card';
// import CardActions from '@mui/material/CardActions';
 //import CardContent from '@mui/material/CardContent';
 //import CardMedia from '@mui/material/CardMedia';
 //import Button from '@mui/material/Button';
 import Typography from '@mui/material/Typography';

const Cookie = new Cookies();

function getFormattedDate(target) {
    let year = target.getFullYear()
    let month = target.getMonth()
    let day = target.getDate()

    return`${year}-${month<9?'0':''}${month+1}-${day}`
}


function convertirFecha(fecha) {
    let fechaString = fecha.toString()
    

    let year = fechaString.substring(0,4)
    
    let month = fechaString.substring(5,7)
    
    let day = fechaString.substring(8,10)

    let yearPlusMonth = year.concat("",month)
    let fechaStringFinal = yearPlusMonth.concat("",day)

    

    var fechaEntero = Number(fechaStringFinal)

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

const notifyNotLoggedIn = () => {
    toast.error("Recuerda que antes debes loggearte!", {
        pauseOnHover: false,
        autoClose: 3000,
    })
}

const notifyBadDate = () => {
    toast.error("No puedes seleccionar una fecha devuelta anterior a la de inicio!", {
        pauseOnHover: false,
        autoClose: 3000,
    })
}

const notifyBadDate2 = () => {
    toast.error("No puedes seleccionar una fecha pasada!", {
        pauseOnHover: false,
        autoClose: 3000,
    })
}

const notifyBadDate3 = () => {
    toast.error("Debes seleccionar minimo una noche de estadia!", {
        pauseOnHover: false,
        autoClose: 3000,
    })
}


function goto(path) {
    setTimeout(()=> {
        window.location = window.location.origin + path
    }, 3000)
}

export const HotelDetalle = ( hotel_id ) => {
    const [hotel, setHotel] = useState()
    const {id} = useParams()
    const navigate = useNavigate();

    const infoHotel = `http://localhost:8090/hotel/${id}`


    const postBooking = "http://localhost:8090/booking"

    async function insertBooking(jsonData) {

        const response = await fetch(postBooking, {
            method: "POST",
            headers: {"content-type": "application/json"},
            body: JSON.stringify(jsonData)
        }).then(response => {
            if (response.status === 400 || response.status === 401 || response.status === 403) {
                console.log("Error al reservar");

                if (Cookie.get("user_id") === "-1") {
                    notifyNotLoggedIn()
                    goto("/login")
                } else {
                    notifyError();
                }

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
        "start_date": "",
        "end_date": "",
    })




        const inputs = [
            {
                id:1,
                name:"start_date",
                type:"date",
                placeholder:"Fecha Inicio",
                errorMessage:"La fecha de inicio no puede ser nula!",
                label:"Fecha Inicio",
                min: new Date().toISOString().split("T")[0] ,
                required:true,
            },
            {
                id:2,
                name:"end_date",
                type:"date",
                placeholder:"Fecha Fin",
                errorMessage:"La fecha de fin no puede ser nula!",
                label:"Fecha Fin",
                disabled: values.start_date === "" ? true: false,
                min : values.start_date ? new Date(values.start_date).toISOString().split("T")[0]: "" ,
                required: true,
            }

        ]


        const jsonData = {
            "start_date": convertirFecha(values.start_date),
            "end_date": convertirFecha(values.end_date),
            "user_booked_id": Number(Cookie.get("user_id")),
            "booked_hotel_id": Number(id) //hotel_id
        }

        console.log(Cookie.get("user_id"))
        console.log(jsonData)


        const handleSubmit = async (e) => {
            e.preventDefault();

            if (chequearfecha() !== 0) {
                insertBooking(jsonData)
            }






        }

        const onChange = (e) => {
            setValues({...values, [e.target.name]: e.target.value})
        }

        console.log(values);
    let hoy = (getFormattedDate(new Date))
  
    function chequearfecha() {
        if (values.end_date < values.start_date) {
            notifyBadDate();
            return 0
         // } else if ( values.start_date < (getFormattedDate(new Date)) ) {
         //     notifyBadDate2();
         //      return 0
        } else if (values.start_date == values.end_date) {
            notifyBadDate3();
            return 0
        }
    }

    return (
        <div className='contenedor-principal'>

            <Navbar />
            <Card sx={{width:'92%'}} className="hotelDetalle">
                <CardMedia
                    component="img"
                    alt={hotel?.hotel_name}
                    height="560"
                    image = {hotel?.hotel_image_url}
                />

                <CardContent>
                    <Typography gutterBottom variant="h5" component="div">
                        {hotel?.hotel_name}
                    </Typography>
                    <Typography variant="body2" color="text.secondary">
                        <p> Esta Ubicado en: {hotel?.hotel_address}</p>
                        <p>{hotel?.hotel_description}</p>

                        <div className='reserva'>
                            <form className= 'reservaForm'  >
                                <h1>Cuando nos quiere visitar?</h1>
                                {inputs.map((input ) => (
                                    <FormInput key={input.id} {...input} value={values[input.name]}  onChange={onChange}/>

                                ))}
                                <button className="reservar-button" onClick={handleSubmit} >RESERVAR</button>
                            </form>
                        </div>
                    </Typography>
                </CardContent>
            </Card>
            <ToastContainer/>
        </div>

    )
}

