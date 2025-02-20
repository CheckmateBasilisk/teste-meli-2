import Typography from '@mui/material/Typography';
import Rating from '@mui/material/Rating';
import { FormatBRL } from '../../../utils';

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



export default function ProductInfo(p : params) {
  return (
    <div>
      <Typography variant="h3" gutterBottom>
        {p.product.name}
      </Typography>

      <Typography component="legend">Ratings:</Typography>
      <Rating name="read-only" precision={0.5} value={p.product.rating} readOnly />

      <Typography variant="body1" gutterBottom>
        {FormatBRL(p.product.price)}
      </Typography>

      <Typography variant="body1" gutterBottom>
        {p.product.descr}
      </Typography>
    </div>
);
}