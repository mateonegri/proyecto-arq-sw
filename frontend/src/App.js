import logo from './logo.svg';
import './App.css';
import {useEffect} from "react";
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import Navbar from './componentes/Navbar';
import {Home} from './paginas/Home';
import {Login} from './paginas/Login';
import {navigation} from './Rutas/Navegacion';



const App = () => {

    useEffect(() => {

    }, []);
    return (
        <div className="App">
            <BrowserRouter>
            <Routes>
                <Route path='' element={<Home />}/>
                <Route path='/home' element={<Home />}/>
                <Route path='/login' element={<Login />}/>
            </Routes>
            </BrowserRouter>
        </div>
  );

}

export default App;
