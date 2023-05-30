import React, { useState } from "react";
import FormInput from "../componentes/FormInput.jsx"
import "../hojas-de-estilo/SignInStyle.css"
import Navbar from "../componentes/Navbar.jsx";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'

function goTo(path){
    setTimeout(() => {
        window.location = window.location.origin + path;
    }, 3000)
}
  
const notifyRegistered = () => {
    toast.success("Registrado!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}

const notifyError = () => {
    toast.error("Email o Nombre de usuario ya registrados!", {
        pauseOnHover: false,
        autoClose: 2000,
    })
}

const postUser = "http://localhost:8090/user"

async function insertUser(jsonData) {

        const response = await fetch(postUser, {
            method: "POST",
            headers:{"content-type":"application/json"},
            body: JSON.stringify(jsonData)
        }).then(response => {
            if (response.status === 400 || response.status === 401 || response.status === 403) {
                console.log("Error al insertar usuario"); 

                notifyError();

                return response.json();

            } else {
                console.log("User added"); 

                notifyRegistered();

                goTo("/login");

                return response.json();
            }

        })

}


export const SignIn = () => {

    const [values, setValues] = useState({
        name:"",
        last_name:"",
        email:"",
        username:"",
        password:"",
        confirmPassword:"",
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
            name:"last_name",
            type:"text",
            placeholder:"Apellido",
            errorMessage: "El apellido no puede ser nulo!",
            label:"Apellido",
            required: true,
        },
        {
            id:3,
            name:"email",
            type:"email",
            placeholder:"Email",
            errorMessage: "Debe ser una direccion de email valida!",
            label:"Email",
            required: true,
        },
        {
            id:4,
            name:"username",
            type:"text",
            placeholder:"Nombre de usuario",
            errorMessage: "El nombre de usuario debe tener entre 3-16 caracteres y no incluir caracteres especiales!",
            label:"Nombre de usuario",
            pattern:"^[A-Za-z0-9]{3,16}$",
            required: true,
        },
        {
            id:5,
            name:"password",
            type:"password",
            placeholder:"Contraseña",
            errorMessage: "La constraseña debe tener entre 8-20 caracteres, y debe tener 1 letra, 1 numero y un caracter especial!",
            label:"Contraseña",
            pattern:"^(?=.*[0-9])(?=.*[a-zA-Z])(?=.*[!@#$%^&*])[a-zA-Z0-9!@#$%^&*]{8,20}$",
            required: true,
        },
        {
            id:6,
            name:"confirmPassword",
            type:"password",
            placeholder:"Confirmar contraseña",
            errorMessage: "Las contraseñas no coinciden!",
            label:"Confirmar contraseña",
            pattern: values.password,
            required: true,
        }
    ]

     const jsonData = {
        "name": values.name,
        "last_name": values.last_name,
        "username": values.username,
        "password": values.password,
        "email": values.email,
        "type": false
    } 

    // const userObj = { name, last_name, username, password, email, type}

    const handleSubmit = async (e) => {
        e.preventDefault();

        insertUser(jsonData)

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
                <h1>Registrese!</h1>
                {inputs.map((input) => (
                    <FormInput key={input.id} {...input} value={values[input.name]} onChange={onChange}/>
                ))}
                <button className="RegisterButton">Registrarse!</button>
            </form>
        </div>
        <ToastContainer />
        </>
    )
        
};
