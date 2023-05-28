import React, { useState } from "react";
import FormInput from "../componentes/FormInput.jsx"
import "../hojas-de-estilo/SignInStyle.css"
import Navbar from "../componentes/Navbar.jsx";

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

    const handleSubmit = (e) => {
        e.preventDefault();
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
        </>
    )
        
};
