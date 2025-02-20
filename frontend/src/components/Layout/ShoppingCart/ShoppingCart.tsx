import * as React from 'react';
import { styled, useTheme } from '@mui/material/styles';
import Box from '@mui/material/Box';
import Drawer from '@mui/material/Drawer';
import Divider from '@mui/material/Divider';
import IconButton from '@mui/material/IconButton';
import Button from '@mui/material/Button';
import ShoppingCartIcon from '@mui/icons-material/ShoppingCart';
import ChevronLeftIcon from '@mui/icons-material/ChevronLeft';
import ChevronRightIcon from '@mui/icons-material/ChevronRight';
import Typography from '@mui/material/Typography';

import List from '@mui/material/List';
import ListItem from '@mui/material/ListItem';

import itemImg from '/home/bagatini/Programming/teste-meli-2/frontend/src/assets/productImages/colorful-didactic-wooden-train-toy-for-preschool-aged-kids.webp'
import ProductSummary from '../../ProductSummary/ProductSummary';
import { FormatBRL } from '../../../utils';
import { useContext } from 'react';
import { CartContext } from '../../../context/CartContext';


const drawerWidth = "50%";

const Main = styled('main', { shouldForwardProp: (prop) => prop !== 'open' })<{
  open?: boolean;
}>(({ theme }) => ({
  flexGrow: 1,
  padding: theme.spacing(3),
  transition: theme.transitions.create('margin', {
    easing: theme.transitions.easing.sharp,
    duration: theme.transitions.duration.leavingScreen,
  }),
  marginRight: -drawerWidth,
  /**
   * This is necessary to enable the selection of content. In the DOM, the stacking order is determined
   * by the order of appearance. Following this rule, elements appearing later in the markup will overlay
   * those that appear earlier. Since the Drawer comes after the Main content, this adjustment ensures
   * proper interaction with the underlying content.
   */
  position: 'relative',
  variants: [
    {
      props: ({ open }) => open,
      style: {
        transition: theme.transitions.create('margin', {
          easing: theme.transitions.easing.easeOut,
          duration: theme.transitions.duration.enteringScreen,
        }),
        marginRight: 0,
      },
    },
  ],
}));

const DrawerHeader = styled('div')(({ theme }) => ({
  display: 'flex',
  alignItems: 'center',
  padding: theme.spacing(0, 1),
  // necessary for content to be below app bar
  ...theme.mixins.toolbar,
  justifyContent: 'flex-start',
}));

export default function ShoppingCart() {
  const theme = useTheme();
  const [open, setOpen] = React.useState(false);
  const { quantity } = useContext(CartContext);

  const handleDrawerOpen = () => {
    setOpen(true);
  };

  const handleDrawerClose = () => {
    setOpen(false);
  };

  var data = [
    { image: itemImg,
      price: 999,
      alt: "item.alt",
      title: "Filme Fotográfico 50D Kodak Color",
      desc : "lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo ",
    },
    { image: itemImg,
      price: 999,
      alt: "item.alt",
      title: "Filme Fotográfico 50D Kodak Color",
      desc : "lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo ",
    },
    { image: itemImg,
      price: 999,
      alt: "item.alt",
      title: "Filme Fotográfico 50D Kodak Color",
      desc : "lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo ",
    },
    { image: itemImg,
      price: 999,
      alt: "item.alt",
      title: "Filme Fotográfico 50D Kodak Color",
      desc : "lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo lalilulelo ",
    }
  ]

  return (
    <Box sx={{ display: 'flex' }}>
      <IconButton
            color="inherit"
            aria-label="open drawer"
            edge="end"
            onClick={handleDrawerOpen}
            sx={[open && { display: 'none' }]}
          >
            <ShoppingCartIcon className='shopping-cart-icon'/>
      </IconButton>

      <Drawer
        sx={{
          width: drawerWidth,
          flexShrink: 0,
          '& .MuiDrawer-paper': {
            width: drawerWidth,
          },
        }}
        variant="persistent"
        anchor="right"
        open={open}
        PaperProps={{
          sx: {
            backgroundColor: "white"
          }
        }}
      
      >
        <DrawerHeader>
          <IconButton onClick={handleDrawerClose}>
            {theme.direction === 'rtl' ? <ChevronLeftIcon /> : <ChevronRightIcon />}
          </IconButton>
          <Typography variant="h6" component="div">
            Shopping Cart
          </Typography>
        </DrawerHeader>
        {`Context counting: ${quantity}`}
        <Divider />
          {/* list of products currently in cart */}
        <List>
          {/* there should be a list of products currently in the shopping cart here! */}
          {data.map((data) => (
            <ListItem disablePadding key={data.title}>
              <ProductSummary image={data.image} price={data.price} alt={data.alt} title={data.title} desc={data.desc}/>
            </ListItem>
          ))}
        </List>
        <Divider />
        {/* RESUMO DA COMPRA E BOTAO DE AVANÇAR */}
        <Typography variant="h6">
          Total: { FormatBRL(data.reduce((tot,d) => {return tot+(d.price)}, 0)) }

        </Typography>
        <Button variant="contained" >Buy Now</Button>

      </Drawer>
    </Box>
  );
}