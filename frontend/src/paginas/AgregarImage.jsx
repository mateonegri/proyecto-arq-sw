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
import { uploadFileToIPFS } from "../pinata";

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

export const AgregarImage = () => {

    const [fileURL, setFileURL] = useState(null);
    const [message, updateMessage] = useState('');

    const {id} = useParams()

    async function disableButton() {
        const listButton = document.getElementById("list-button")
        listButton.disabled = true
        listButton.style.backgroundColor = "grey";
        listButton.style.opacity = 0.3;
    }

    async function enableButton() {
        const listButton = document.getElementById("list-button")
        listButton.disabled = false
        listButton.style.backgroundColor = "#A500FF";
        listButton.style.opacity = 1;
    }

     //This function uploads the NFT image to IPFS
     async function OnChangeFile(e) {
        var file = e.target.files[0];
        //check for file extension
        try {
            //upload the file to IPFS
            disableButton();
            updateMessage("Subiendo imagen... no presione nada!")
            const response = await uploadFileToIPFS(file);
            if(response.success === true) {
                enableButton();
                updateMessage("")
                console.log("Uploaded image to Pinata: ", response.pinataURL)
                setFileURL(response.pinataURL);
            }
        }
        catch(e) {
            console.log("Error during file upload", e);
        }
    }

    const handleSubmit = (event) => {
        event.preventDefault();
    
        const jsonData = {
            "url": fileURL
        }

        console.log(Cookie.get("user_id"))
        if (Number(Cookie.get("user_id")) === 1){ 
        // Realizar la solicitud PUT al backend con los datos seleccionados
        const hotelId = id;

        fetch(`http://localhost:8090/hotel/${id}/add-image`, {
            method: "POST",
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

    }

    return (
        <>
         <Navbar />
        <div className="Page">
            <form>
                <input type={"file"} onChange={OnChangeFile}></input>
                <button type="submit" className="RegisterButton" id="list-button" onClick={handleSubmit}>Actualizar!</button>
            </form>
        </div>
        <ToastContainer />
        </>
    )
}