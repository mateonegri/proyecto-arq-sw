import logo from './logo.svg';
import './App.css';
import { BrowserRouter } from 'react-router-dom'
import Navbar from './componentes/Navbar';
import {Home} from './paginas/Home';
import {Login} from "./paginas/Login";
function App() {

        let component
        switch (window.location.pathname) {
            case '/':
            component = <Home />
            break
            case '/home':
            component = <Home />
            break
            case '/login':
            component = <Login />
                break
    }
    return (
        <div className="App">
        <Navbar />
            {component}
        </div>
  );

}

export default App;
