import React, {useEffect, useState} from "react";
import {useNavigate} from "react-router-dom";
import HotelCard from "./HotelCard";
import Cookies from "universal-cookie";

const Cookie = new Cookies();

const hotels = "http://localhost:8090/hotel"

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

console.log(Cookie.get("user_id"))

return (
    <div className='contenedor-principal'>
            <h1>Tus vacaciones estan esperandote...</h1>
            {
                hotel?.length ? hotel.map((hotel) => <HotelCard key={hotel.id} hotel_name={hotel.hotel_name} hotel_address={hotel.hotel_address} hotel_description={hotel.hotel_description} hotel_id={hotel.id} hotel_image={hotel.hotel_image_url}/>): <p>No hay hoteles</p>
            }
    </div>
)
}

export default Hotel