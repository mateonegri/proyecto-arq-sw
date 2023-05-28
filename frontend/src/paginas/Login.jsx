import React, { useState } from "react";
import ReactDOM from "react-dom";
import Cookies from "universal-cookie";
import Navbar from "../componentes/Navbar.jsx";
import "../hojas-de-estilo/Login.css"

const Cookie = new Cookies();

async function login(username, password) {
<<<<<<< HEAD
    return await fetch('http://localhost:8090/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({"username": username, "password":password})
    })
        .then(response => {
            if (response.status === 400 || response.status === 401)
            {
                return {"user_id": -1}
            }
            return response.json()
        })
        .then(response => {
            Cookie.set("user_id", response.user_id, {path: '/'})
            Cookie.set("username", username, {path: '/login'})
        })
=======
  return await fetch('http://localhost:8090/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({"username": username, "password":password})
  })
      .then(response => {
        if (response.status === 400 || response.status === 401)
        {
          return {"user_id": -1}
        }
        return response.json()
      })
      .then(response => {
        Cookie.set("user_id", response.user_id, {path: '/'})
        Cookie.set("username", username, {path: '/login'})
      })
>>>>>>> 4faf5513968713b4aafb1e3bad2c30a2287b6ed4
}

function goto(path){
    window.location = window.location.origin + path
}

export function Login() {
    // React States
    const [errorMessages, setErrorMessages] = useState({});
    const [isSubmitted, setIsSubmitted] = useState(false);

    const error = "Contraseña o Usuario invalido";

    const handleSubmit = (event) => {
        //Prevent page reload
        event.preventDefault();

        var { uname, pass } = document.forms[0];

<<<<<<< HEAD
        // Find user login info
        const userData = login(uname.value, pass.value).then(data => {
            if (Cookie.get("user_id") > -1) {
                goto("/")
            }
            else if (Cookie.get("user_id") === -1) {
                setErrorMessages({name: "default", message: error})
            }
        })
    };
=======
    // Find user login info
    const userData = login(uname.value, pass.value).then(data => {
      if (Cookie.get("user_id") > -1) {
        goto("/")
      }
      else if (Cookie.get("user_id") === -1) {
        setErrorMessages({name: "default", message: error})
      }
    })
  };
>>>>>>> 4faf5513968713b4aafb1e3bad2c30a2287b6ed4


    // Generate JSX code for error message
    const renderErrorMessage = (name) =>
        name === errorMessages.name && (
            <div className="error">{errorMessages.message}</div>
        );

<<<<<<< HEAD
    // JSX code for login form
    const renderForm = (
        <div className="form">
            <form onSubmit={handleSubmit}>
                <div className="input-container">
                    <label>Usuario </label>
                    <input type="text" name="uname" required />
                </div>
                <div className="input-container">
                    <label>Contraseña</label>
                    <input type="password" name="pass" required />
                </div>

                {renderErrorMessage("default")}
                <div className="button-container">
                    <input type="submit" />
                </div>
            </form>
        </div>

    );

    function logout(){
        Cookie.set("user_id", -1, {path: "/"})
        document.location.reload()
    }

    return (
        <>
            <Navbar />
            <div className="app">
                <div className="login-form">
                    <div className="title">BIENVENIDOS</div>
                    {isSubmitted || Cookie.get("user_id") > -1 ? Cookie.get("username") : renderForm}
                </div>
            </div>
            <div className="logout-button">
                <button  onClick={logout}>Log Out</button>
            </div>
        </>
    );
=======
  // JSX code for login form
  const renderForm = (
      <div className="form">
        <form onSubmit={handleSubmit}>
          <div className="input-container">
            <label>Usuario </label>
            <input type="text" name="uname" required />
          </div>
          <div className="input-container">
            <label>Contraseña</label>
            <input type="password" name="pass" required />
          </div>

          {renderErrorMessage("default")}
          <div className="button-container">
            <input type="submit" />
          </div>
        </form>
      </div>

  );

  function logout(){
    Cookie.set("user_id", -1, {path: "/"})
    document.location.reload()
  }

  return (
    <> 
    <Navbar /> 
    <div className="app"> 
        <div className="login-form">
          <div className="title">BIENVENIDOS</div>
          {isSubmitted || Cookie.get("user_id") > -1 ? Cookie.get("username") : renderForm}
        </div>
    </div>
    <div className="logout-button">
          <button  onClick={logout}>Log Out</button>
    </div>
    </>
  );
>>>>>>> 4faf5513968713b4aafb1e3bad2c30a2287b6ed4
}

