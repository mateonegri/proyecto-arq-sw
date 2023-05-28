import './App.css';
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import {Home} from './paginas/Home';
import {Login} from './paginas/Login';
import {navigation} from './Rutas/Navegacion';
import {HotelDetalle} from "./paginas/HotelDetalle";
import { SignIn } from './paginas/SignIn';

function App() {

    return (
        <div>
            <BrowserRouter>
            <Routes>
                <Route path='' element={<Home />}/>
                <Route path='/home' element={<Home />}/>
                <Route path='/login' element={<Login />}/>
                <Route path='/home/hotel/:id' element={<HotelDetalle />}/>
                <Route path='/signin' element={<SignIn />}/>
            </Routes>
            </BrowserRouter>
        </div>
  );

}

export default App;
