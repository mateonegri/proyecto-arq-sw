import React from "react";
import useLocalState from 'use-local-storage';
import {Navigate} from "react-router-dom";


const RutasPrivadas = ({children}) => {
    const [jvt, setJvt] = useLocalState("", "jvt");
    return jvt ?  children : <Navigate to="/login" />;
}

export default RutasPrivadas;