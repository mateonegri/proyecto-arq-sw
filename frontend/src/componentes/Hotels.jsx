import React, {useEffect, useState} from "react";


const hotels = "http://localhost:3000/hotel.json"

const Hotel = () => {

const [hotel, setHotel] = useState([]);

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

return (
    <div className='contenedor-principal'>
            <h1>Tus vacaciones te estan esperando</h1>
            {
                hotel.length ? hotel.map((hotel) =>
                    <div className='contenedor-hotel'>
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
                    </div>):null
            }

    </div>
)
}

export default Hotel