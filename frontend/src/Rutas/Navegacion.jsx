import { lazy } from "react";
import {Home} from '../paginas/Home'
const Home1 = lazy(() => import('../paginas/Home'));
const Login = lazy(() => import('../paginas/Login'))
export const navigation = [
    {
        id: 1,
        path: "/",
        Element: Home1,
    },
    {
        id: 2,
        path: "/login",
        Element: Login,
    }

];

export {Home};

