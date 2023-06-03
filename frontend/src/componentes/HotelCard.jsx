import React, { useState } from "react";
import { useNavigate } from "react-router-dom";


import Card from '@mui/material/Card';
import CardActions from '@mui/material/CardActions';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Button from '@mui/material/Button';
import Typography from '@mui/material/Typography';
import Cookies from "universal-cookie";

const Cookie = new Cookies();

 const HotelCard = ({ hotel_name, hotel_address, hotel_id, hotel_description, hotel_image }) => {
    const navigate = useNavigate();
    const [isAdmin, setIsAdmin] = useState(false);

    const selectHotel = () => {
        navigate(`/home/hotel/${hotel_id}`);
    };

    const editHotel = () => {
        navigate(`/hotel/edit/${hotel_id}`)
    }

    const renderButton = (
        <>
        <Button className="masinfo-boton" onClick={editHotel}>Editar</Button>
        </>
    )

    const renderNothing = (
        <>
        </>
    )


    console.log(hotel_image)
        return (
        <Card sx={{width:'92%'}} className="cartaHotel">
          <CardMedia
          component="img"
          alt={hotel_name}
          height="140"
          image = {hotel_image}
          
          />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            {hotel_name}
          </Typography>
          <Typography variant="body2" color="text.secondary">
          <p>{hotel_address}</p>
          <p>{hotel_description}</p>
          </Typography>
        
        </CardContent>
        <CardActions>
            {Cookie.get("user_type") === "true" ? renderButton : renderNothing} 
          <Button className="masinfo-boton" onClick={selectHotel}>Mas Informacion</Button>
        </CardActions>
      </Card>
          );
        }

export default HotelCard;
