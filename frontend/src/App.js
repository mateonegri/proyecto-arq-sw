import './App.css';
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import {Home} from './paginas/Home';
import {Login} from './paginas/Login';
import {navigation} from './Rutas/Navegacion';
import {HotelDetalle} from "./paginas/HotelDetalle";
import { SignIn } from './paginas/SignIn';
import {MisReservas} from './paginas/MisReservas';
import {ReservasAdmin} from './paginas/ReservasAdmin';
import RutasPrivadas from "./Rutas/RutasPrivadas";
import Reserva from './paginas/Reserva'
import {useEffect} from "react";
import { ToastContainer, toast } from "react-toastify";
import 'react-toastify/dist/ReactToastify.css'

function App() {


    return (
        <div>
            <BrowserRouter>
            <Routes>
                <Route path='' element={<Home />}/>
                <Route path='/home' element={<Home />}/>
                <Route path='/login' element={<Login />}/>
                <Route path='/signin' element={<SignIn />}/>
                <Route path='/home/hotel/:id' element={<HotelDetalle />} />

                <Route element={<RutasPrivadas/>}>
                    <Route path='/misreservas' element={<MisReservas />} />
                </Route>
            </Routes>
            </BrowserRouter>
        </div>
  );

}

export default App;
