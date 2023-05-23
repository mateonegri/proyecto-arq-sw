import React, {useEffect, useState} from "react";
import HotelCard from '../componentes/HotelCard';

const hotels = "https://localhost:8090/hotels"

const Hotel = () => {

const [hotel, setHotel] = useState([]);

const getHotel = async () => {
    const response = await fetch(hotels);
    const resolve = await response.json();
    setHotel(resolve)
}
useEffect(() => {
    getHotel();
},[])

return (
    <div>
    <h3>Hoteles...</h3>
    <div className='hotelCard'>

            {
                hotel.length ? hotel.map((hotel) => <HotelCard HotelName={hotel.HotelName} HotelDescription={hotel.HotelDescription} />): null
            }

    </div>
    </div>
)
}

export default Hotel