import { lazy } from "react";
import { Home } from '../paginas/Home'

const Login = lazy(() => import('../paginas/Login'))
const HotelDetalle = lazy(() => import('../paginas/HotelDetalle'))
const SignIn = lazy(() => import('../paginas/SignIn'))
const Reserva = lazy(() => import('../paginas/Reserva'))
const ReservasAdmin = lazy(() => import('../paginas/ReservasAdmin'))
const MisReservas = lazy(() => import('../paginas/MisReservas'))
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
        path: "/home/hotel/reserva/:id",
        Element: Reserva,
    },
    {
        id: 5,
        path: "/signin",
        Element: SignIn,
    },
    {
        id: 6,
        path: "/admin/reservas",
        Element: ReservasAdmin,
    },
    {
        id: 7,
        path: "/misreservas",
        Element: MisReservas,
    }

];

export {Home};

