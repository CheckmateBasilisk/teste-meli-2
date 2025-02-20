import Typography from '@mui/material/Typography';
import Grid2 from '@mui/material/Grid2'
import { NumberField } from '@base-ui-components/react/number-field';
import React, { useState } from 'react';
import styles from './style.module.scss';
import RemoveIcon from '@mui/icons-material/Remove';
import AddIcon from '@mui/icons-material/Add';
import { FormatBRL } from '../../utils';

interface params {
  image: string
  alt: string
  title: string
  desc: string
  price: number
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
  const [quantity, setQuantity] = useState(1);

  return (

    // FIXME: THE SPACING ISNT RIGHT. THE IMGS KEEP ON CLIPPING
      <div className='product-summary'>
      <Grid2 container size={12}>
        <Grid2 size={3}>
        <div className='summary-img' >
          <img src={p.image} alt={p.alt} height="100rem"/>
        </div>
        </Grid2>
        <Grid2 size={5}>
          <Typography variant="h6" gutterBottom className='summary-title'>
            {p.title}
          </Typography>
          <Typography variant="body2" gutterBottom>
            {p.desc}
          </Typography>
        </Grid2>
        <Grid2 size={4}>

          {/* i should move this to a component eventually... */}
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
            Total: {FormatBRL(p.price*quantity)}
          </Typography>

          {/* product amount */}
        </Grid2>
      </Grid2>
      </div>
  )
}


