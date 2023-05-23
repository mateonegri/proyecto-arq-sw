import React, {useState} from "react";
import {useNavigate} from "react-router-dom";
import '../hojas-de-estilo/Hotel.css'

const HotelCard = ({ HotelName, HotelDescription }) => {
    const navigate = useNavigate();
    return (
        <div className='hotel'>
            <div>
                <h1>{HotelCard.HotelName}</h1>
            </div>
            <div>
                <p>{HotelCard.HotelDescription}</p>
            </div>
        </div>

    )
}

export default HotelCard;