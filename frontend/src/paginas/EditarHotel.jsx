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
    toast.success("Actualizado!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}

const notifyError = () => {
    toast.error("Error al actualizar!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}


const postUser = "http://localhost:8090/hotel/update"

async function updateHotel(jsonData) {

        const response = await fetch(postUser, {
            method: "PUT",
            headers:{"content-type":"application/json"},
            body: JSON.stringify(jsonData)
        }).then(response => {
            if (response.status === 400 || response.status === 401 || response.status === 403) {
                console.log("Error al actualizar hotel"); 

                notifyError();

                return response.json();

            } else {
                console.log("Hotel updated"); 

                notifyRegistered();

                goTo("/");

                return response.json();
            }

        })

}



export const EditarHotel = (hotel_id) => {

    const navigate = useNavigate();

    const {id} = useParams()

    let deleteHotelURL = `http://localhost:8090/hotel/delete/${id}/${Cookie.get("user_id")}`

    console.log(deleteHotelURL)

    function deleteHotel() {
         fetch(deleteHotelURL, { 
            method: 'DELETE' 
        })
        .then(res => res.json())
        .then(data => console.log(data))
        .catch(error => console.error(error));
    }

    const [values, setValues] = useState({
        name:"",
        description:"",
        rooms:"",
        address:"",
        ImageURL:"",    
    })

    const inputs = [
        {
            id:1,
            name:"name",
            type:"text",
            placeholder:"Nombre",
            label:"Nombre",
        },
        {
            id:2,
            name:"description",
            type:"text",
            placeholder:"Descripcion del hotel",
            label:"Descripcion del hotel",
        },
        {
            id:3,
            name:"rooms",
            type:"number",
            placeholder:"Habitaciones",
            label:"Email",
        },
        {
            id:4,
            name:"address",
            type:"text",
            placeholder:"Direccion del hotel",
            label:"Direccion del hotel",
        },
        {
            id:5,
            name:"ImageURL",
            type:"text",
            placeholder:"URL de la Imagen",
            label:"URL de la Imagen",
        },
    ]

    console.log(id)

     const jsonData = {
        "id": Number(id),
        "hotel_name": values.name,
        "hotel_description": values.description,
        "hotel_rooms": Number(values.rooms),
        "hotel_address": values.address,
        "hotel_image_url": values.ImageURL,
        "user_id": Number(Cookie.get("user_id"))
    } 

    // const userObj = { name, last_name, username, password, email, type}

    const handleSubmit = async (e) => {
        e.preventDefault();

        updateHotel(jsonData)

    }

    const handleDelete = async (e) => {
        e.preventDefault();

        deleteHotel();
    }

    const onChange = (e) => {
        setValues({...values, [e.target.name]: e.target.value})
    }

    const handleInsert = () => {
        navigate(`/hotel/insert`)
    }

    console.log(values);

    return (
        <>
        <Navbar />
        <div className="Page">
            <form className = "SignInForm" onSubmit={handleSubmit}>
                <h1>Actualizacion!</h1>
                {inputs.map((input) => (
                    <FormInput key={input.id} {...input} value={values[input.name]} onChange={onChange}/>
                ))}
                <button className="RegisterButton">Actualizar!</button>
            </form>
            <div className="AdminHotelOptions">
                <button className="DeleteButton" onClick={handleDelete}>Borrar!</button>
                <button className="AddHotelButton" onClick={handleInsert}>Agregar Hotel!</button>
            </div>
        </div>
        <ToastContainer />
        </>
    )
        
};
