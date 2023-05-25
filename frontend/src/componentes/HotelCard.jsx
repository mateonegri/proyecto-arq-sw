import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
 const HotelCard = ({ hotel_name, hotel_address, hotel_id, hotel_description, hotel_image }) => {
    const navigate = useNavigate();

    const selectHotel = () => {
        navigate(`/home/hotel/${hotel_id}`);
    };

    return (
        <div className='contenedor-hotel'>
            <section onClick={selectHotel}>
                <div className='contenedor-informacion-hotel'>
                    <div className='nombre-hotel'>
                        <h1>{hotel_name}</h1>
                    </div>
                    <div className='ubicacion-hotel'>
                        <p>{hotel_address}</p>
                    </div>
                    <div className='descripcion-hotel'>
                        <p>{hotel_description}</p>
                    </div>
                    {/*<div className='imagen-hotel'>*/}
                    {/*    <img src={require(hotel.hotel_image_url)}/>*/}
                    {/*</div>*/}
                </div>
            </section>
        </div>
    );
}
export default HotelCard;