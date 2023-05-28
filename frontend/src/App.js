import './App.css';
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import {Home} from './paginas/Home';
import {Login} from './paginas/Login';
import {navigation} from './Rutas/Navegacion';
import {HotelDetalle} from "./paginas/HotelDetalle";
import RutasPrivadas from "./Rutas/RutasPrivadas";
import {useEffect} from "react";

function App() {


    return (
        <div className="App">
            <BrowserRouter>
            <Routes>
                <Route path='' element={<Home />}/>
                <Route path='/home' element={<Home />}/>
                <Route path='/login' element={<Login />}/>
                <Route path='/home/hotel/:id' element={
                    <RutasPrivadas>
                    <HotelDetalle />
                    </RutasPrivadas>
                    }
                    />
            </Routes>
            </BrowserRouter>
        </div>
  );

}

export default App;
