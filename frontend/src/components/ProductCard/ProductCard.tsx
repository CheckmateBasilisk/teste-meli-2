import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import CardActionArea from '@mui/material/CardActionArea';
import CardActions from '@mui/material/CardActions';
import Divider from '@mui/material/Divider';
import { FormatBRL } from '../../utils';



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

export default function ProductCard(p : params) {
  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardActionArea>
        <CardMedia
          component="img"
          height="140"
          image={p.product.image}
          alt={p.product.alt}
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            {p.product.name}
          </Typography>
          <Typography variant="body2" sx={{ color: 'text.secondary' }}>
            {p.product.descr}
          </Typography>
        </CardContent>
      </CardActionArea>

      <Divider variant="middle" flexItem />
      
      <CardContent >
        <Typography variant="body1" gutterBottom>
          {FormatBRL(p.product.price)}
        </Typography>
      </CardContent>
      
      <CardActions>
        <Button size="small">Add to Cart</Button>
      </CardActions>
    </Card>
  );
}