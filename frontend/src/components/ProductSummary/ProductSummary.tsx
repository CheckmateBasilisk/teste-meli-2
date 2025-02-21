import Typography from '@mui/material/Typography';
import Grid2 from '@mui/material/Grid2'
import { NumberField } from '@base-ui-components/react/number-field';
import React, { useContext, useEffect, useState } from 'react';
import styles from './style.module.scss';
import RemoveIcon from '@mui/icons-material/Remove';
import AddIcon from '@mui/icons-material/Add';
import DeleteIcon from '@mui/icons-material/Delete';
import { FormatBRL } from '../../utils';
import { CartContext, cartEntry } from '../../context/CartContext';

interface params {
  cartEntry: cartEntry
  image: string
  alt: string
}

/*

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


*/


export default function ProductSummary(p : params) {
  const id = React.useId();
  const { cartEntries, setCartEntries, updateCart, deleteCartEntry } = useContext(CartContext);

  const [quantity, setQuantity] = useState(0);
  const [disableButton, setDisableButton] = useState(false);

  useEffect(() => {
    setQuantity(p?.cartEntry?.amount)
  }, [p?.cartEntry?.amount])

  return (

    // FIXME: THE SPACING ISNT RIGHT. THE IMGS KEEP ON CLIPPING
      <div className='product-summary'>
      <Grid2 container size={12}>
        <Grid2 size={3}>
        <div className='summary-img' >
          <img src={p.image} alt={p.alt} height="auto" width="100%"/>
        </div>
        </Grid2>
        <Grid2 size={4}>
          <Typography variant="h6" gutterBottom className='summary-title'>
            {p.cartEntry.product.name}
          </Typography>
          <Typography variant="body2" gutterBottom>
            {p.cartEntry.product.descr}
          </Typography>
        </Grid2>
        <Grid2 size={4}>
        <NumberField.Root id={id} defaultValue={1} min={1} value={p?.cartEntry?.amount} className={styles.Field} disabled={disableButton} onValueChange={async (value) => {
            setDisableButton(true);
            let currentEntryIndex = cartEntries.findIndex(e => e.product.id == p.cartEntry.product.id);

            cartEntries[currentEntryIndex] = {
              ...cartEntries[currentEntryIndex],
              amount: value as number
            }
            
            await updateCart(p.cartEntry.product.id, value as number, cartEntries[currentEntryIndex].id)

            setCartEntries([...cartEntries]);
            setDisableButton(false);
          }
        }>
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
            Total: {FormatBRL(p.cartEntry.product.price*p.cartEntry.amount)}
          </Typography>
        </Grid2>
        <Grid2>
          <DeleteIcon onClick={async ()=> {
            if (disableButton) {
              return; 
            }
            setDisableButton(true)

            await deleteCartEntry(p.cartEntry.id)

            const newCartEntries = cartEntries?.filter((cartEntry) => cartEntry.id !== p.cartEntry.id);

            setCartEntries(newCartEntries);
            setDisableButton(false)
          }} />
        </Grid2>
      </Grid2>
      </div>
  )
}


