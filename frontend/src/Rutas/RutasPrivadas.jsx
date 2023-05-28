import React from "react";
import Cookies from "universal-cookie";
const Cookie = new Cookies();

function goto(path){
    window.location = window.location.origin + path
}
const RutasPrivadas = () => {
    if (Cookie.get("user_id") > -1 ) {
        goto("/home/hotel/reserva/:id")
    } else {
        goto("/login")
    }

}

export default RutasPrivadas;