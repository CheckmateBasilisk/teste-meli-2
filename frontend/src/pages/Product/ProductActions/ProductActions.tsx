import Typography from '@mui/material/Typography';
import { NumberField } from '@base-ui-components/react/number-field';
import RemoveIcon from '@mui/icons-material/Remove';
import AddIcon from '@mui/icons-material/Add';
import styles from './style.module.scss';
import React, { useContext, useEffect, useState } from 'react';
import Button from '@mui/material/Button';
import { CartContext } from '../../../context/CartContext';


// TODO: move this to a single location...
interface params {
  product : {
    id: string,
    barcode: string,
    name: string,
    rating: number,
    price: number,
    stock: number,
    descr: string,
    image: string,
    alt: string, 
  }
}

export default function ProductActions(p : params) {
  const id = React.useId();

  const [quantity, setQuantity] = useState(1);
  const { cartEntries, setCartEntries } = useContext(CartContext);

  //const [cartEntries, setCartEntries] = useState( CartContextInitialState.cartEntries );
  // console.log("currentEntry: ",currentEntry)
  
  
  console.log("cartEntries: ",cartEntries)

  return (
    <div style={{ padding: '1rem', border: '1px solid #ccc', borderRadius: '0.5rem' }}>
      {/* TODO: i should move this to a component eventually... */}
      <NumberField.Root id={id} defaultValue={1} min={1} value={quantity} onValueChange={(value) => {setQuantity(value as number)}} className={styles.Field}>
      <NumberField.Group className={styles.Group}>
        <NumberField.Decrement className={styles.Decrement}>
          <RemoveIcon />
        </NumberField.Decrement>
        <NumberField.Input className={styles.Input} />
        <NumberField.Increment className={styles.Increment}>
          <AddIcon />
        </NumberField.Increment>
      </NumberField.Group>
      </NumberField.Root>

      <Typography variant="body1" gutterBottom>
        Total: {Intl.NumberFormat('pt-BR', { style: 'currency', currency: 'BRL' }).format(p.product.price*quantity)}
      </Typography>

      {/* botao adicionar ao carrinho */}
      {/* TODO: IMPLEMENTAR ADICIONAR AO CARRINHO, preciso mexer na API pra isso... */}
      <Button variant="contained" href="#contained-buttons" onClick={() => {
        let currentEntry = cartEntries.find(e => e.product.id == p.product.id);

        if (!currentEntry) {
          const newEntry = { // FIXME: Sub for post api call 
            id: "lalilulelo",
            amount: quantity,
            product: p.product
          }
          return setCartEntries([...cartEntries, newEntry]) //FIXME: cartEntries is initialized with an awful starting value. Upon adding new Entries, that value stays on the array...
        }

        currentEntry.amount = quantity 
        setCartEntries(cartEntries);
        //given cartEntries, find current entry. Add quantity to it
      }}>
        Add to Cart
      </Button>
    </div>
);
}