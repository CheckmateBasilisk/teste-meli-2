import { createContext, Dispatch, SetStateAction, useContext } from "react";
import { useState } from 'react';
import { AuthContext } from "./AuthContext";
import { api } from "../config/api";

// saves data from customer's shopping cart

export interface cartEntry {
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
    createCart: (productId: string, amount: number) => Promise<string>,
    isCartCreated: boolean,
    updateCart: (productId: string, amount: number, cartEntryId: string) => Promise<boolean>
    deleteCartEntry: (cartEntryId: string) => Promise<boolean>
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
    createCart: async () => "",
    updateCart: async () => false,
    isCartCreated: false,
    deleteCartEntry: async () => false,
}

export const CartContext = createContext <CartContextProps> (CartContextInitialState)

export const CartContextProvider = ({children}:any) => {
    const [cartEntries, setCartEntries] = useState<cartEntry[]>( [] );
    const [isCartCreated, setIsCartCreated] = useState(false);
    const { customer } = useContext(AuthContext);

    const createCart = async (productId: string, amount: number) => {
        try {
            const data = await api.post("/v1/cart", JSON.stringify({
                product_id: productId,
                customer_id: customer.id,
                amount: amount,
            }))
            
            const cartId = data?.id

            let currentEntryIndex = cartEntries.findIndex(e => e.product.id == productId);

            cartEntries[currentEntryIndex] = {
              ...cartEntries[currentEntryIndex],
              id: cartId
            }

            setCartEntries([...cartEntries])

            setIsCartCreated(true)
            return cartId
        } catch (error) {
            console.error(error)
            setIsCartCreated(false)
            return ""
        }
    }

    const updateCart = async (productId: string, amount: number, cartEntryId: string) => {
        try {
            await api.put(`/v1/cart/${cartEntryId}`, JSON.stringify({
                product_id: productId,
                customer_id: customer.id,
                amount: amount,
            }))
            
            return true
        } catch (error) {
            console.error(error)
            return false
        }
    }

    const deleteCartEntry = async (cartEntryId: string) => {
        try {
            await api.delete(`/v1/cart/${cartEntryId}`)
            return true
        } catch (error) {
            console.error(error)
            return false
        }
    }

    return (
        <CartContext.Provider value={{cartEntries, setCartEntries, createCart, isCartCreated, updateCart, deleteCartEntry}} >
            {children}
        </CartContext.Provider>
    )
}  

