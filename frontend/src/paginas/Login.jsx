import React, { useState } from "react";
import ReactDOM from "react-dom";
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function login(username, password) {
  return await fetch('http://localhost:8090/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
	body: JSON.stringify({"username": username, "password":password})
  })
    .then(response => {
      if (response.status == 400 || response.status == 401)
      {
        return {"user_id": -1}
      }
      return response.json()
    })
    .then(response => {
      Cookie.set("user_id", response.user_id, {path: '/'})
      Cookie.set("username", username, {path: '/login'})
    })
 }

function goto(path){
  window.location = window.location.origin + path
}

export function Login() {
<<<<<<< HEAD
    return (
        <div className='Login'>
             <div className='contenedor-principal'>
                 <p>¡Hola! Registrate</p>
                 <Login />
             </div>
=======
  // React States
  const [errorMessages, setErrorMessages] = useState({});
  const [isSubmitted, setIsSubmitted] = useState(false);

  const error = "Contraseña o Usuario invalido";

  const handleSubmit = (event) => {
    //Prevent page reload
    event.preventDefault();

    var { uname, pass } = document.forms[0];

    // Find user login info
    const userData = login(uname.value, pass.value).then(data => {
      if (Cookie.get("user_id") > -1) {
        goto("/")
      }
      else if (Cookie.get("user_id") == -1) {
        setErrorMessages({name: "default", message: error})
      }
    })
  };


  // Generate JSX code for error message
  const renderErrorMessage = (name) =>
    name === errorMessages.name && (
      <div className="error">{errorMessages.message}</div>
    );

  // JSX code for login form
  const renderForm = (
    <div className="form">
      <form onSubmit={handleSubmit}>
        <div className="input-container">
          <label>Usuario </label>
          <input type="text" name="uname" placeholder="Usuario" required />
>>>>>>> 2067c45f64935e3d1db8ef3d7b1c3f4031d5bc19
        </div>
        <div className="input-container">
          <label>Contraseña</label>
          <input type="password" name="pass" placeholder="Contraseña" required />
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

    <div className="app">
      <div className="login-form">
        <div className="title">BIENVENIDOS</div>

        {isSubmitted || Cookie.get("user_id") > -1 ? Cookie.get("username") : renderForm}
      </div>

      <button onClick={logout}>Log Out</button>
    </div>
  );
}



