import React from "react";
import Cookies from "universal-cookie";
const Cookie = new Cookies();
function goto(path){
    window.location = window.location.origin + path
}
export const MisReservas=() => {

   if (Cookie.get("user_type") === "true"){
        goto("/admin/reservas")
    } else if (Cookie.get("user_id") === "-1"){
        goto("/login")
    }
    return (
        <>

        </>
    )

}