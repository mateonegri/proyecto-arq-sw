import React from 'react';
import Logo from '../imagenes/TU HOTEL.png';
import {Link} from "react-router-dom"

function Navbar (){


        return <nav className='nav'>
        <a href='/' className='Titulo-de-la-web'>TU HOTEL</a>


        <ul>
            <li className='active'>
                <a href='/home'>Home</a>
            </li>
            <li>
                <a href='/login'>Log In</a>
            </li>
            <li>
                <a href='/misreservas'>Mis Reservas</a>
            </li>
        </ul>
    </nav>
}
export default Navbar;