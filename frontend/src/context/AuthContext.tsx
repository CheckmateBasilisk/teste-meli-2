import { createContext, Dispatch, SetStateAction } from "react";
import React, { useState } from 'react';

// saves data from currently authenticated customer

interface customerInfo {
    id: string,
    login: string,
    name: string,
    email: string,
}

interface AuthContextProps {
    // fields from customer + "setters" ? 

    // this isnt secure (man in the middle)
    JWTToken: string,
    setJWTToken: Dispatch<SetStateAction<string>>,
    customer: customerInfo,
    setCustomer: Dispatch<SetStateAction<customerInfo>>,
}

// these are values just for startup, not the actual innital values (theyre set at AuthContextProvider)
const AuthContextInitialState = {
    // fields from customer + "setters" ? 
    JWTToken: "",
    setJWTToken: () => {},
    customer: {
        id: "",
        login: "",
        name: "",
        email: "",
    },
    setCustomer: () => {},
}

// setting up a context
export const AuthContext = createContext <AuthContextProps> (AuthContextInitialState)

export const AuthContextProvider = ({children}:any) => {
    //setting up context's actual innitial values
    const [JWTToken, setJWTToken] = useState("lalilulelo");
    const [customer, setCustomer] = useState({id: "",login: "",name: "",email: ""},);
    
    return (
        <AuthContext.Provider value={{JWTToken, setJWTToken, customer, setCustomer}} >
            {children}
        </AuthContext.Provider>
    )
}