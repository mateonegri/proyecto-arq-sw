import React, { useEffect, useState } from "react";
import FormInput from "../componentes/FormInput.jsx"
import "../hojas-de-estilo/SignInStyle.css"
import Navbar from "../componentes/Navbar.jsx";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'
import { useNavigate } from "react-router-dom";
import {useParams} from "react-router-dom";
import Cookies from "universal-cookie";
import { MenuItem, Button, InputLabel, Select } from "@mui/material";

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

export const AgregarAmenitie = () => {

    const [amenities, setAmenities] = useState([]);
    const [selectedAmenity, setSelectedAmenity] = useState("");

    const {id} = useParams()

    const getAmenities = async () => {
        try {
          const response = await fetch(`http://localhost:8090/amenities`);
          const data = await response.json();
          if (data.amenities != "") {
            setAmenities(data.amenities);
          }
        } catch (error) {
            console.error('Error al obtener las amenities:', error);
        }
    }

    useEffect(() => {  
        // Obtener lista de amenities
        getAmenities();
      }, []);


    const handleAmenityChange = (event) => {
        setSelectedAmenity(event.target.value);
    };

    const handleSubmit = (event) => {
        event.preventDefault();
    

        console.log(Cookie.get("user_id"))
        if (Number(Cookie.get("user_id")) === 1){ 
        console.log(Cookie.get("adentro"))
        // Realizar la solicitud PUT al backend con los datos seleccionados
        const hotelId = id;
        const amenityId = selectedAmenity;

        fetch(`http://localhost:8090/hotel/${hotelId}/add-amenitie/${amenityId}`, {
            method: "PUT",
            headers:{"content-type":"application/json"}
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

    }


    return (
        <>
         <Navbar />
        <div className="Page">
            <form className = "SignInForm" onSubmit={handleSubmit}>
                <div>
                <h1>Agregar Amenitie!</h1>
                {amenities?.length ? (
                <Select
                    labelId="amenitie-select-label"
                    id="amenitie-select"
                    value={selectedAmenity}
                    onChange={handleAmenityChange}
                >
                {amenities.map((amenitie) => (
                    <MenuItem key={amenitie.id} value={amenitie.id}>
                        {amenitie.name}
                    </MenuItem>
                ))}
            </Select>
            ) : (
                <p>No hay amenities disponibles</p>
            )}
            </div>
                <button type="submit" className="RegisterButton" onClick={handleSubmit}>Actualizar!</button>
            </form>
        </div>
        <ToastContainer />
        </>
    )
}