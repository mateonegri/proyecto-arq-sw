import { useState } from "react";
import "../hojas-de-estilo/FormInput.css"

const FormInput = (props) => {

    const [focused, setFocused] = useState(false);

    const {label, errorMessage, onChange, id, ...inputProps } = props;

    const handleFocus = (e) => {
        setFocused(true);
    }

    return (
        <div className="formInput">
            <label className="SignInLabel">{label}</label>
            <input className="SignInInput"
            {...inputProps} 
            onChange={onChange} 
            onBlur={handleFocus} 
            onFocus={() => inputProps.name === "confirmPassword" && setFocused(true)} 
            focused={focused.toString()}/>
            <span>{errorMessage}</span>
        </div>
    )
};

export default FormInput