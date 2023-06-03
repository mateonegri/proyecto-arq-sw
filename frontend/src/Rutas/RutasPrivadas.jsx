import React from "react";
import Cookies from "universal-cookie";
const Cookie = new Cookies();

function goto(path){
    window.location = window.location.origin + path
}
const RutasPrivadas = () => {
    if (Cookie.get("user_type") === true ) {
        goto("/admin/reservas")

    } else {
        goto("/misreservas")
    }
}

export default RutasPrivadas;