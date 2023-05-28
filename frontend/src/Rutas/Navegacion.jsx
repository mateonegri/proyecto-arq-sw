import { lazy } from "react";
import { Home } from '../paginas/Home'

const Login = lazy(() => import('../paginas/Login'))
const HotelDetalle = lazy(() => import('../paginas/HotelDetalle'))
const SignIn = lazy(() => import('../paginas/SignIn'))
export const navigation = [
    {
        id: 1,
        path: "",
        Element: Home,
    },
    {
        id: 2,
        path: "/login",
        Element: Login,
    },
    {
        id: 3,
        path: "/home/hotel/:id",
        Element: HotelDetalle,
    },
    {
        id: 4,
        path: "/signin",
        Element: SignIn, 
    }

];

export {Home};

