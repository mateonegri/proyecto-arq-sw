import './App.css';
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import {Home} from './paginas/Home';
import {Login} from './paginas/Login';
import {navigation} from './Rutas/Navegacion';
import {HotelDetalle} from "./paginas/HotelDetalle";
import { SignIn } from './paginas/SignIn';
import {MisReservas} from './paginas/MisReservas';
import {ReservasAdmin} from './paginas/ReservasAdmin';
import Reserva from './paginas/Reserva'
import {useEffect} from "react";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'
import { EditarHotel } from './paginas/EditarHotel'
import { InsertarHotel } from './paginas/InsertarHotel'

function App() {


    return (
        <div className='App'>
            <BrowserRouter>
            <Routes>
                <Route path='' element={<Home />}/>
                <Route path='/home' element={<Home />}/>
                <Route path='/login' element={<Login />}/>
                <Route path='/signin' element={<SignIn />}/>
                <Route path='/home/hotel/:id' element={<HotelDetalle />} />
                <Route path='/hotel/edit/:id' element={<EditarHotel />} />
                <Route path='/misreservas' element={<MisReservas />} />
                <Route path='/hotel/insert' element={<InsertarHotel />} />

                <Route path='/admin/reservas' element={<ReservasAdmin />} />
                <Route path='/misreservas' element={<MisReservas />} />

            </Routes>
            </BrowserRouter>
        </div>
  );

}

export default App;
