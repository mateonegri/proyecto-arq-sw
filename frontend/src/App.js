import logo from './logo.svg';
import './App.css';
import {useEffect} from "react";
import { BrowserRouter, Routes, Route, useNavigate, Navigate } from 'react-router-dom'
import Navbar from './componentes/Navbar';
import {Home} from './paginas/Home';
<<<<<<< HEAD
import {Login} from './paginas/Login';
import {navigation} from './Rutas/Navegacion';



const App = () => {

    useEffect(() => {

    }, []);
=======
import {Login} from "./paginas/Login";
import {navigation} from "./Rutas/Navegacion";

function App() {
>>>>>>> 2067c45f64935e3d1db8ef3d7b1c3f4031d5bc19
    return (
        <div className="App">
            <BrowserRouter>
            <Routes>
                <Route path='' element={<Home />}/>
                <Route path='/home' element={<Home />}/>
                <Route path='/login' element={<Login />}/>
<<<<<<< HEAD
            </Routes>
=======
        </Routes>
>>>>>>> 2067c45f64935e3d1db8ef3d7b1c3f4031d5bc19
            </BrowserRouter>
        </div>
  );

}

export default App;
