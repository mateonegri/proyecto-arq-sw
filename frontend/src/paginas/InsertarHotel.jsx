import React, { useState } from "react";
import FormInput from "../componentes/FormInput.jsx"
import "../hojas-de-estilo/SignInStyle.css"
import Navbar from "../componentes/Navbar.jsx";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'
import { useNavigate } from "react-router-dom";
import {useParams} from "react-router-dom";
import Cookies from "universal-cookie";


const Cookie = new Cookies()

function goTo(path){
    setTimeout(() => {
        window.location = window.location.origin + path;
    }, 3000)
}
  
const notifyRegistered = () => {
    toast.success("Agregado!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}

const notifyError = () => {
    toast.error("Error al agregar!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}


const postUser = "http://localhost:8090/hotel"

async function updateHotel(jsonData) {

        const response = await fetch(postUser, {
            method: "POST",
            headers:{"content-type":"application/json"},
            body: JSON.stringify(jsonData)
        }).then(response => {
            if (response.status === 400 || response.status === 401 || response.status === 403) {
                console.log("Error al agregar hotel"); 

                notifyError();

                return response.json();

            } else {
                console.log("Hotel inserted"); 

                notifyRegistered();

                goTo("/");

                return response.json();
            }

        })

}



export const InsertarHotel = (hotel_id) => {

    const navigate = useNavigate();

    const {id} = useParams()

    const [values, setValues] = useState({
        name:"",
        description:"",
        rooms:"",
        address:"",
    })

    const inputs = [
        {
            id:1,
            name:"name",
            type:"text",
            placeholder:"Nombre",
            errorMessage: "El nombre no puede ser nulo!",
            label:"Nombre",
            required: true,
        },
        {
            id:2,
            name:"description",
            type:"text",
            placeholder:"Descripcion del hotel",
            errorMessage: "Ingrese una descripcion del hotel",
            label:"Descripcion del hotel",
            required: true,
        },
        {
            id:3,
            name:"rooms",
            type:"number",
            placeholder:"Habitaciones",
            errorMessage: "No puede dejar este campo vacio!",
            label:"Habitaciones",
            required: true,
        },
        {
            id:4,
            name:"address",
            type:"text",
            placeholder:"Direccion del hotel",
            errorMessage: "No puede dejar este campo vacio!",
            label:"Direccion del hotel",
            required: true,
        },
    ]

    console.log(id)

     const jsonData = {
        "id": Number(id),
        "hotel_name": values.name,
        "hotel_description": values.description,
        "hotel_rooms": Number(values.rooms),
        "hotel_address": values.address,
        "user_id": Number(Cookie.get("user_id"))
    } 

    // const userObj = { name, last_name, username, password, email, type}

    const handleSubmit = async (e) => {
        e.preventDefault();

        updateHotel(jsonData)

    }

    const onChange = (e) => {
        setValues({...values, [e.target.name]: e.target.value})
    }

    console.log(values);

    return (
        <>
        <Navbar />
        <div className="Page">
            <form className = "SignInForm" onSubmit={handleSubmit}>
                <h1>Agregue un hotel!</h1>
                {inputs.map((input) => (
                    <FormInput key={input.id} {...input} value={values[input.name]} onChange={onChange}/>
                ))}
                <button className="RegisterButton">Agregar!</button>
            </form>
        </div>
        <ToastContainer />
        </>
    )
        
};