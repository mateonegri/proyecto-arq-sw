import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";

const hotels = "http://localhost:3000/hotel.json"

const Hotel = () => {

const [hotel, setHotel] = useState([]);
    const navigate = useNavigate();
const getHotel = async () => {
    const response = await fetch(hotels);
    const hotel = await response.json();
    // setHotel(hotel)
    console.log(hotel)
    return hotel;
}
useEffect(() => {
    getHotel().then ((hotel) => setHotel(hotel));
},[])

const id = hotel.id;
const selectHotel = () => {
    navigate(`/home/hotel/${hotel.id}`);
}

return (
    <div className='contenedor-principal'>
            <h1>Tus vacaciones estan esperandote...</h1>
            {
                hotel.length ? hotel.map((hotel) =>
                    <div className='contenedor-hotel'>
                        <section onClick={selectHotel}>
                        <div className='contenedor-informacion-hotel'>
                            <div className='nombre-hotel'>
                            <h1>{hotel.hotel_name}</h1>
                            </div>
                            <div className='ubicacion-hotel'>
                                <p>{hotel.hotel_address}</p>
                            </div>
                        <div className='descripcion-hotel'>
                            <p>{hotel.hotel_description}</p>
                        </div>
                            {/*<div className='imagen-hotel'>*/}
                            {/*    <img src={require(hotel.hotel_image_url)}/>*/}
                            {/*</div>*/}

                        </div>
                        </section>
                    </div>):null
            }

    </div>
)
}

export default Hotel