import {useEffect, useState} from "react";
import HotelCard from '../componentes/HotelCard';

const [hotel, setHotel] = useState([]);

const getHotel = async () => {
    const response = await fetch(hotels);
    const resolve = await response.json();
    setHotel(resolve)
}
useEffect(() => {
    getHotel();
},[])
