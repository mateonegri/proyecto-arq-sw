import { lazy } from "react";
const Home = lazy(() => import('../Paginas/Home'));


export const navigation = [
    {
        id: 1,
        path: "/home",
        Element: Home,
    },

];
