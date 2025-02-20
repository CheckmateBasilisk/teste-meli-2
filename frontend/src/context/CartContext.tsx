import { createContext, Dispatch, SetStateAction } from "react";
import { useState } from 'react';

// saves data from customer's shopping cart

interface cartEntry {
    id: string,
    amount: number,
    product: {
        id: string,
        barcode: string,
        name: string,
        price: number,
        stock: number,
        rating: number,
        descr: string,
    },
}

interface CartContextProps {
    cartEntries: cartEntry[],
    setCartEntries: Dispatch<SetStateAction<cartEntry[]>>,
}


const CartContextInitialState = {
    cartEntries: [{
        id: "",
        amount: 0,
        product: {
            id: "",
            barcode: "",
            name: "",
            price: 0,
            stock: 0,
            rating: 0,
            descr: "",
        },
    }],
    setCartEntries: ()=>{},
}

export const CartContext = createContext <CartContextProps> (CartContextInitialState)

export const CartContextProvider = ({children}:any) => {
    const [cartEntries, setCartEntries] = useState( CartContextInitialState.cartEntries );

    return (
        <CartContext.Provider value={{cartEntries, setCartEntries}} >
            {children}
        </CartContext.Provider>
    )
}  

