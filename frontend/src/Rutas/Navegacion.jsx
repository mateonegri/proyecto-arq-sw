import { lazy } from "react";

const Home = lazy(() => import('../paginas/Home'));
const Login = lazy(() => import('../paginas/Login'))
export const navigation = [
    {
        id: 1,
        path: "/home",
        Element: Home,
    },
    {
        id: 2,
        path: "/login",
        Element: Login,
    }

];

